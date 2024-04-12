package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/util/app"
	xhttp "github.com/usual2970/meta-forge/internal/util/http"
	"github.com/usual2970/meta-forge/internal/util/xchains"
	"github.com/usual2970/meta-forge/internal/util/xtools"
	"github.com/usual2970/meta-forge/internal/util/zhipu"

	"github.com/google/uuid"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
	"github.com/tmc/langchaingo/tools"
	"github.com/tmc/langchaingo/vectorstores/weaviate"

	"github.com/tmc/langchaingo/chains"
)

type usecase struct {
	secretRepo domain.ISecretRepository
}

func NewUsecase(secretRepo domain.ISecretRepository) domain.IAiUsecase {
	return &usecase{
		secretRepo: secretRepo,
	}
}

func (a *usecase) OrderNum(ctx context.Context, req *domain.AiOrderNumReq) (*domain.AiOrderNumResp, error) {
	zhipuConfig, err := a.secretRepo.Get(ctx, "uri='zhipu'")
	if err != nil {
		return nil, err
	}
	llm := zhipu.NewZhipu(zhipuConfig.SecretKey)

	toolList := []tools.Tool{xtools.GetOrderNum{}}

	e, err := agents.Initialize(llm, toolList, agents.ZeroShotReactDescription)
	if err != nil {
		app.Get().Logger().Error("Initialize angents failed", "err", err)
		return nil, err
	}

	rs, err := chains.Run(context.Background(), e, req.Question)
	if err != nil {
		app.Get().Logger().Error("Get result failed", "err", err)
		return nil, err
	}

	return &domain.AiOrderNumResp{
		Result: rs,
	}, nil
}

func (a *usecase) UploadDocByLink(ctx context.Context, req *domain.AiUploadDocByLinkReq) (*domain.AiUploadDocResp, error) {

	resp, err := xhttp.Req(req.Link, http.MethodGet, nil, map[string]string{})

	if err != nil {
		return nil, err
	}

	return a.uploadDoc(ctx, req.Index, bytes.NewReader(resp), int64(len(resp)))
}

func (a *usecase) uploadDoc(ctx context.Context, index string, f io.ReaderAt, size int64) (*domain.AiUploadDocResp, error) {
	pdf := documentloaders.NewPDF(f, size)

	split := textsplitter.NewRecursiveCharacter()
	split.ChunkOverlap = 80
	split.ChunkSize = 512

	docs, err := pdf.LoadAndSplit(context.Background(), split)
	if err != nil {
		return nil, err
	}

	// 生成向量数据并保存到向量数据库

	zhipuConfig, err := a.secretRepo.Get(ctx, "uri='zhipu'")
	if err != nil {
		return nil, err
	}
	llm := zhipu.NewZhipu(zhipuConfig.SecretKey)
	// wenxinConfig, err := a.secretRepo.Get(ctx, "uri='baidu_wenxin'")
	// if err != nil {
	// 	return nil, err
	// }

	// llm := wenxin.NewWenxin(wenxinConfig.ApiKey, wenxinConfig.SecretKey)

	embed, err := embeddings.NewEmbedder(llm)
	if err != nil {
		return nil, err
	}

	weaviateConfig, err := a.secretRepo.Get(ctx, "uri='weaviate'")
	if err != nil {
		return nil, err
	}

	uuid := uuid.New().String()
	store, err := weaviate.New(
		weaviate.WithScheme(weaviateConfig.Ext["scheme"]),
		weaviate.WithHost(weaviateConfig.Ext["host"]),
		weaviate.WithEmbedder(embed),
		weaviate.WithNameSpace(uuid),
		weaviate.WithIndexName(index),

		weaviate.WithAPIKey(weaviateConfig.SecretKey),
	)

	if err != nil {
		return nil, err
	}

	docChunks := Chunk[schema.Document](docs, 16)
	for _, doc := range docChunks {
		_, err = store.AddDocuments(ctx, doc)

		if err != nil {
			return nil, err
		}
	}

	return &domain.AiUploadDocResp{
		Namespace: uuid,
		Index:     index,
	}, nil
}

