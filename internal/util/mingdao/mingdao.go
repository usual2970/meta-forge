package mingdao

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/usual2970/meta-forge/internal/domain"
	xhttp "github.com/usual2970/meta-forge/internal/util/http"
	"github.com/usual2970/meta-forge/internal/util/str"
)

type mingdaoResp struct {
	Data      interface{} `json:"data"`
	Success   bool        `json:"success"`
	ErrorCode int         `json:"error_code"`
	ErrorMsg  string      `json:"error_msg"`
}

type mingdao struct {
}

func NewMingdao() *mingdao {
	return &mingdao{}
}

func (d *mingdao) WorksheetGetFilterRows(ctx context.Context, param *domain.ThirdApiReq) (interface{}, error) {

	return d.worksheetGetFilterRows(ctx, param.Param)
}

func (d *mingdao) worksheetGetFilterRows(ctx context.Context, param map[string]interface{}) (interface{}, error) {
	path := "/api/v2/open/worksheet/getFilterRows"

	return d.process(ctx, http.MethodPost, path, param)
}

func (d *mingdao) WorksheetAddRow(ctx context.Context, param *domain.ThirdApiReq) (interface{}, error) {
	path := "/api/v2/open/worksheet/addRow"
	return d.process(ctx, http.MethodPost, path, param.Param)
}

func (d *mingdao) WorksheetAddRows(ctx context.Context, param *domain.ThirdApiReq) (interface{}, error) {
	path := "/api/v2/open/worksheet/addRows"
	return d.process(ctx, http.MethodPost, path, param.Param)
}

func (d *mingdao) WorksheetGetRowById(ctx context.Context, param *domain.ThirdApiReq) (interface{}, error) {
	worksheetId, ok := param.Param["worksheetId"]
	if !ok || !str.IsNotEmptyString(worksheetId) {
		return nil, errors.New("worksheetId is empty")
	}
	rowId, ok := param.Param["rowId"]
	if !ok || !str.IsNotEmptyString(rowId) {
		return nil, errors.New("rowId is empty")
	}
	return d.worksheetGetRowById(ctx, worksheetId.(string), rowId.(string))
}

type worksheetGetRowByIdResp struct {
	err error
	rs  interface{}
}

func (d *mingdao) worksheetGetRowById(ctx context.Context, worksheetId, id string, fields ...string) (interface{}, error) {
	path := "/api/v2/open/worksheet/getRowById"
	param := &domain.ThirdApiReq{
		Param: map[string]interface{}{
			"worksheetId": worksheetId,
			"rowId":       id,
		},
	}
	resp, err := d.process(ctx, http.MethodGet, path, param.Param)
	if err != nil {
		return nil, err
	}

	if len(fields) == 0 {
		return resp, nil
	}

	rs := make(map[string]interface{})
	res := resp.(map[string]interface{})
	for _, field := range fields {
		rs[field] = res[field]
	}
	return rs, nil
}

func (d *mingdao) WorksheetGetRowByIdPost(ctx context.Context, param *domain.ThirdApiReq) (interface{}, error) {
	path := "/api/v2/open/worksheet/getRowByIdPost"
	return d.process(ctx, http.MethodPost, path, param.Param)
}

func (d *mingdao) WorksheetEditRow(ctx context.Context, param *domain.ThirdApiReq) (interface{}, error) {
	path := "/api/v2/open/worksheet/editRow"
	return d.process(ctx, http.MethodPost, path, param.Param)
}

func (d *mingdao) WorksheetEditRows(ctx context.Context, param *domain.ThirdApiReq) (interface{}, error) {
	path := "/api/v2/open/worksheet/editRows"
	return d.process(ctx, http.MethodPost, path, param.Param)
}

func (d *mingdao) WorksheetDeleteRow(ctx context.Context, param *domain.ThirdApiReq) (interface{}, error) {
	path := "/api/v2/open/worksheet/deleteRow"
	return d.process(ctx, http.MethodPost, path, param.Param)
}

func (d *mingdao) WorksheetGetRowRelations(ctx context.Context, param *domain.ThirdApiReq) (interface{}, error) {
	path := "/api/v2/open/worksheet/getRowRelations"
	return d.process(ctx, http.MethodPost, path, param.Param)
}

func (d *mingdao) WorksheetGetRowShareLink(ctx context.Context, param *domain.ThirdApiReq) (interface{}, error) {
	path := "/api/v2/open/worksheet/getRowShareLink"
	return d.process(ctx, http.MethodPost, path, param.Param)
}

func (d *mingdao) WorksheetGetFilterRowsTotalNum(ctx context.Context, param *domain.ThirdApiReq) (interface{}, error) {
	path := "/api/v2/open/worksheet/getFilterRowsTotalNum"
	return d.process(ctx, http.MethodPost, path, param.Param)
}

func (d *mingdao) process(ctx context.Context, method, path string, body map[string]interface{}) (interface{}, error) {
	appKey, sign, domain, err := d.getConfig(ctx)
	if err != nil {
		return nil, err
	}

	body["appKey"] = appKey
	body["sign"] = sign

	url := fmt.Sprintf("%s%s", domain, path)

	var resp []byte
	var reqErr error
	if method == http.MethodPost {
		body, _ := json.Marshal(body)
		header := map[string]string{

			"Content-Type": "application/json",
		}
		resp, reqErr = xhttp.Req(url, http.MethodPost, bytes.NewReader(body), header)
	} else {
		queryStr := ""
		for key, item := range body {
			queryStr += fmt.Sprintf("%s=%v&", key, item)
		}
		url = fmt.Sprintf("%s?%s", url, strings.Trim(queryStr, "&"))
		resp, reqErr = xhttp.Req(url, http.MethodGet, nil, nil)
	}

	if reqErr != nil {
		return nil, reqErr
	}

	rs := &mingdaoResp{}
	if err := json.Unmarshal(resp, rs); err != nil {
		return nil, err
	}

	if !rs.Success {
		return nil, errors.New(rs.ErrorMsg)
	}

	return rs.Data, nil
}

func (d *mingdao) getConfig(ctx context.Context) (string, string, string, error) {
	return "6c64064632c7abd0", "ZGEyYTViODNjNmYwYzliYmZlMzFjZmExMDk3MmY2ZGE4NzZjMDg2MGUzMTIxZjQyZDUxMzY3MmRjYmRlMWJiOQ==", "https://z.daopuqifu.com", nil
}
