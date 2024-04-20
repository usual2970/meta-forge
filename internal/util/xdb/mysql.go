package xdb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sort"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/util/helper"
)

type Mysql struct {
	db *sql.DB
}

type mysqlField struct {
	TableName              string  `json:"tableName"`
	ColumnName             string  `json:"columnName"`
	DataType               string  `json:"dataType"`
	IsNullable             string  `json:"isNullable"`
	ColumnKey              string  `json:"columnKey"`
	ColumnDefault          *string `json:"columnDefault"`
	Extra                  string  `json:"extra"`
	OrdinalPosition        int     `json:"ordinalPosition"`
	CharacterMaximumLength *int    `json:"characterMaximumLength"`
	ColumnType             string  `json:"columnType"`
}

func (m mysqlField) getType() string {
	switch m.DataType {
	case "varchar", "text", "json":
		return FieldTypeString
	case "datetime":
		return FieldTypeDate
	case "int", "tinyint", "bigint":
		return FieldTypeNumber
	case "enum":
		return FieldTypeEnum

	default:
		return FieldTypeString
	}
}

func (m mysqlField) isRequired() bool {
	return m.IsNullable == "NO"
}

func (m mysqlField) isId() bool {
	return m.ColumnKey == "PRI"
}

func (m mysqlField) length() int {
	if m.CharacterMaximumLength == nil {
		return 0
	}
	return *m.CharacterMaximumLength
}

func (m mysqlField) enumeration() []string {
	if m.DataType != "enum" {
		return nil
	}

	columnType := strings.TrimLeft(m.ColumnType, "enum(")

	columnType = strings.TrimRight(columnType, ")")

	ts := strings.Split(columnType, ",")

	rs := make([]string, 0, len(ts))

	for _, t := range ts {
		rs = append(rs, strings.Trim(t, "'"))
	}
	return rs
}

type mysqlFields []mysqlField

func (m mysqlFields) Len() int {
	return len(m)
}
func (m mysqlFields) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m mysqlFields) Less(i, j int) bool {
	return m[i].OrdinalPosition < m[j].OrdinalPosition
}

type mysqlIndex struct {
	TableCatalog string
	TableSchema  string
	TableName    string
	NonUnique    int
	IndexSchema  string
	IndexName    string
	SeqInIndex   int
	ColumnName   string
}

type mysqlIndexes []mysqlIndex

func (m mysqlIndexes) Len() int {
	return len(m)
}
func (m mysqlIndexes) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m mysqlIndexes) Less(i, j int) bool {
	return m[i].SeqInIndex < m[j].SeqInIndex
}

func InitialDbMysql(ctx context.Context, req *domain.InitializeReq) (XDB, error) {
	if req.User == "" || req.Password == "" || req.Host == "" || req.Database == "" {
		return nil, errors.New("invalid mysql config")
	}

	if req.Port == "" {
		req.Port = "3306"
	}

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", req.User, req.Password, req.Host, req.Port, req.Database)

	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	return &Mysql{
		db: db,
	}, nil
}

func (m *Mysql) GetSchemas() ([]TableSchema, error) {

	rows, err := m.DB().Query("SELECT TABLE_NAME, COLUMN_NAME, DATA_TYPE, IS_NULLABLE, COLUMN_KEY, COLUMN_DEFAULT, EXTRA,ORDINAL_POSITION,CHARACTER_MAXIMUM_LENGTH,COLUMN_TYPE FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = DATABASE()")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fields := make([]mysqlField, 0)
	for rows.Next() {
		var r mysqlField
		err := rows.Scan(
			&r.TableName,
			&r.ColumnName,
			&r.DataType,
			&r.IsNullable,
			&r.ColumnKey,
			&r.ColumnDefault,
			&r.Extra,
			&r.OrdinalPosition,
			&r.CharacterMaximumLength,
			&r.ColumnType,
		)
		if err != nil {
			return nil, err
		}
		fields = append(fields, r)
	}

	fieldsMap := helper.GroupBy(fields, func(mf mysqlField) string {
		return mf.TableName
	})

	// 从当前数据库中获取所有的索引信息
	indexRows, err := m.DB().Query("SELECT TABLE_CATALOG, TABLE_SCHEMA, TABLE_NAME, NON_UNIQUE, INDEX_SCHEMA, INDEX_NAME, SEQ_IN_INDEX, COLUMN_NAME FROM INFORMATION_SCHEMA.STATISTICS WHERE TABLE_SCHEMA = DATABASE()")
	if err != nil {
		return nil, err
	}
	defer indexRows.Close()

	indexes := make([]mysqlIndex, 0)
	for indexRows.Next() {
		var r mysqlIndex
		err := indexRows.Scan(
			&r.TableCatalog,
			&r.TableSchema,
			&r.TableName,
			&r.NonUnique,
			&r.IndexSchema,
			&r.IndexName,
			&r.SeqInIndex,
			&r.ColumnName,
		)
		if err != nil {
			return nil, err
		}
		indexes = append(indexes, r)
	}

	indexesMap := helper.GroupBy(indexes, func(mi mysqlIndex) string {
		return mi.TableName
	})

	rs := make([]TableSchema, 0)
	for tableName, fields := range fieldsMap {
		tableSchema := TableSchema{
			Name: tableName,
		}

		sort.Sort(mysqlFields(fields))

		for _, field := range fields {
			tableSchema.Fields = append(tableSchema.Fields, TableSchemaField{
				Name:        field.ColumnName,
				Type:        field.getType(),
				IsRequired:  field.isRequired(),
				IsId:        field.isId(),
				Length:      field.length(),
				Enumeration: field.enumeration(),
			})
		}

		if indexes, ok := indexesMap[tableName]; ok {
			m := helper.GroupBy(indexes, func(mi mysqlIndex) int {
				return mi.NonUnique
			})

			// 获取所有unique索引
			uniqueFields, ok := m[0]
			if ok {
				uniqueFieldsMap := helper.GroupBy(uniqueFields, func(mi mysqlIndex) string {
					return mi.IndexName
				})

				for _, uniqueFields := range uniqueFieldsMap {
					sort.Sort(mysqlIndexes(uniqueFields))

					var uniqueFieldNames []string
					for _, uniqueField := range uniqueFields {
						uniqueFieldNames = append(uniqueFieldNames, uniqueField.ColumnName)
					}

					tableSchema.UniqueFields = append(tableSchema.UniqueFields, uniqueFieldNames)
				}

			}

		}

		rs = append(rs, tableSchema)
	}

	return rs, nil
}

func (m *Mysql) DB() *sql.DB {

	return m.db
}

func (m *Mysql) Close() error {

	return m.db.Close()
}
