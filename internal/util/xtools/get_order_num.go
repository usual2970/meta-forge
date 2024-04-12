package xtools

import (
	"context"
	"fmt"
	"strings"

	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/util/mingdao"

	"github.com/tmc/langchaingo/callbacks"
	"github.com/tmc/langchaingo/tools"
)

// Calculator is a tool that can do math.
type GetOrderNum struct {
	CallbacksHandler callbacks.Handler
}

var _ tools.Tool = GetOrderNum{}

// Description returns a string describing the calculator tool.
func (c GetOrderNum) Description() string {
	return `"各种订单统计封装"
			"用于统计各种类型的订单的数量"
			"输入参数为:订单类型-订单状态-业务员"
			"订单状态：有审批完成、全部等状态，不确定则输入：全部"
			"业务员：不确定则输入：全部"`
}

// Name returns the name of the tool.
func (c GetOrderNum) Name() string {
	return "getOrderNum"
}

func (c GetOrderNum) Call(ctx context.Context, input string) (string, error) {
	if c.CallbacksHandler != nil {
		c.CallbacksHandler.HandleToolStart(ctx, input)
	}

	md := mingdao.NewMingdao()

	rs, err := md.WorksheetGetFilterRowsTotalNum(context.Background(), &domain.ThirdApiReq{
		Param: map[string]interface{}{
			"worksheetId": getWorksheetId(input),
			"filters":     getFilters(input),
		},
	})

	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%v", rs)

	if c.CallbacksHandler != nil {
		c.CallbacksHandler.HandleToolEnd(ctx, result)
	}

	return result, nil
}

func getWorksheetId(input string) string {
	if strings.Contains(input, "进件") {
		return "ddlb"
	} else if strings.Contains(input, "评估") {
		return "dycwpg"
	}
	return "ddlb"
}

func getControlId(worksheetId string, name string) string {

	if worksheetId == "ddlb" && name == "状态" {
		return "fkjg"
	}

	if worksheetId == "ddlb" && name == "业务员" {
		return "yewuyuan"
	}

	if worksheetId == "dycwpg" && name == "业务员" {
		return "xiadanren"
	}

	if worksheetId == "dycwpg" && name == "状态" {
		return "pgzt"
	}

	return ""
}

func getFilters(input string) []map[string]interface{} {
	rs := make([]map[string]interface{}, 0)

	worksheetId := getWorksheetId(input)
	params := strings.Split(input, "-")
	if len(params) == 1 {
		return rs
	}
	if len(params) >= 2 && params[1] != "全部" {
		rs = append(rs, map[string]interface{}{
			"controlId":  getControlId(worksheetId, "状态"),
			"dataType":   11,
			"spliceType": 1,
			"filterType": 2,
			"values":     []string{params[1]},
		})
	}

	if len(params) == 3 && params[2] != "全部" {
		rs = append(rs, map[string]interface{}{
			"controlId":  getControlId(worksheetId, "业务员"),
			"dataType":   26,
			"spliceType": 1,
			"filterType": 2,
			"values":     []string{getUserId(params[2])},
		})
	}

	return rs
}

func getUserId(name string) string {
	name = strings.ReplaceAll(name, "\nObservation:", "")
	return users[name].Id
}

type MingdaoUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var users map[string]MingdaoUser = map[string]MingdaoUser{
	"袁华生": {Id: "3ccf6ffa-d0f0-40c8-a404-18c17347ed2f", Name: "袁华生"}, "王万利": {Id: "972a5b5d-43b0-4140-b1a1-13a563364009", Name: "王万利"}, "王兰青": {Id: "0b601188-2653-44c5-9238-c66bcdf0a30d", Name: "王兰青"}, "王英": {Id: "a7ef3b57-059e-453f-a03b-ba5e9d0bd08e", Name: "王英"}, "刘洋": {Id: "086059c7-ef05-43ba-9155-19f871d9f236", Name: "刘洋"}, "刘备": {Id: "d948b883-a5c2-4b86-bd3f-c45719c0a0df", Name: "刘备"}, "应勇庆": {Id: "5981bd55-8b61-4b5d-ae79-e43042a8cd34", Name: "应勇庆"}, "牛建华": {Id: "da6e8ae6-427e-4bf1-863e-16f162df1fda", Name: "牛建华"}, "张喜财": {Id: "0a139f14-971c-4f5c-b6b7-79131a765819", Name: "张喜财"}, "王星星": {Id: "0e16c060-3fc5-4c34-933e-be98acf0ddeb", Name: "王星星"}, "庄妮": {Id: "b381b543-ba42-4fde-8fc8-15ffae7548c0", Name: "庄妮"}, "何建忠": {Id: "7208f0f4-41d9-4e2d-b36d-c068d250969e", Name: "何建忠"}, "杜立衡": {Id: "8752c957-b08c-4a38-99df-6ddd27ba1b7e", Name: "杜立衡"}, "孙喆 ": {Id: "69de627a-2448-4e65-a5f6-656e57a60cd1", Name: "孙喆 "}, "朱杰": {Id: "a3928624-8feb-45f4-a812-51c1f018812c", Name: "朱杰"}, "田洁洁": {Id: "103d89d6-b704-4c86-877d-5f8c0eed02f5", Name: "田洁洁"}, "李莉": {Id: "d564e158-361a-4ada-93fd-e21769ad6b12", Name: "李莉"}, "钟继龙": {Id: "8dbf9470-7365-4eab-8671-c52d55032b4a", Name: "钟继龙"}, "吴亦婷": {Id: "5c24de60-e6a0-4cee-8f2d-d8500815019c", Name: "吴亦婷"}, "15012345678": {Id: "2cb73c96-d62c-481c-b2e6-405e58836cd1", Name: "15012345678"},
	"x": {Id: "5581551d-3f47-44a9-b19e-8b10583fb65d", Name: "x"},
}
