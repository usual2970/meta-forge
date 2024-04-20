package systemsettings

import (
	"context"

	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/util/xdb"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type usecase struct {
	repo domain.ISystemSettingsRepository
}

func NewUsecase(repo domain.ISystemSettingsRepository) domain.ISystemSettingsUsecase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) Get(ctx context.Context, key string) (map[string]interface{}, error) {
	return nil, nil
}
func (u *usecase) Initialize(ctx context.Context, req *domain.InitializeReq) error {

	// 先连接数据库
	db, err := xdb.InitialDb(ctx, req)
	if err != nil {
		return err
	}

	// 获取数据库中的表结构
	schemas, err := db.GetSchemas()
	if err != nil {
		return err
	}

	// 保存起来返回
	settings := []domain.SystemSetting{
		{
			Data:        req,
			Uri:         "dbconfig",
			Description: "数据库配置",
		},
		{
			Uri: "@hasInitialized",
			Data: map[string]int{
				"value": 1,
			},
			Description: "初始化标识",
		},
		{
			Uri:         "schemas",
			Data:        schemas,
			Description: "数据库表结构",
		},
	}

	if err := u.repo.BatchSave(ctx, settings); err != nil {
		return err
	}
	return nil

}
