package domain

import "context"

const (
	DbKindMysql  = "mysql"
	DbKindSqlite = "sqlite"
)

type InitializeReq struct {
	Kind     string `json:"kind"`
	File     string `json:"file"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type ISystemSettingsUsecase interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Initialize(ctx context.Context, req *InitializeReq) error
}

type SystemSetting struct {
	Meta
	Uri         string
	Description string
	Data        interface{}
}

type ISystemSettingsRepository interface {
	Get(ctx context.Context, key string) (interface{}, error)
	BatchSave(ctx context.Context, settings []SystemSetting) error
}
