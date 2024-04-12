package openapi

import (
	"context"

	"github.com/usual2970/meta-forge/internal/util/xcontext"
)

func PostWpsNotices(ctx context.Context, next Next) {
	next()

	openapiContext, err := xcontext.GetNoticeContext(ctx)
	if err != nil {
		openapiContext = &xcontext.NoticeContext{
			Error: err,
		}
		openapiContext.ResponseType = "string"
		openapiContext.Response = "Failed"
		xcontext.SetNoticeContext(ctx, openapiContext)
		return
	}

	openapiContext.ResponseType = "string"
	openapiContext.Response = "Success"
	xcontext.SetNoticeContext(ctx, openapiContext)
}
