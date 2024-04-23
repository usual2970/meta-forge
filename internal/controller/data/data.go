package data

import (
	"github.com/labstack/echo/v5"
	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/util/resp"
)

type controller struct {
	uc domain.IDataUsecase
}

func (c *controller) List(ctx echo.Context) error {
	req := &domain.DataListReq{}
	if err := ctx.Bind(req); err != nil {
		return err
	}

	rs, err := c.uc.List(ctx.Request().Context(), req)
	if err != nil {
		return resp.Err(ctx, err)
	}

	return resp.Succ(ctx, rs)
}

func Register(route *echo.Group, usecase domain.IDataUsecase) {
	c := &controller{
		uc: usecase,
	}

	route.GET("/data/list", c.List)
}
