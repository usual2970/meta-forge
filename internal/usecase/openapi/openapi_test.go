package openapi

import (
	"testing"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/tests"
)

const testDataDir = "/Users/liuxuanyao/work/github.com/usual2970/meta-forge/pb_data"

func TestRelation(t *testing.T) {
	app, err := tests.NewTestApp(testDataDir)
	if err != nil {
		t.Fatal(err)
	}
	defer app.Cleanup()

	rs, err := app.Dao().FindFirstRecordByFilter("app_services", "app={:id}", dbx.Params{"id": "ebytkqacnrfz6ru"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rs)
}
