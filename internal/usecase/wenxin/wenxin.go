package wenxin

import (
	"context"

	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/util/app"

	ernie "github.com/anhao/go-ernie"
)

var wenxinClient *ernie.Client

func getWenxinClient() *ernie.Client {

	if wenxinClient == nil {
		record, err := app.Get().Dao().FindFirstRecordByData("secrets", "uri", "baidu_wenxin")
		if err != nil {
			panic(err)
		}

		wenxinClient = ernie.NewDefaultClient(record.GetString("api_key"), record.GetString("secret_key"))
	}

	return wenxinClient
}

type usecase struct{}

func NewUsecase() domain.IWenxinUsecase {
	return &usecase{}
}

func (s *usecase) Completion(ctx context.Context, req *domain.WenxinCompletionReq) (string, error) {
	client := getWenxinClient()

	resp, err := client.CreateErnieBotTurboChatCompletion(ctx, ernie.ErnieBotTurboRequest{
		Messages: []ernie.ChatCompletionMessage{
			{
				Role:    ernie.MessageRoleUser,
				Content: req.Content,
			},
		},
	})

	if err != nil {
		return "", err
	}

	return resp.Result, nil
}
