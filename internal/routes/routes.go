package routes

import (
	"github.com/usual2970/meta-forge/internal/controller/data"
	"github.com/usual2970/meta-forge/internal/controller/systemsettings"

	systemSettingsRepository "github.com/usual2970/meta-forge/internal/repository/system_settings"
	systemSettingsUC "github.com/usual2970/meta-forge/internal/usecase/systemsettings"

	dataRepository "github.com/usual2970/meta-forge/internal/repository/data"
	dataUC "github.com/usual2970/meta-forge/internal/usecase/data"

	"github.com/labstack/echo/v5"
)

func Route(echo *echo.Echo) {

	systemSettingsRepo := systemSettingsRepository.NewRepository()
	dataRepo := dataRepository.NewRepository()

	systemSettingsUc := systemSettingsUC.NewUsecase(systemSettingsRepo)
	dataUc := dataUC.NewUsecase(dataRepo)

	group := echo.Group("/api/v1")

	systemsettings.Register(group, systemSettingsUc)
	data.Register(group, dataUc)
}
