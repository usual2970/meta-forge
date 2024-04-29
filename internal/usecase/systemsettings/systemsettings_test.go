package systemsettings

import (
	"context"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/usual2970/meta-forge/internal/domain"
	systemsettings "github.com/usual2970/meta-forge/internal/repository/system_settings"
	"github.com/usual2970/meta-forge/internal/util/app"
)

const testDataDir = "/Users/liuxuanyao/work/metaforge/pb_data"

func Test_usecase_Initialize(t *testing.T) {
	_ = app.GetTest()
	type fields struct {
		repo domain.ISystemSettingsRepository
	}
	type args struct {
		ctx context.Context
		req *domain.InitializeReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "mysql",
			fields: fields{
				repo: systemsettings.NewRepository(),
			},
			args: args{
				ctx: context.Background(),
				req: &domain.InitializeReq{
					Kind:     "mysql",
					Host:     "127.0.0.1",
					Port:     "13306",
					User:     "root",
					Password: "root",
					Database: "jichedi",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				repo: tt.fields.repo,
			}
			if err := u.Initialize(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("usecase.Initialize() error = %v, wantErr %v", err, tt.wantErr)
			}

	}
}
