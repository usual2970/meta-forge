package xdb

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/usual2970/meta-forge/internal/domain"
)

func MustDb() *sql.DB {
	url := "root:root@tcp(127.0.0.1:13306)/jichedi"

	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	return db
}

func TestMysql_GetSchemas(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domain.TableSchema
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				db: MustDb(),
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mysql{
				db: tt.fields.db,
			}
			got, err := m.GetSchemas()
			if (err != nil) != tt.wantErr {
				t.Errorf("Mysql.GetSchemas() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mysql.GetSchemas() = %v, want %v", got, tt.want)
			}
		})
	}
}
