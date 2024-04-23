package domain

import (
	"context"
)

const (
	DbKindMysql  = "mysql"
	DbKindSqlite = "sqlite"
)

type TableSchema struct {
	Name         string                `json:"name"`
	Fields       []TableSchemaField    `json:"fields"`
	Relations    []TableSchemaRelation `json:"relations"`
	UniqueFields [][]string            `json:"uniqueFields"`
}

type TableSchemaRelation struct {
	Name               string `json:"name"`
	FieldName          string `json:"fieldName"`
	ReferenceTable     string `json:"referenceTable"`
	ReferenceFieldName string `json:"referenceFieldName"`
}

type TableSchemaField struct {
	Name        string   `json:"name"`
	IsRequired  bool     `json:"isRequired"`
	IsId        bool     `json:"isId"`
	Type        string   `json:"type"`
	Enumeration []string `json:"enumeration"`
	Length      int      `json:"length"`
}

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
	BatchGet(ctx context.Context, keys []string) (map[string]interface{}, error)
}

type SystemSetting struct {
	Meta
	Uri         string
	Description string
	Data        interface{}
}

type ISystemSettingsRepository interface {
	Get(ctx context.Context, key string) (interface{}, error)
	GetSchemas(ctx context.Context) (map[string]TableSchema, error)
	BatchSave(ctx context.Context, settings []SystemSetting) error
	BatchGet(ctx context.Context, keys []string) (map[string]interface{}, error)
}
