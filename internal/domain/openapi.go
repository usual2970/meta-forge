package domain

import (
	"context"

	"github.com/labstack/echo/v5"
)

const OpenapiAppChecked = "checked"

type OpenapiCallReq struct {
	Service  string                 `json:"service"`
	Api      string                 `json:"api"`
	Args     map[string]interface{} `json:"args"`
	AsyncUrl string                 `json:"asyncUrl"`
}
type IOpenapiUsecase interface {
	// Call 调用service 下的 api
	Call(ctx context.Context, req *OpenapiCallReq) (interface{}, error)

	Notice(ctx echo.Context) error
}
