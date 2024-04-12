package middleware

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/util/app"
	"github.com/usual2970/meta-forge/internal/util/xcontext"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/pocketbase/dbx"
)

func AuthToken() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(c echo.Context, user, password string) (bool, error) {
		ts := c.Request().Header.Get("X-Ts")
		if user == "" || password == "" || ts == "" {
			return false, nil
		}

		// 获取 app
		application, err := app.Get().Dao().FindFirstRecordByFilter("apps",
			"state={:checked} && app_key={:appKey}",
			dbx.Params{"appKey": user, "checked": domain.OpenapiAppChecked})
		if err != nil {
			return false, nil
		}

		appSecret := application.GetString("app_secret")

		path := "/" + strings.Trim(c.Request().URL.Path, "/")

		ctx := c.Request().Context()

		ctx = xcontext.SetAppKey(ctx, user)

		c.SetRequest(c.Request().WithContext(ctx))

		return checkToken(user, appSecret, path, ts, password), nil
	})
}

func checkToken(appKey, appSecret, path, ts, requestToken string) bool {
	token := generateToken(appKey, appSecret, path, ts)

	return token == requestToken
}

func generateToken(appKey, appSecret, path, ts string) string {
	raw := fmt.Sprintf("%s-%s-%s-%s", appKey, path, ts, appSecret)
	hash := sha1.New()

	hash.Write([]byte(raw))

	return hex.EncodeToString(hash.Sum(nil))
}
