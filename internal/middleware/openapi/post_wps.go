package openapi

import (
	"context"
	"encoding/json"

	"github.com/usual2970/meta-forge/internal/domain/constant"
	"github.com/usual2970/meta-forge/internal/util/xcontext"
)

type WpsConvert2DocResp struct {
	Code    int         `json:"code"`
	Hint    string      `json:"hint"`
	Message string      `json:"message"`
	Extra   string      `json:"extra"`
	Data    interface{} `json:"data"`
}

func PostWps(ctx context.Context, next Next) {
	next()

	openapiContext, err := xcontext.GetOpenapiContext(ctx)
	if err != nil {
		openapiContext = &xcontext.OpenapiContext{
			Error: err,
		}
		xcontext.SetOpenapiContext(ctx, openapiContext)
		return
	}

	bts, _ := json.Marshal(openapiContext.Response)

	resp := &WpsConvert2DocResp{}
	if err := json.Unmarshal(bts, resp); err != nil {
		openapiContext.Error = err
		xcontext.SetOpenapiContext(ctx, openapiContext)
		return
	}

	if resp.Code != 0 {
		openapiContext.Error = constant.NewXError(resp.Code, string(bts))
		xcontext.SetOpenapiContext(ctx, openapiContext)
		return
	}

	openapiContext.Response = resp.Data

	xcontext.SetOpenapiContext(ctx, openapiContext)
}
