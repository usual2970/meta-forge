package domain

import "context"

type WenxinCompletionReq struct {
	Model   string `form:"model"`
	Content string `form:"content"`
}
type IWenxinUsecase interface {
	Completion(ctx context.Context, req *WenxinCompletionReq) (string, error)
}
