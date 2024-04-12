package domain

import "context"

type IBotUsecase interface {
	Process(ctx context.Context) error

	Stop(ctx context.Context)
}

