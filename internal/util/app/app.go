package app

import (
	"sync"

	"github.com/pocketbase/pocketbase"
)

var app *pocketbase.PocketBase
var appOnce sync.Once

const testDataDir = "/Users/liuxuanyao/work/metaforge/pb_data"

// func GetTest() *pocketbase.PocketBase {
// 	appOnce.Do(func() {
// 		appTest, err := tests.NewTestApp(testDataDir)
// 		if err != nil {
// 			panic(err)
// 		}
// 		app = appTest

// 	})
// 	return app
// }

func Get() *pocketbase.PocketBase {
	appOnce.Do(func() {
		app = pocketbase.New()

	})
	return app
}
