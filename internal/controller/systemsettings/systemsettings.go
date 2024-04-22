package systemsettings

import (
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/util/resp"
)

type controller struct {
	usecase domain.ISystemSettingsUsecase
}

func (c *controller) Initail(ctx echo.Context) error {
	param := &domain.InitializeReq{}
	if err := ctx.Bind(param); err != nil {
		return resp.Err(ctx, err)
	}

	if err := c.usecase.Initialize(ctx.Request().Context(), param); err != nil {
		return resp.Err(ctx, err)
	}

	return resp.Succ(ctx, nil)

}

func (c *controller) Get(ctx echo.Context) error {
	key := ctx.QueryParam("key")
	value, err := c.usecase.Get(ctx.Request().Context(), key)
	if err != nil {
		return resp.Err(ctx, err)
	}

	return resp.Succ(ctx, value)
}

func (c *controller) BatchGet(ctx echo.Context) error {
	keys := ctx.QueryParam("keys")
	value, err := c.usecase.BatchGet(ctx.Request().Context(), strings.Split(keys, ","))
	if err != nil {
		return resp.Err(ctx, err)
	}

	return resp.Succ(ctx, value)
}

func Register(route *echo.Group, usecase domain.ISystemSettingsUsecase) {
	c := &controller{
		usecase: usecase,
	}
	route.POST("/systemsettings/initialize", c.Initail)
	route.GET("/systemsettings/get", c.Get)
	route.GET("/systemsettings/batch-get", c.BatchGet)
}
