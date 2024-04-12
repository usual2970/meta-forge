package openapi

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/middleware/openapi"

	"github.com/usual2970/meta-forge/internal/util/app"
	"github.com/usual2970/meta-forge/internal/util/xcontext"

	"github.com/labstack/echo/v5"
	"github.com/panjf2000/ants/v2"
	"github.com/pocketbase/dbx"

	"github.com/gojek/heimdall/v7/httpclient"
)

type usecase struct {
	ch chan *xcontext.OpenapiContext
}

func NewUsecase() domain.IOpenapiUsecase {
	u := &usecase{
		ch: make(chan *xcontext.OpenapiContext, 1),
	}
	for i := 0; i < 5; i++ {
		ants.Submit(u.processJob)
	}
	return u
}

func (a *usecase) Notice(ctx echo.Context) error {
	//serviceUri := ctx.PathParam("service")
	noticeUri := ctx.PathParam("notice")

	notification, err := app.Get().Dao().FindFirstRecordByFilter("app_notifications",
		"state='published' && uri={:uri}",
		dbx.Params{"uri": noticeUri})
	if err != nil {
		return err
	}

	noticeContext := &xcontext.NoticeContext{
		Url:    notification.GetString("notice_url"),
		Method: notification.GetString("notice_method"),
		Echo:   ctx,
	}

	c := xcontext.SetNoticeContext(ctx.Request().Context(), noticeContext)

	// 有的话就调用全局中间件
	app.Get().Dao().ExpandRecord(notification, []string{"middlewares"}, nil)
	middles := notification.ExpandedAll("middlewares")

	middlewares := []string{"logNotice"}
	for _, middle := range middles {
		middlewares = append(middlewares, middle.GetString("uri"))
	}

	pipeline := openapi.NewPipeline(middlewares...)

	pipeline.Use(a.notice)

	pipeline.Execute(c)

	noticeContext, err = xcontext.GetNoticeContext(c)

	if err != nil {
		return err
	}

	if noticeContext.ResponseType == "string" {
		return ctx.String(http.StatusOK, noticeContext.Response.(string))
	}

	return ctx.JSON(http.StatusOK, noticeContext.Response)
}

func (a *usecase) notice(ctx context.Context, _ openapi.Next) {
	noticeCtx, err := xcontext.GetNoticeContext(ctx)
	if err != nil {
		noticeCtx.Error = err
		xcontext.SetNoticeContext(ctx, noticeCtx)
		return
	}

	if noticeCtx.Url == "" {
		return
	}

	var body io.Reader
	if noticeCtx.Method == http.MethodPost {
		bts, _ := json.Marshal(noticeCtx.Param)
		body = bytes.NewReader(bts)
	}

	rs, err := req(noticeCtx.Url, noticeCtx.Method, body, noticeCtx.Header)
	if err != nil {
		noticeCtx.Error = err
		xcontext.SetNoticeContext(ctx, noticeCtx)
		return
	}

	var resp interface{}

	json.Unmarshal(rs, &resp)

	noticeCtx.Response = resp

	xcontext.SetNoticeContext(ctx, noticeCtx)
}

