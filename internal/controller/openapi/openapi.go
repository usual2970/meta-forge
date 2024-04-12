package openapi

import (
	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/middleware"
	"github.com/usual2970/meta-forge/internal/util/resp"

	"github.com/labstack/echo/v5"
)

type controller struct {
	uc domain.IOpenapiUsecase
}

func (c *controller) Call(ctx echo.Context) error {
	param := &domain.OpenapiCallReq{}
	if err := ctx.Bind(param); err != nil {
		return resp.Err(ctx, err)
	}
	param.Service = ctx.PathParam("service")
	param.Api = ctx.PathParam("api")

	rs, err := c.uc.Call(ctx.Request().Context(), param)
	if err != nil {
		return resp.Err(ctx, err)
	}
	return resp.Succ(ctx, rs)
}

func (c *controller) Notice(ctx echo.Context) error {
	return c.uc.Notice(ctx)
}

func Register(route *echo.Echo, uc domain.IOpenapiUsecase) {
	c := &controller{uc: uc}
	group := route.Group("/openapi/v1")
	group.Use(middleware.AuthToken())

	group.POST("/:service/:api", c.Call)

	notifyGroup := route.Group("/notify/v1")

	notifyGroup.Any("/:service/:notice", c.Notice)
}