func (a *usecase) UploadDoc(ctx context.Context, req *domain.AiUploadDocReq) (*domain.AiUploadDocResp, error) {

	f, err := req.File.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return a.uploadDoc(ctx, req.Index, f, req.File.Size)
}

func (a *usecase) SearchDoc(ctx context.Context, req *domain.AiSearchDocReq) (*domain.AiSearchDocResp, error) {
	zhipuConfig, err := a.secretRepo.Get(ctx, "uri='zhipu'")
	if err != nil {
		return nil, err
	}

	defaultOptions := make([]llms.CallOption, 0)
	if req.Func != "" {
		function := &zhipu.Function{}
		err = json.Unmarshal([]byte(req.Func), function)
		if err != nil {
			return nil, err
		}
		defaultOptions = append(defaultOptions, llms.WithFunctions([]llms.FunctionDefinition{
			{
				Name:        function.Name,
				Parameters:  function.Parameters,
				Description: function.Description,
			},
		}))
	}
	llm := zhipu.NewZhipu(zhipuConfig.SecretKey, defaultOptions...)
	embed, err := embeddings.NewEmbedder(llm)
	if err != nil {
		return nil, err
	}

	weaviateConfig, err := a.secretRepo.Get(ctx, "uri='weaviate'")
	if err != nil {
		return nil, err
	}

	store, err := weaviate.New(
		weaviate.WithScheme(weaviateConfig.Ext["scheme"]),
		weaviate.WithHost(weaviateConfig.Ext["host"]),
		weaviate.WithEmbedder(embed),
		weaviate.WithNameSpace(req.Namespace),
		weaviate.WithIndexName(req.Index),

		weaviate.WithAPIKey(weaviateConfig.SecretKey),
	)

	if err != nil {
		return nil, err
	}

	docs, err := store.SimilaritySearch(ctx, req.Question, 2)
	if err != nil {
		return nil, err
	}
	qaChain := chains.LoadStuffQA(llm)

	rs, err := chains.Call(ctx, qaChain, map[string]any{
		"input_documents": docs,
		"question":        req.Question,
	})

	if err != nil {
		return nil, err
	}

	return &domain.AiSearchDocResp{
		Result: rs,
		Docs:   docs,
	}, nil
}

func (a *usecase) SummaryDoc(ctx context.Context, req *domain.AiSummaryDocReq) (*domain.AiSummaryDocResp, error) {
	file, err := xhttp.Req(req.Link, http.MethodGet, nil, map[string]string{})

	if err != nil {
		return nil, err
	}

	pdf := documentloaders.NewPDF(bytes.NewReader(file), int64(len(file)))

	split := textsplitter.NewRecursiveCharacter()
	split.ChunkOverlap = 80
	split.ChunkSize = 512

	docs, err := pdf.LoadAndSplit(context.Background(), split)
	if err != nil {
		return nil, err
	}

	// 生成向量数据并保存到向量数据库

	zhipuConfig, err := a.secretRepo.Get(ctx, "uri='zhipu'")
	if err != nil {
		return nil, err
	}
	llm := zhipu.NewZhipu(zhipuConfig.SecretKey)

	chain := xchains.LoadStuffSummarization(llm)

	rs, err := chains.Run(ctx, chain, docs, chains.WithModel("glm-4"))

	if err != nil {
		return nil, err
	}

	return &domain.AiSummaryDocResp{
		Result: rs,
	}, nil
}

func Chunk[T any](arr []T, chunkSize int) [][]T {
	var chunks [][]T
	for i := 0; i < len(arr); i += chunkSize {
		end := i + chunkSize
		if end > len(arr) {
			end = len(arr)
		}
		chunk := arr[i:end]
		chunks = append(chunks, chunk)
	}
	return chunks
}
