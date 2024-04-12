package dashscope

import (
	"reflect"
	"testing"
)

func TestDashscope_MultimodalGeneration(t *testing.T) {
	type fields struct {
		apiKey string
	}
	type args struct {
		req *MultimodalGenerationReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *MultimodalGenerationResp
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				apiKey: "sk-b4efe5ac3aa646c085dbe78a4f32c689",
			},
			args: args{
				req: &MultimodalGenerationReq{
					Model: "qwen-audio-turbo",
					Input: MultimodalGenerationInput{
						Messages: []MultimodalGenerationMessage{
							{
								Role: "user",
								Content: []MultimodalGenerationMessageContent{
									{
										Audio: "https://dashscope.oss-cn-beijing.aliyuncs.com/audios/2channel_16K.wav",
									},
									{
										Text: "这段音频在说什么?",
									},
								},
							},
						},
					},
					Parameters: map[string]interface{}{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dashscope{
				apiKey: tt.fields.apiKey,
			}
			got, err := d.MultimodalGeneration(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Dashscope.MultimodalGeneration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dashscope.MultimodalGeneration() = %v, want %v", got, tt.want)
			}
		})
	}
}
