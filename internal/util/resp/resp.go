package resp

import (
	"net/http"

	"github.com/usual2970/meta-forge/internal/domain/constant"

	"github.com/labstack/echo/v5"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Succ(e echo.Context, data interface{}) error {
	rs := &Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
	return e.JSON(http.StatusOK, rs)
}

func Err(e echo.Context, err error) error {

	xerr, ok := err.(*constant.XError)
	code := 100
	if ok {
		code = xerr.GetCode()
	}

	rs := &Response{
		Code: code,
		Msg:  err.Error(),
		Data: nil,
	}
	return e.JSON(http.StatusOK, rs)
}

func WecomVerifyUrl(e echo.Context, msg string) error {

	return e.String(http.StatusOK, msg)
}
