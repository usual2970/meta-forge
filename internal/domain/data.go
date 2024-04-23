package domain

import "context"

type DataListReq struct {
	Table    string `json:"table" query:"table"`
	Page     int    `json:"page" query:"page"`
	PageSize int    `json:"pageSize" query:"pageSize"`
	OrderBy  string `json:"orderBy" query:"orderBy"`
	Filter   string `json:"filter" query:"filter"`
	Params   []any  `json:"params" query:"params"`
}

type DataListResp struct {
	Data         []map[string]interface{} `json:"data"`
	TotalRecords int                      `json:"totalRecords"`
	Page         int                      `json:"page"`
	PageSize     int                      `json:"pageSize"`
}

type IDataUsecase interface {
	List(ctx context.Context, req *DataListReq) (*DataListResp, error)
}

type IDataRepository interface {
	List(ctx context.Context, req *DataListReq) (*DataListResp, error)
}
