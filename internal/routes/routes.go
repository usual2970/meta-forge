package routes

import (
	"github.com/usual2970/meta-forge/internal/controller/ai"
	"github.com/usual2970/meta-forge/internal/controller/mingdao"
	"github.com/usual2970/meta-forge/internal/controller/openapi"
	"github.com/usual2970/meta-forge/internal/controller/wecom"
	"github.com/usual2970/meta-forge/internal/controller/wenxin"
	wenxinUC "github.com/usual2970/meta-forge/internal/usecase/wenxin"

	aiUC "github.com/usual2970/meta-forge/internal/usecase/ai"
	mingdaoUC "github.com/usual2970/meta-forge/internal/usecase/mingdao"
	openapiUC "github.com/usual2970/meta-forge/internal/usecase/openapi"
	wecomUC "github.com/usual2970/meta-forge/internal/usecase/wecom"

	secretRepository "github.com/usual2970/meta-forge/internal/repository/secret"
	wecomRepository "github.com/usual2970/meta-forge/internal/repository/wecom"

	"github.com/labstack/echo/v5"
)

func Route(echo *echo.Echo) {

	secretRepo := secretRepository.NewRepository()
	wecomRepo := wecomRepository.NewRepository()

	wenxinUc := wenxinUC.NewUsecase()

	openapiUc := openapiUC.NewUsecase()

	aiUc := aiUC.NewUsecase(secretRepo)

	mingdaoUc := mingdaoUC.NewUsecase()

	wecomUc := wecomUC.NewUsecase(secretRepo, wecomRepo)

	wenxin.Register(echo, wenxinUc)

	openapi.Register(echo, openapiUc)

	ai.Register(echo, aiUc)

	mingdao.Register(echo, mingdaoUc)

	wecom.Register(echo, wecomUc)
}
