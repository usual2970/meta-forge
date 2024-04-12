package ai

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/repository/secret"
	"github.com/usual2970/meta-forge/internal/util/xtools"
	"github.com/usual2970/meta-forge/internal/util/zhipu"

	"github.com/pocketbase/pocketbase/tests"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/tools"
	"github.com/tmc/langchaingo/tools/serpapi"
	client "github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate/entities/models"
)

const testDataDir = "/Users/liuxuanyao/work/github.com/usual2970/meta-forge/pb_data"

func Test_usecase_UploadDoc(t *testing.T) {
	app, err := tests.NewTestApp(testDataDir)
	if err != nil {
		t.Fatal(err)
	}
	defer app.Cleanup()
	type args struct {
		ctx context.Context
		req *domain.AiUploadDocReq
	}
	tests := []struct {
		name    string
		a       *usecase
		args    args
		wantErr bool
	}{
		{
			name: "1",
			a: &usecase{
				secretRepo: secret.NewRepository(),
			},
			args: args{
				ctx: context.Background(),
				req: &domain.AiUploadDocReq{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &usecase{
				secretRepo: secret.NewRepository(),
			}
			if rs, err := a.UploadDoc(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("usecase.UploadDoc() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				t.Log(rs)
			}
		})
	}
}

func Test_usecase_SearchDoc(t *testing.T) {
	type args struct {
		ctx context.Context
		req *domain.AiSearchDocReq
	}
	tests := []struct {
		name    string
		a       *usecase
		args    args
		want    *domain.AiSearchDocResp
		wantErr bool
	}{
		{
			name: "1",
			a:    &usecase{},
			args: args{
				ctx: context.Background(),
				req: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &usecase{}
			got, err := a.SearchDoc(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.SearchDoc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.SearchDoc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateCollections(t *testing.T) {

	config := client.Config{
		Scheme: "https",
		Host:   "weaviate.ikit.fun",
	}
	client, err := client.NewClient(config)
	if err != nil {
		t.Fatal(err)
	}
	className := "Contract"

	emptyClass := &models.Class{
		Class: className,
	}

	// Add the class to the schema
	err = client.Schema().ClassCreator().
		WithClass(emptyClass).
		Do(context.Background())

	if err != nil {
		t.Fatal(err)
	}
}

func TestAgent(t *testing.T) {
	llm := zhipu.NewZhipu("7a909dd632f00f28a0d08efba999b3dc.50FHMHrklQIaaD16", llms.WithTemperature(0.95), llms.WithTopP(0.70))

	os.Setenv("SERPAPI_API_KEY", "0cc261b4c1814a7f2d9ed79df97df998a0232c3059df1733532cce6610f04c42")

	searchTool, err := serpapi.New()
	if err != nil {
		t.Fatal(err)
	}

	toolList := []tools.Tool{xtools.GetInfo{}, xtools.GetOrderNum{}, searchTool}

	// a := agents.NewOpenAIFunctionsAgent(llm,
	// 	toolList,
	// 	agents.NewOpenAIOption().WithSystemMessage("you are a helpful assistant"),
	// 	agents.NewOpenAIOption().WithExtraMessages([]prompts.MessageFormatter{
	// 		prompts.NewHumanMessagePromptTemplate("please be strict", nil),
	// 	}),
	// )

	a, err := agents.Initialize(llm, toolList, agents.ConversationalReactDescription, agents.WithParserErrorHandler(agents.NewParserErrorHandler(func(s string) string {
		return "Check your output and make sure it conforms, use the Action/Action Input syntax"
	})), agents.WithMemory(memory.NewConversationBuffer()))
	if err != nil {
		t.Fatal(err)
	}
	// e := agents.NewExecutor(a, toolList)

	rs, err := chains.Run(context.Background(), a, "陈奕迅今年多大了")

	t.Log(rs, err)

	rs1, err := chains.Run(context.Background(), a, "请用英文回答")

	t.Log(rs1, err)

	rs2, err := chains.Run(context.Background(), a, "你好")

	t.Log(rs2, err)

	rs3, err := chains.Run(context.Background(), a, `
	业务员王万利和何建忠各有多少个审批完成的进件订单`)

	t.Log(rs3, err)

	rs4, err := chains.Run(context.Background(), a, "请用中文回答")

	t.Log(rs4, err)

}
