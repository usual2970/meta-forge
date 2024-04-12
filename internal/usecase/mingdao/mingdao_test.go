package mingdao

import (
	"context"
	"testing"

	"github.com/usual2970/meta-forge/internal/domain"
)

func Test_usecase_GetPassId(t *testing.T) {
	type args struct {
		ctx   context.Context
		param *domain.MingdaoGetPassIdReq
	}
	tests := []struct {
		name    string
		u       *usecase
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "1",
			u:    &usecase{},
			args: args{
				ctx: context.Background(),
				param: &domain.MingdaoGetPassIdReq{
					UserName: "lif@yicheshuke.com",
					Password: "yc888888",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{}
			got, err := u.GetPassId(tt.args.ctx, tt.args.param)
			t.Log(got, err)
			got1, err := u.GetPassId(tt.args.ctx, tt.args.param)
			t.Log(got1, err)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetPassId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("usecase.GetPassId() = %v, want %v", got, tt.want)
			}
		})
	}
}
