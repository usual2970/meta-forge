package wenxin

import (
	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/util/resp"

	"github.com/labstack/echo/v5"
)

type controller struct {
	uc domain.IWenxinUsecase
}

func (c *controller) Completion(ctx echo.Context) error {
	param := &domain.WenxinCompletionReq{}
	if err := ctx.Bind(param); err != nil {
		return resp.Err(ctx, err)
	}

	rs, err := c.uc.Completion(ctx.Request().Context(), param)
	if err != nil {
		return resp.Err(ctx, err)
	}
	return resp.Succ(ctx, rs)
}

func Register(route *echo.Echo, uc domain.IWenxinUsecase) {
	c := &controller{uc: uc}
	group := route.Group("/v1/wenxin")

	group.POST("/completion", c.Completion)
}
