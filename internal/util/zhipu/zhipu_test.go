package zhipu

import (
	"context"
	"reflect"
	"testing"

	"github.com/tmc/langchaingo/llms"
)

func TestZhipu_Call(t *testing.T) {
	type fields struct {
		apiKey string
	}
	type args struct {
		ctx     context.Context
		prompt  string
		options []llms.CallOption
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
				apiKey: "84ca0ef664ff6ac12ba94618a433cb77.KcWwtJgAQf4VJ83e",
			},
			args: args{
				ctx:     context.Background(),
				prompt:  "你好",
				options: []llms.CallOption{llms.WithModel("GLM-4")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Zhipu{
				apiKey: tt.fields.apiKey,
			}
			got, err := z.Call(tt.args.ctx, tt.args.prompt, tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zhipu.Call() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zhipu.Call() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZhipu_CreateEmbedding(t *testing.T) {
	type fields struct {
		apiKey string
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
				apiKey: "7a909dd632f00f28a0d08efba999b3dc.50FHMHrklQIaaD16",
			},
			args: args{
				ctx:   context.Background(),
				texts: []string{"你好"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Zhipu{
				apiKey: tt.fields.apiKey,
			}
			got, err := z.CreateEmbedding(tt.args.ctx, tt.args.texts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zhipu.CreateEmbedding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zhipu.CreateEmbedding() = %v, want %v", got, tt.want)
			}
		})
	}
}
