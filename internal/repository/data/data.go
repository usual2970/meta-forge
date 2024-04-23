package data

import (
	"context"
	osql "database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/usual2970/meta-forge/internal/domain"
	systemsettings "github.com/usual2970/meta-forge/internal/repository/system_settings"

	"github.com/usual2970/meta-forge/internal/util/xdb"
)

type repository struct {
	xdb         xdb.XDB
	settingRepo domain.ISystemSettingsRepository
}

func NewRepository() domain.IDataRepository {

	settingRepo := systemsettings.NewRepository()
	ctx := context.Background()
	config, err := settingRepo.Get(ctx, "dbconfig")
	if err != nil {
		panic(err)
	}

	req := &domain.InitializeReq{}

	byts, _ := json.Marshal(config)

	json.Unmarshal(byts, req)

	db := xdb.DB(ctx, req)

	return &repository{
		settingRepo: settingRepo,
		xdb:         db,
	}
}

func (r *repository) List(ctx context.Context, req *domain.DataListReq) (*domain.DataListResp, error) {
	count, err := r.count(ctx, req)
	if err != nil {
		return nil, err
	}

	schemas, err := r.settingRepo.GetSchemas(ctx)
	if err != nil {
		return nil, err
	}

	schema, ok := schemas[req.Table]
	if !ok {
		return nil, errors.New("table not found")
	}

	fields := fields(schema)

	sql := fmt.Sprintf("select %s from %s where %s order by %s limit %s",
		buildSelect(schema),
		buildFrom(req.Table),
		buildWhere(req.Filter),
		buildOrderBy(req.OrderBy),
		buildLimit(req),
	)

	rows, err := r.xdb.DB().Query(sql, req.Params...)
	if err != nil {
		return nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]osql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var result []map[string]interface{}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		entry := make(map[string]interface{})
		for i, col := range columns {
			entry[col] = toValue(values[i], fields[col])
		}
		result = append(result, entry)
	}
	return &domain.DataListResp{
		Data:         result,
		Page:         page(req),
		PageSize:     pageSize(req),
		TotalRecords: count,
	}, nil
}

func (r *repository) count(ctx context.Context, req *domain.DataListReq) (int, error) {
	sql := fmt.Sprintf("select count(*) from %s where %s",
		buildFrom(req.Table),
		buildWhere(req.Filter),
	)

	var totalRecords int
	if err := r.xdb.DB().QueryRowContext(ctx, sql, req.Params...).Scan(&totalRecords); err != nil {
		return 0, err
	}

	return totalRecords, nil
}

func toValue(bts []byte, field domain.TableSchemaField) interface{} {
	if bts == nil {
		return nil
	}

	switch field.Type {
	case xdb.FieldTypeString, xdb.FieldTypeEnum:
		return string(bts)
	case xdb.FieldTypeNumber:
		f, err := strconv.ParseFloat(string(bts), 64)
		if err != nil {
			return nil
		}
		return f
	case xdb.FieldTypeDate:
		t, err := time.Parse("2006-01-02", string(bts))
		if err != nil {
			return nil
		}
		return t

	default:
		return string(bts)
	}

}

func page(req *domain.DataListReq) int {
	page := 1
	if req.Page > 0 {
		page = req.Page
	}
	return page
}

func pageSize(req *domain.DataListReq) int {
	pageSize := 10
	if req.PageSize > 0 {
		pageSize = req.PageSize
	}
	return pageSize
}

func fields(schema domain.TableSchema) map[string]domain.TableSchemaField {
	fields := make(map[string]domain.TableSchemaField)
	for _, field := range schema.Fields {
		fields[field.Name] = field
	}
	return fields
}

func buildFrom(table string) string {
	return fmt.Sprintf("`%s`", table)
}

func buildOrderBy(orderBy string) string {
	if orderBy == "" {
		return "id asc"
	}
	return orderBy
}

func buildWhere(filter string) string {
	if filter == "" {
		return "1=1"
	}
	filter = strings.ReplaceAll(filter, "&&", "and")
	filter = strings.ReplaceAll(filter, "||", "or")
	return filter
}

func buildLimit(req *domain.DataListReq) string {
	page := page(req)
	pageSize := pageSize(req)
	offset := (page - 1) * pageSize

	return fmt.Sprintf("%d, %d", offset, pageSize)
}

func buildSelect(schema domain.TableSchema) string {
	fields := []string{}
	for _, field := range schema.Fields {

		fields = append(fields, "`"+field.Name+"`")

	}

	return strings.Join(fields, ",")
}
