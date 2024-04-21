package systemsettings

import (
	"context"
	"errors"

	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/util/app"
)

type SystemSettingsRepository struct {
}

func NewRepository() domain.ISystemSettingsRepository {
	return &SystemSettingsRepository{}
}

func (r *SystemSettingsRepository) BatchSave(ctx context.Context, settings []domain.SystemSetting) error {
	collection, err := app.GetDao().FindCollectionByNameOrId("mf_system_settings")
	if err != nil {
		return err
	}

	if err := app.GetDao().RunInTransaction(func(txDao *daos.Dao) error {
		for _, setting := range settings {
			record := models.NewRecord(collection)
			record.Set("uri", setting.Uri)
			record.Set("description", setting.Description)
			record.Set("data", map[string]interface{}{
				"value": setting.Data,
			})

			if err := txDao.SaveRecord(record); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil

}

func (r *SystemSettingsRepository) Get(ctx context.Context, key string) (interface{}, error) {

	record, err := app.GetDao().FindFirstRecordByFilter("mf_system_settings",
		"uri='"+key+"'",
	)
	if err != nil {
		return nil, err
	}
	var rs map[string]interface{}
	if err := record.UnmarshalJSONField("data", &rs); err != nil {
		return nil, err
	}
	if res, ok := rs["value"]; ok {
		return res, nil
	}
	return nil, errors.New("not found")
}
