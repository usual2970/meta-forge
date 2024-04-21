package routes

import (
	"github.com/usual2970/meta-forge/internal/controller/systemsettings"

	systemSettingsUC "github.com/usual2970/meta-forge/internal/usecase/systemsettings"

	systemSettingsRepository "github.com/usual2970/meta-forge/internal/repository/system_settings"

	"github.com/labstack/echo/v5"
)

func Route(echo *echo.Echo) {

	systemSettingsRepo := systemSettingsRepository.NewRepository()

	systemSettingsUc := systemSettingsUC.NewUsecase(systemSettingsRepo)
	group := echo.Group("/api/v1")
	systemsettings.Register(group, systemSettingsUc)
}
