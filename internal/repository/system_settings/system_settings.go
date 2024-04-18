package systemsettings

import (
	"context"

	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/util/app"
)

type SystemSettingsRepository struct {
}

func NewRepository() domain.ISystemSettingsRepository {
	return &SystemSettingsRepository{}
}

func (r *SystemSettingsRepository) Get(ctx context.Context, key string) (map[string]interface{}, error) {

	record, err := app.Get().Dao().FindFirstRecordByFilter("mf_system_settings",
		"uri='"+key+"'",
	)
	if err != nil {
		return nil, err
	}
	var rs map[string]interface{}
	if err := record.UnmarshalJSONField("data", &rs); err != nil {
		return nil, err
	}
	return rs, nil
}