func (a *usecase) Call(ctx context.Context, req *domain.OpenapiCallReq) (interface{}, error) {
	// 校验app_key下有没有service
	appKey, err := xcontext.GetAppKey(ctx)
	if err != nil {
		return nil, err
	}

	//appKey := "YuoN5mL7mTXPfXhetcza"

	application, err := app.Get().Dao().FindFirstRecordByFilter("apps",
		"state='checked' && app_key={:appKey}",
		dbx.Params{"appKey": appKey})
	if err != nil {
		return nil, err
	}

	srv, err := app.Get().Dao().FindFirstRecordByFilter("services",
		"state='published' && uri={:service}",
		dbx.Params{"service": req.Service})
	if err != nil {
		return nil, err
	}

	appService, err := app.Get().Dao().FindFirstRecordByFilter("app_services",
		"app={:app} && service={:service} && enabled='enabled'",
		dbx.Params{"app": application.Id, "service": srv.Id},
	)

	if err != nil {
		return nil, err
	}

	// 校验servcie下有没有api

	serviceApi, err := app.Get().Dao().FindFirstRecordByFilter("service_apis",
		"service={:service} && state='published' && uri={:uri}",
		dbx.Params{"service": srv.Id, "uri": req.Api},
	)
	if err != nil {
		return nil, err
	}

	config := map[string]string{}
	appService.UnmarshalJSONField("config", &config)

	openapiCtx := &xcontext.OpenapiContext{
		Path:   serviceApi.GetString("path"),
		Domain: srv.GetString("gateway"),
		Config: config,
		Method: serviceApi.GetString("method"),
		Param:  req.Args,
		Async:  serviceApi.GetBool("async"),
	}
	if openapiCtx.Async && req.AsyncUrl != "" {
		openapiCtx.AsyncUrl = req.AsyncUrl
	}

	ctx = xcontext.SetOpenapiContext(ctx, openapiCtx)
	// 有的话就调用全局中间件
	app.Get().Dao().ExpandRecord(srv, []string{"middlewares"}, nil)
	srvMiddles := srv.ExpandedAll("middlewares")

	middlewares := []string{"log"}
	for _, middle := range srvMiddles {
		middlewares = append(middlewares, middle.GetString("uri"))
	}

	app.Get().Dao().ExpandRecord(serviceApi, []string{"middlewares"}, nil)
	apiMiddles := serviceApi.ExpandedAll("middlewares")
	for _, middle := range apiMiddles {
		middlewares = append(middlewares, middle.GetString("uri"))
	}
	pipeline := openapi.NewPipeline(middlewares...)

	pipeline.Use(a.call)

	pipeline.Execute(ctx)

	// 返回结果
	openapiCtx, err = xcontext.GetOpenapiContext(ctx)
	if err != nil {
		return nil, err
	}

	if openapiCtx.Error != nil {
		return nil, openapiCtx.Error
	}

	return openapiCtx.Response, nil
}

func (a *usecase) call(ctx context.Context, _ openapi.Next) {
	openapiCtx, err := xcontext.GetOpenapiContext(ctx)
	if err != nil {
		openapiCtx.Error = err
		xcontext.SetOpenapiContext(ctx, openapiCtx)
		return
	}
	var body io.Reader
	if openapiCtx.Method == http.MethodPost {
		bts, _ := json.Marshal(openapiCtx.Param)
		body = bytes.NewReader(bts)
	}

	if openapiCtx.Async {
		a.pushJob(openapiCtx)
		openapiCtx.Error = nil
		xcontext.SetOpenapiContext(ctx, openapiCtx)
		return
	}

	rs, err := req(openapiCtx.Domain+"/"+strings.Trim(openapiCtx.Path, "/"), openapiCtx.Method, body, openapiCtx.Header)
	if err != nil {
		openapiCtx.Error = err
		xcontext.SetOpenapiContext(ctx, openapiCtx)
		return
	}

	var resp interface{}

	json.Unmarshal(rs, &resp)

	openapiCtx.Response = resp

	xcontext.SetOpenapiContext(ctx, openapiCtx)
}

func (a *usecase) pushJob(ctx *xcontext.OpenapiContext) {
	ants.Submit(func() {
		a.ch <- ctx
	})
}

func (a *usecase) processJob() {
	for openapiCtx := range a.ch {
		var body io.Reader
		if openapiCtx.Method == http.MethodPost {
			bts, _ := json.Marshal(openapiCtx.Param)
			body = bytes.NewReader(bts)
		}

		rs, err := req(openapiCtx.Domain+"/"+strings.Trim(openapiCtx.Path, "/"), openapiCtx.Method, body, openapiCtx.Header)
		if err != nil {
			app.Get().Logger().Error("openapi async request failed", "err", err, "params", openapiCtx)
			return
		}
		if openapiCtx.AsyncUrl == "" {
			return
		}
		asyncRs, err := req(openapiCtx.AsyncUrl, http.MethodPost, bytes.NewReader(rs), map[string]string{"Content-Type": "application/json"})
		if err != nil {
			app.Get().Logger().Error("openapi async request failed: ", err, string(rs))
		}
		app.Get().Logger().Info("openapi async request Succ", "wpt rs", string(rs), "async rs", string(asyncRs), "params", openapiCtx)

	}
}

func req(url string, method string, body io.Reader, head map[string]string) ([]byte, error) {
	timeout := 10000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	// Create an http.Request instance
	req, _ := http.NewRequest(method, url, body)
	for k, v := range head {
		req.Header.Set(k, v)
	}
	// Call the `Do` method, which has a similar interface to the `http.Do` method
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}
