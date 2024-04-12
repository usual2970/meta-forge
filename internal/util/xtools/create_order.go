package xtools

import (
	"context"

	"github.com/tmc/langchaingo/callbacks"
)

type CreateOrder struct {
	CallbacksHandler callbacks.Handler
}

func (c CreateOrder) Name() string {
	return "createOrder"
}

func (c CreateOrder) Description() string {
	return `"创建订单"
			"用于创建订单并返回订单号"
			"输入参数为:姓名-手机号-身份证号"
			"姓名:用户姓名"
			"手机号：用户手机号"
			"身份证号：用户身份证号"`
}

func (c CreateOrder) Call(ctx context.Context, input string) (string, error) {
	if c.CallbacksHandler != nil {
		c.CallbacksHandler.HandleToolStart(ctx, input)
	}
	return "12341asdf23", nil
}
