package openapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/usual2970/meta-forge/internal/util/hash"
	"github.com/usual2970/meta-forge/internal/util/xcontext"
)

func PreWps(ctx context.Context, next Next) {
	openapiContext, err := xcontext.GetOpenapiContext(ctx)
	if err != nil {
		openapiContext = &xcontext.OpenapiContext{
			Error: err,
		}
		xcontext.SetOpenapiContext(ctx, openapiContext)
		return
	}

	location, err := time.LoadLocation("UTC")
	if err != nil {
		openapiContext.Error = err
		return
	}

	appId, ok := openapiContext.Config["app_id"]
	if !ok {
		openapiContext.Error = errors.New("app_id is required")
		return
	}

	appSecret, ok := openapiContext.Config["app_secret"]
	if !ok {
		openapiContext.Error = errors.New("app_secret is required")
		return
	}

	// 创建一个带有纽约时区的当前时间
	currentTime := time.Now().In(location)
	date := currentTime.Format("Mon, 02 Jan 2006 15:04:05 GMT")
	md5 := ""
	param := openapiContext.Param
	if openapiContext.Method == http.MethodGet {
		if len(openapiContext.Param) == 0 {
			md5 = hash.Md5(openapiContext.Path)
		} else {

			docType, ok := param["docType"]
			if ok {
				delete(param, "docType")
				openapiContext.Path = strings.Replace(openapiContext.Path, ":office_type", docType.(string), -1)
			}

			taskId, ok := param["taskId"]
			if !ok {
				openapiContext.Error = errors.New("missing taskId")
				xcontext.SetOpenapiContext(ctx, openapiContext)
				return
			}
			delete(param, "taskId")
			openapiContext.Path = strings.Replace(openapiContext.Path, ":task_id", taskId.(string), -1)

			openapiContext.Param = param
			md5 = hash.Md5(openapiContext.Path)
		}

	} else {
		bts, _ := json.Marshal(openapiContext.Param)
		md5 = hash.Md5(string(bts))
	}

	contentType := "application/json"

	signature := hash.Sha1(appSecret + md5 + contentType + date)

	if len(openapiContext.Header) == 0 {
		openapiContext.Header = make(map[string]string)
	}
	header := openapiContext.Header

	header["Date"] = date
	header["Content-Type"] = contentType
	header["Content-Md5"] = md5
	header["Authorization"] = fmt.Sprintf("WPS-2:%s:%s", appId, signature)

	xcontext.SetOpenapiContext(ctx, openapiContext)

	next()
}
