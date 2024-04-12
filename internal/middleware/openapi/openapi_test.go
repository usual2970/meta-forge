package openapi

import (
	"context"
	"testing"
	"time"
)

func TestPipeline_Execute(t *testing.T) {
	type fields struct {
		stack    []Middelware
		preIndex int
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "1",
			fields: fields{
				stack: []Middelware{
					func(ctx context.Context, next Next) {

						next()
						t.Log(1)
					},
					func(ctx context.Context, next Next) {
						t.Log(2)
						next()
					},

					func(ctx context.Context, next Next) {
						t.Log("正式逻辑")
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pipeline{
				stack:    tt.fields.stack,
				preIndex: tt.fields.preIndex,
			}
			p.Execute(tt.args.ctx)
		})
	}
}

func TestTime(t *testing.T) {

	location, err := time.LoadLocation("UTC")
	if err != nil {
		t.Log(err)
		return
	}

	// 创建一个带有纽约时区的当前时间
	currentTime := time.Now().In(location)
	formattedTime := currentTime.Format("Mon, 02 Jan 2006 15:04:05 GMT")

	t.Log(formattedTime)
}
