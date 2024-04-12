package routes

import (
	"context"

	"github.com/usual2970/meta-forge/internal/domain"
	botUC "github.com/usual2970/meta-forge/internal/usecase/bot"
)

var botUc domain.IBotUsecase

func Register() error {
	var err error
	botUc, err = botUC.New()

	return err
}

func UnRegister() {
	if botUc == nil {
		return
	}
	botUc.Stop(context.Background())
}
