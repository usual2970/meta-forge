package xdb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/usual2970/meta-forge/internal/domain"
)

const (
	FieldTypeString = "string"
	FieldTypeNumber = "number"
	FieldTypeDate   = "date"
	FieldTypeEnum   = "enum"
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

type XDB interface {
	GetSchemas() ([]TableSchema, error)
	DB() *sql.DB
	Close() error
}

func InitialDb(ctx context.Context, req *domain.InitializeReq) (XDB, error) {
	switch req.Kind {
	case domain.DbKindMysql:
		return InitialDbMysql(ctx, req)
	default:
		return nil, errors.New("initial database failed, " + req.Kind + " is not supported")
	}
}
