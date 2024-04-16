package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/usual2970/meta-forge/internal/routes"
	"github.com/usual2970/meta-forge/internal/util/app"
	"github.com/usual2970/meta-forge/ui"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/pocketbase/pocketbase/core"

	"github.com/usual2970/meta-forge/internal/repository/secret"
	wecomRepository "github.com/usual2970/meta-forge/internal/repository/wecom"
	"github.com/usual2970/meta-forge/internal/usecase/wecom"

	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	_ "github.com/pocketbase/pocketbase/migrations"
)

const mfPath = "/mf/"

func main() {
	os.Setenv("SERPAPI_API_KEY", "0cc261b4c1814a7f2d9ed79df97df998a0232c3059df1733532cce6610f04c42")
	app := app.Get()

	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {

		e.Router.GET(
			strings.TrimRight(mfPath, "/"),
			func(c echo.Context) error {
				return c.Redirect(http.StatusTemporaryRedirect, strings.TrimLeft(mfPath, "/"))
			},
		)
		e.Router.GET(
			mfPath+"*",
			echo.StaticDirectoryHandler(ui.DistDirFS, false),
			installerRedirect(app),
			middleware.Gzip(),
		)

		e.Router.GET("/hello/:name", func(c echo.Context) error {
			name := c.PathParam("name")

			return c.JSON(http.StatusOK, map[string]string{"message": "Hello " + name})
		} /* optional middlewares */)

		routes.Route(e.Router)

		return nil
	})

	app.OnTerminate().Add(func(e *core.TerminateEvent) error {
		routes.UnRegister()

		secretRepo := secret.NewRepository()
		wecomRepo := wecomRepository.NewRepository()
		wc := wecom.NewUsecase(secretRepo, wecomRepo)

		return wc.Exit()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func installerRedirect(app core.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// skip redirect checks for non-root level index.html requests
			path := c.Request().URL.Path
			if path != mfPath && path != mfPath+"index.html" {
				return next(c)
			}

			return next(c)
		}
	}
}
