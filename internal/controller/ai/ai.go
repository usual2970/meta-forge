package ai

import (
	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/util/resp"

	"github.com/labstack/echo/v5"
)

type controller struct {
	uc domain.IAiUsecase
}

func (c *controller) SearchDoc(ctx echo.Context) error {
	req := &domain.AiSearchDocReq{}
	if err := ctx.Bind(req); err != nil {
		return err
	}
	rs, err := c.uc.SearchDoc(ctx.Request().Context(), req)
	if err != nil {
		return resp.Err(ctx, err)
	}
	return resp.Succ(ctx, rs)
}

func (c *controller) UploadDoc(ctx echo.Context) error {
	fh, err := ctx.FormFile("file")
	if err != nil {
		return resp.Err(ctx, err)
	}

	req := &domain.AiUploadDocReq{}

	if err := ctx.Bind(req); err != nil {
		return resp.Err(ctx, err)
	}

	req.File = fh
	rs, err := c.uc.UploadDoc(ctx.Request().Context(), req)
	if err != nil {
		return resp.Err(ctx, err)
	}
	return resp.Succ(ctx, rs)
}

func (c *controller) UploadDocByLink(ctx echo.Context) error {

	req := &domain.AiUploadDocByLinkReq{}

	if err := ctx.Bind(req); err != nil {
		return resp.Err(ctx, err)
	}
	rs, err := c.uc.UploadDocByLink(ctx.Request().Context(), req)
	if err != nil {
		return resp.Err(ctx, err)
	}
	return resp.Succ(ctx, rs)
}

func (c *controller) SummaryDoc(ctx echo.Context) error {
	req := &domain.AiSummaryDocReq{}
	if err := ctx.Bind(req); err != nil {
		return err
	}
	rs, err := c.uc.SummaryDoc(ctx.Request().Context(), req)
	if err != nil {
		return resp.Err(ctx, err)
	}
	return resp.Succ(ctx, rs)
}

func (c *controller) OrderNum(ctx echo.Context) error {
	req := &domain.AiOrderNumReq{}
	if err := ctx.Bind(req); err != nil {
		return err
	}
	rs, err := c.uc.OrderNum(ctx.Request().Context(), req)
	if err != nil {
		return resp.Err(ctx, err)
	}
	return resp.Succ(ctx, rs)
}

func Register(route *echo.Echo, uc domain.IAiUsecase) {
	c := &controller{uc: uc}
	group := route.Group("/ai/v1")

	group.POST("/upload-doc", c.UploadDoc)

	group.POST("/upload-doc-by-link", c.UploadDocByLink)

	group.POST("/search-doc", c.SearchDoc)

	group.POST("/summary-doc", c.SummaryDoc)

	group.POST("/order-num", c.OrderNum)
}
