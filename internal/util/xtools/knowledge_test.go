package xtools

import (
	"context"
	"testing"

	"github.com/tmc/langchaingo/callbacks"
)

func TestKnowledge_Call(t *testing.T) {
	type fields struct {
		apiKey           string
		CallbacksHandler callbacks.Handler
	}
	type args struct {
		ctx   context.Context
		input string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				apiKey: "7a909dd632f00f28a0d08efba999b3dc.50FHMHrklQIaaD16",
			},
			args: args{
				ctx:   context.Background(),
				input: "林润云下载地址",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Knowledge{
				apiKey:           tt.fields.apiKey,
				CallbacksHandler: tt.fields.CallbacksHandler,
			}
			got, err := c.Call(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Knowledge.Call() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Knowledge.Call() = %v, want %v", got, tt.want)
			}
		})
	}
}
