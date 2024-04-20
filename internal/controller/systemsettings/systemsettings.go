package systemsettings

import (
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

func Register(route *echo.Echo, usecase domain.ISystemSettingsUsecase) {
	c := &controller{
		usecase: usecase,
	}
	route.POST("/systemsettings/initialize", c.Initail)
}
