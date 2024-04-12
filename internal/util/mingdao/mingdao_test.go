package mingdao

import (
	"context"
	"reflect"
	"testing"

	"github.com/usual2970/meta-forge/internal/domain"
)

func Test_mingdao_WorksheetGetFilterRowsTotalNum(t *testing.T) {
	type args struct {
		ctx   context.Context
		param *domain.ThirdApiReq
	}
	tests := []struct {
		name    string
		d       *mingdao
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "1",
			d:    &mingdao{},
			args: args{
				ctx: context.Background(),
				param: &domain.ThirdApiReq{
					Param: map[string]interface{}{
						"worksheetId": "ddlb",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &mingdao{}
			got, err := d.WorksheetGetFilterRowsTotalNum(tt.args.ctx, tt.args.param)
			t.Log(got, err)
			if (err != nil) != tt.wantErr {
				t.Errorf("mingdao.WorksheetGetFilterRowsTotalNum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mingdao.WorksheetGetFilterRowsTotalNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
