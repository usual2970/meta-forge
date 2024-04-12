package mingdao

import (
	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/util/resp"

	"github.com/labstack/echo/v5"
)

type controller struct {
	uc domain.IMingdaoUsecase
}

func (c *controller) GetPassId(ctx echo.Context) error {
	param := &domain.MingdaoGetPassIdReq{}
	if err := ctx.Bind(param); err != nil {
		return resp.Err(ctx, err)
	}

	rs, err := c.uc.GetPassId(ctx.Request().Context(), param)
	if err != nil {
		return resp.Err(ctx, err)
	}
	return resp.Succ(ctx, rs)
}

func Register(route *echo.Echo, uc domain.IMingdaoUsecase) {
	c := &controller{uc: uc}
	group := route.Group("/v1/mingdao")
	group.POST("/passid", c.GetPassId)
}
