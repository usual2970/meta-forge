package wecom

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/bytedance/gopkg/cache/asynccache"
)

func Test_wecom_GetAccessToken(t *testing.T) {
	type fields struct {
		corpId string
		secret string
		cache  asynccache.AsyncCache
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name:   "1",
			fields: fields{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWecom("ww4ea4ed42029597ff", "pu35X51mz_HVFAHPygHqb7xUCOhjGbmvkXtWQIWuOco")
			got, err := w.GetAccessToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("wecom.GetAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("wecom.GetAccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wecom_getAccessToken(t *testing.T) {
	type fields struct {
		corpId string
		secret string
		cache  asynccache.AsyncCache
	}
	type args struct {
		in0 string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &wecom{
				corpId: tt.fields.corpId,
				secret: tt.fields.secret,
				cache:  tt.fields.cache,
			}
			got, err := w.getAccessToken(tt.args.in0)
			if (err != nil) != tt.wantErr {
				t.Errorf("wecom.getAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("wecom.getAccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wecom_SyncMsg(t *testing.T) {
	type fields struct {
		corpId string
		secret string
		cache  asynccache.AsyncCache
	}
	type args struct {
		req *SyncMsgReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SyncMsgResp
		wantErr bool
	}{
		{
			name:   "1",
			fields: fields{},
			args: args{
				req: &SyncMsgReq{
					OpenKfid: "wkJK4fbAAAyn9MtqHggaquT3Y_yMZbIw",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWecom("ww4ea4ed42029597ff", "pu35X51mz_HVFAHPygHqb7xUCOhjGbmvkXtWQIWuOco")
			got, err := w.SyncMsg(tt.args.req)
			rs, _ := json.Marshal(got)
			t.Log(string(rs))
			if (err != nil) != tt.wantErr {
				t.Errorf("wecom.SyncMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wecom.SyncMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}
