package xtools

import (
	"context"
	"strings"

	"github.com/tmc/langchaingo/callbacks"
)

type GetInfo struct {
	CallbacksHandler callbacks.Handler
}

func (c GetInfo) Name() string {
	return "getUserInfo"
}

func (c GetInfo) Description() string {
	return `"获取信息"
			"用于获取生成订单所需的各种信息"
			"输入参数为:信息类型"
			"信息类型：一次只输入一种类型，如：姓名"`
}

func (c GetInfo) Call(ctx context.Context, input string) (string, error) {
	if c.CallbacksHandler != nil {
		c.CallbacksHandler.HandleToolStart(ctx, input)
	}

	input = strings.Replace(input, "\nObservation", "", -1)
	input = strings.Replace(input, ":", "", -1)

	switch input {
	case "姓名":
		return "刘旋尧", nil
	case "手机号":
		return "18257140570", nil
	case "身份证号":
		return "440123199909090000", nil
	}

	return input, nil
}
