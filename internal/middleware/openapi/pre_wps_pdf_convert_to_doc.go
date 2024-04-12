package openapi

import (
	"context"
	"errors"
	"strings"

	"github.com/usual2970/meta-forge/internal/util/xcontext"
)

func PreWpsPdfConvertToDoc(ctx context.Context, next Next) {
	openapiContext, err := xcontext.GetOpenapiContext(ctx)
	if err != nil {
		openapiContext = &xcontext.OpenapiContext{
			Error: err,
		}
		xcontext.SetOpenapiContext(ctx, openapiContext)
		return
	}

	param := openapiContext.Param

	docType, ok := param["docType"]
	if !ok {
		openapiContext.Error = errors.New("missing doc type")
		xcontext.SetOpenapiContext(ctx, openapiContext)
		return
	}

	openapiContext.Path = strings.Replace(openapiContext.Path, ":office_type", docType.(string), -1)
	xcontext.SetOpenapiContext(ctx, openapiContext)

	next()
}
