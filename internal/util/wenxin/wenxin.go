package wenxin

import (
	"context"
	"errors"
	"sync"

	ernie "github.com/anhao/go-ernie"
	"github.com/tmc/langchaingo/llms"
)

var wenxinClient *ernie.Client

var wenxinOnce sync.Once

func getWenxinClient(apiKey, secretkey string) *ernie.Client {

	wenxinOnce.Do(func() {

		wenxinClient = ernie.NewDefaultClient(apiKey, secretkey)
	})

	return wenxinClient
}

type Wenxin struct {
	client *ernie.Client
}

func NewWenxin(apiKey, secretKey string) *Wenxin {
	client := getWenxinClient(apiKey, secretKey)

	return &Wenxin{
		client: client,
	}
}

func (w *Wenxin) Call(ctx context.Context, prompt string, options ...llms.CallOption) (string, error) {

	r, err := w.Generate(ctx, []string{prompt}, options...)
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", errors.New("no response")
	}
	return r[0].Text, nil
}

func (w *Wenxin) Generate(ctx context.Context, prompts []string, options ...llms.CallOption) ([]*llms.Generation, error) {

	opts := llms.CallOptions{}
	for _, opt := range options {
		opt(&opts)
	}

	rs := make([]*llms.Generation, 0)
	for _, prompt := range prompts {

		resp, err := w.client.CreateErnieBotTurboChatCompletion(ctx, ernie.ErnieBotTurboRequest{
			Messages: []ernie.ChatCompletionMessage{
				{
					Role:    ernie.MessageRoleUser,
					Content: prompt,
				},
			},
			TopP:        float32(opts.TopP),
			Temperature: float32(opts.Temperature),
		})

		if err != nil {
			return nil, err
		}

		item := &llms.Generation{
			Text: resp.Result,
		}

		rs = append(rs, item)
	}
	return rs, nil
}

func (w *Wenxin) CreateEmbedding(ctx context.Context, texts []string) ([][]float32, error) {

	resp, err := w.client.CreateEmbeddings(ctx, ernie.EmbeddingRequest{
		Input: texts,
	})
	if err != nil {
		return nil, err
	}
	rs := make([][]float32, 0, len(resp.Data))
	for _, items := range resp.Data {
		temp := make([]float32, 0, len(items.Embedding))

		for _, item := range items.Embedding {
			temp = append(temp, float32(item))
		}

		rs = append(rs, temp)
	}

	return rs, nil

}
