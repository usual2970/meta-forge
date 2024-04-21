package app

import (
	"sync"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/tests"
)

var app core.App
var appOnce sync.Once

const testDataDir = "/Users/liuxuanyao/work/metaforge/pb_data"

func GetTest() core.App {
	appOnce.Do(func() {
		appTest, err := tests.NewTestApp(testDataDir)
		if err != nil {
			panic(err)
		}
		app = appTest

	})
	return app
}

func Get() *pocketbase.PocketBase {
	rs := get()
	return rs.(*pocketbase.PocketBase)
}

func get() core.App {
	appOnce.Do(func() {
		app = pocketbase.New()

	})
	return app
}

func GetDao() *daos.Dao {
	return get().Dao()
}
