package openapi

import (
	"context"

	"github.com/usual2970/meta-forge/internal/util/app"
	"github.com/usual2970/meta-forge/internal/util/xcontext"
)

func LogNotice(ctx context.Context, next Next) {
	next()

	openapiContext, err := xcontext.GetNoticeContext(ctx)
	if err != nil {
		openapiContext = &xcontext.NoticeContext{
			Error: err,
		}
		xcontext.SetNoticeContext(ctx, openapiContext)
		return
	}

	if openapiContext.Error != nil {
		app.Get().Logger().Error("openapi notice failed", "ctx", openapiContext)
	} else {
		app.Get().Logger().Info("openapi notice Succ", "ctx", openapiContext)
	}

}
