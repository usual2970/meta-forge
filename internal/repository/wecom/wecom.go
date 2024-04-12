package wecom

import (
	"context"
	"sync"

	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/util/app"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/filesystem"
)

var once sync.Once
var instance domain.IWecomRepository

type repository struct{}

func NewRepository() domain.IWecomRepository {
	once.Do(func() {
		instance = &repository{}
	})
	return instance
}

func (r *repository) GetMedia(ctx context.Context, filter string) (*domain.RxMedia, error) {

	record, err := app.Get().Dao().FindFirstRecordByFilter("rx_media",
		filter,
	)
	if err != nil {
		return nil, err
	}

	rs := &domain.RxMedia{

		Meta: domain.Meta{
			Id:      record.Id,
			Created: record.GetTime("created"),
			Updated: record.GetTime("updated"),
		},
		MediaId: record.GetString("media_id"),
		File:    app.Get().Settings().Meta.AppUrl + "/api/files/" + record.BaseFilesPath() + "/" + record.GetString("file"),
	}

	return rs, nil
}

func (r *repository) SaveMedia(ctx context.Context, data *domain.RxMedia) error {
	collection, err := app.Get().Dao().FindCollectionByNameOrId("rx_media")
	if err != nil {
		return err
	}

	record := models.NewRecord(collection)

	form := forms.NewRecordUpsert(app.Get(), record)

	form.LoadData(map[string]any{
		"media_id": data.MediaId,
		"id":       data.Id,
	})

	// manually upload file(s)
	f1, _ := filesystem.NewFileFromUrl(ctx, data.File)

	form.AddFiles("file", f1)

	return form.Submit()

}

func (r *repository) Get(ctx context.Context, filter string) (*domain.RxSupportStuff, error) {

	record, err := app.Get().Dao().FindFirstRecordByFilter("rx_support_stuff",
		filter,
	)
	if err != nil {
		return nil, err
	}

	ext := make(map[string]string)

	record.UnmarshalJSONField("ext", &ext)
	meta := domain.Meta{
		Id:      record.Id,
		Created: record.GetTime("created"),
		Updated: record.GetTime("updated"),
	}
	rs := &domain.RxSupportStuff{

		OpenKfId:   record.GetString("open_kf_id"),
		NextCursor: record.GetString("next_cursor"),

		Meta: meta,
	}

	return rs, nil
}

func (r *repository) SetNextCursor(ctx context.Context, data *domain.RxSupportStuff) error {
	var record *models.Record
	if data.Id == "" {
		collection, _ := app.Get().Dao().FindCollectionByNameOrId("rx_support_stuff")
		record = models.NewRecord(collection)
	} else {
		var err error
		record, err = app.Get().Dao().FindFirstRecordByFilter("rx_support_stuff",
			"id = {:id}",
			dbx.Params{"id": data.Id},
		)
		if err != nil {
			return err
		}
	}

	record.Set("next_cursor", data.NextCursor)
	record.Set("open_kf_id", data.OpenKfId)
	return app.Get().Dao().SaveRecord(record)
}
