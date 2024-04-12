package main

import (
	"log"
	"net/http"
	"os"

	"github.com/usual2970/meta-forge/internal/routes"
	"github.com/usual2970/meta-forge/internal/util/app"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"

	"github.com/usual2970/meta-forge/internal/repository/secret"
	wecomRepository "github.com/usual2970/meta-forge/internal/repository/wecom"
	"github.com/usual2970/meta-forge/internal/usecase/wecom"
)

func main() {
	os.Setenv("SERPAPI_API_KEY", "0cc261b4c1814a7f2d9ed79df97df998a0232c3059df1733532cce6610f04c42")
	app := app.Get()
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/hello/:name", func(c echo.Context) error {
			name := c.PathParam("name")

			return c.JSON(http.StatusOK, map[string]string{"message": "Hello " + name})
		} /* optional middlewares */)

		routes.Route(e.Router)

		// if err := routes.Register(); err != nil {
		// 	return err
		// }
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
