package bot

import (
	"context"

	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/util/app"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type usecase struct {
	ch     tgbotapi.UpdatesChannel
	bot    *tgbotapi.BotAPI
	cancel context.CancelFunc
}

func New() (domain.IBotUsecase, error) {

	bot, err := tgbotapi.NewBotAPI("7083555675:AAFL9qFVC1h8nSFuBColrm7renrO0YSh1rQ")
	if err != nil {
		return nil, err
	}

	rs := &usecase{
		bot: bot,
	}

	rs.ch = bot.GetUpdatesChan(tgbotapi.NewUpdate(0))

	ctx, cancel := context.WithCancel(context.Background())
	rs.cancel = cancel

	for i := 0; i < 10; i++ {
		go func() {
			rs.Process(ctx)
		}()
	}

	return rs, nil
}

func (u *usecase) Process(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case update := <-u.ch:

			session := GetSession(update)
			reply, err := session.Process(ctx, update)
			if err != nil {
				app.Get().Logger().Info("process update error: %v", err)
				continue
			}
			if reply != nil {
				u.bot.Send(reply)
			}

		}
	}

}

func (u *usecase) Exit() error {
	u.cancel()
	return nil
}

func (u *usecase) Stop(ctx context.Context) {
	app.Get().Logger().Info("stop bot")
	u.cancel()
}
