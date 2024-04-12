package wenxin

import (
	"context"
	"reflect"
	"testing"

	ernie "github.com/anhao/go-ernie"
)

func TestWenxin_CreateEmbedding(t *testing.T) {
	type fields struct {
		client *ernie.Client
	}
	type args struct {
		ctx   context.Context
		texts []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    [][]float32
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				client: getWenxinClient("DF5iCGyTVWPX3UsgKugTOyVp", "1jpP60hGRzA28IRbf3Gy1j3F5IKUC66X"),
			},
			args: args{
				ctx:   context.Background(),
				texts: []string{"我们都有一个家名字叫中国"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Wenxin{
				client: tt.fields.client,
			}
			got, err := w.CreateEmbedding(tt.args.ctx, tt.args.texts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Wenxin.CreateEmbedding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wenxin.CreateEmbedding() = %v, want %v", got, tt.want)
			}
		})
	}
}
