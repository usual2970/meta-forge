package domain

import "context"

type ISystemSettingsUsecase interface {
	Get(ctx context.Context, key string) (map[string]interface{}, error)
	Initialize(ctx context.Context) error
}

type ISystemSettingsRepository interface {
	Get(ctx context.Context, key string) (map[string]interface{}, error)
}
