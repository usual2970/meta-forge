package openapi

import (
	"context"

	"github.com/usual2970/meta-forge/internal/util/app"
	"github.com/usual2970/meta-forge/internal/util/xcontext"
)

func Log(ctx context.Context, next Next) {
	next()

	openapiContext, err := xcontext.GetOpenapiContext(ctx)
	if err != nil {
		openapiContext = &xcontext.OpenapiContext{
			Error: err,
		}
		xcontext.SetOpenapiContext(ctx, openapiContext)
		return
	}

	if openapiContext.Error != nil {
		app.Get().Logger().Error("openapi call failed", "ctx", openapiContext)
	} else {
		app.Get().Logger().Info("openapi call Succ", "ctx", openapiContext)
	}

}
