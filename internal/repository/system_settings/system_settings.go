package systemsettings

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/util/app"
)

type repository struct {
}

func NewRepository() domain.ISystemSettingsRepository {

	return &repository{}
}

func (r *repository) GetByType(ctx context.Context, t string) (map[string]interface{}, error) {
	records, err := app.GetDao().FindRecordsByFilter("mf_system_settings", "type='"+t+"'", "-created", 1000, 0)
	if err != nil {
		return nil, err
	}

	rs := make(map[string]interface{})

	for _, record := range records {
		var data map[string]interface{}
		if err := record.UnmarshalJSONField("data", &data); err != nil {
			continue
		}
		res, ok := data["value"]
		if !ok {
			continue
		}
		rs[record.GetString("uri")] = res
	}
	return rs, nil
}

func (r *repository) BatchGet(ctx context.Context, keys []string) (map[string]interface{}, error) {

	filter := []string{}
	for _, key := range keys {
		filter = append(filter, fmt.Sprintf("uri = '%s'", key))
	}

	records, err := app.GetDao().FindRecordsByFilter("mf_system_settings", strings.Join(filter, " || "), "-created", 10, 0)
	if err != nil {
		return nil, err
	}

	rs := make(map[string]interface{})

	for _, record := range records {
		var data map[string]interface{}
		if err := record.UnmarshalJSONField("data", &data); err != nil {
			continue
		}
		res, ok := data["value"]
		if !ok {
			continue
		}
		rs[record.GetString("uri")] = res
	}
	return rs, nil
}

func (r *repository) BatchSave(ctx context.Context, settings []domain.SystemSetting) error {
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

func (r *repository) Save(ctx context.Context, req *domain.SystemSettingSaveReq) error {

	collection, err := app.GetDao().FindCollectionByNameOrId("mf_system_settings")
	if err != nil {
		return err
	}
	record, err := app.GetDao().FindFirstRecordByFilter("mf_system_settings", "uri='"+req.Uri+"'")
	if err != nil {
		record = models.NewRecord(collection)
	}

	form := forms.NewRecordUpsert(app.Get(), record)

	form.LoadData(
		map[string]any{
			"id":  record.Id,
			"uri": req.Uri,
			"data": map[string]interface{}{
				"value": req.Data,
			},
			"type": req.Type,
		},
	)

	if err := form.Submit(); err != nil {
		return err
	}

	app.Get().Store().Remove(req.Uri)
	return nil
}

func (r *repository) GetSchemas(ctx context.Context) (map[string]domain.TableSchema, error) {
	key := "schemas"
	skey := "@" + key
	data := app.Get().Store().Get(skey)
	if data != nil {
		return data.(map[string]domain.TableSchema), nil
	}
	schemas, err := r.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	bts, _ := json.Marshal(schemas)

	rs := make([]domain.TableSchema, 0)

	err = json.Unmarshal(bts, &rs)
	if err != nil {
		return nil, fmt.Errorf("invalid schemas:%w", err)
	}

	rsMap := make(map[string]domain.TableSchema)

	for _, item := range rs {
		rsMap[item.Name] = item
	}

	app.Get().Store().Set(skey, rsMap)

	return rsMap, nil
}

func (r *repository) Get(ctx context.Context, key string) (interface{}, error) {
	data := app.Get().Store().Get(key)
	if data != nil {
		return data, nil
	}

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
		app.Get().Store().Set(key, rs["value"])
		return res, nil
	}

	return nil, errors.New("not found")
}
