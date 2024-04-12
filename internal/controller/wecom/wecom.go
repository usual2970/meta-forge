package wecom

import (
	"io"

	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/util/resp"
	"github.com/usual2970/meta-forge/internal/util/wecom"

	"github.com/labstack/echo/v5"
)

type controller struct {
	uc domain.IWecomUsecase
}

func (c *controller) Notify(ctx echo.Context) error {
	param := &domain.WecomNoticeReq{}

	post, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return resp.Err(ctx, err)
	}

	param.Post = post

	param.Echostr = ctx.QueryParam("echostr")
	param.MsgSignature = ctx.QueryParam("msg_signature")
	param.Timestamp = ctx.QueryParam("timestamp")
	param.Nonce = ctx.QueryParam("nonce")

	rs, err := c.uc.Notify(ctx.Request().Context(), param)
	if err != nil {
		return resp.Err(ctx, err)
	}

	return resp.WecomVerifyUrl(ctx, rs)
}

func (c *controller) GetServiceState(ctx echo.Context) error {
	param := &wecom.GetServiceStateReq{}
	if err := ctx.Bind(param); err != nil {
		return resp.Err(ctx, err)
	}

	rs, err := c.uc.GetServiceState(ctx.Request().Context(), param)
	if err != nil {
		return resp.Err(ctx, err)
	}
	return resp.Succ(ctx, rs)
}

func (c *controller) TransServiceState(ctx echo.Context) error {
	param := &wecom.TransServiceStateReq{}
	if err := ctx.Bind(param); err != nil {
		return resp.Err(ctx, err)
	}

	err := c.uc.TransServiceState(ctx.Request().Context(), param)
	if err != nil {
		return resp.Err(ctx, err)
	}
	return resp.Succ(ctx, nil)
}

func (c *controller) ClearExecutor(ctx echo.Context) error {
	c.uc.ClearExecutor(ctx.Request().Context())
	return resp.Succ(ctx, nil)
}

func (c *controller) Accesstoken(ctx echo.Context) error {
	rs, err := c.uc.Accesstoken(ctx.Request().Context())
	if err != nil {
		return resp.Err(ctx, err)
	}
	return resp.Succ(ctx, rs)
}

func Register(route *echo.Echo, uc domain.IWecomUsecase) {
	c := &controller{uc: uc}
	group := route.Group("/api/v1")

	group.Any("/wecom/notify", c.Notify)

	group.GET("/wecom/service-state", c.GetServiceState)

	group.GET("/wecom/accesstoken", c.Accesstoken)

	group.GET("/wecom/clear-executor", c.ClearExecutor)
	group.POST("/wecom/trans-service-state", c.TransServiceState)
}
