package xdb

import (
	"context"
	"database/sql"
	"errors"
	"sync"

	"github.com/usual2970/meta-forge/internal/domain"
)

const (
	FieldTypeString = "string"
	FieldTypeNumber = "number"
	FieldTypeDate   = "date"
	FieldTypeEnum   = "enum"
)

type XDB interface {
	GetSchemas() ([]domain.TableSchema, error)
	DB() *sql.DB
	Close() error
}

var xdb XDB
var xdbOnce sync.Once

func DB(ctx context.Context, req *domain.InitializeReq) XDB {
	xdbOnce.Do(func() {
		var err error
		xdb, err = InitialDb(ctx, req)
		if err != nil {
			panic(err)
		}

	})
	return xdb
}

func InitialDb(ctx context.Context, req *domain.InitializeReq) (XDB, error) {
	switch req.Kind {
	case domain.DbKindMysql:
		return InitialDbMysql(ctx, req)
	default:
		return nil, errors.New("initial database failed, " + req.Kind + " is not supported")
	}
}
