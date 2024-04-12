package wxbizmsgcrypt

import (
	"reflect"
	"testing"
)

func TestWXBizMsgCrypt_VerifyURL(t *testing.T) {
	type fields struct {
		token              string
		encoding_aeskey    string
		receiver_id        string
		protocol_processor ProtocolProcessor
	}
	type args struct {
		msg_signature string
		timestamp     string
		nonce         string
		echostr       string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
		want1  *CryptError
	}{
		{
			name: "1",
			fields: fields{
				token:              "gNDaF6TSqWVpGrqrUExzXxireKK",
				encoding_aeskey:    "b8B2REP6rARDR9kKPKwGdDjHdbcnkHfRVknOIv2WRjQ=",
				receiver_id:        "",
				protocol_processor: &XmlProcessor{},
			},
			args: args{
				echostr:       "LOqeD45hyEGpXPmLi2TWSAb1+GT9Sc3pRmA+Dfmo365U3XZA/Yc/PF8MT99OT0+zxWz/UsfUYk/pR9o9U1ROUw==",
				msg_signature: "e3cf08430f60f724beb4b9814bc64f208608458a",
				nonce:         "1710999898",
				timestamp:     "1711088357",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := &WXBizMsgCrypt{
				token:              tt.fields.token,
				encoding_aeskey:    tt.fields.encoding_aeskey,
				receiver_id:        tt.fields.receiver_id,
				protocol_processor: tt.fields.protocol_processor,
			}
			got, got1 := self.VerifyURL(tt.args.msg_signature, tt.args.timestamp, tt.args.nonce, tt.args.echostr)
			t.Log(string(got))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WXBizMsgCrypt.VerifyURL() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("WXBizMsgCrypt.VerifyURL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
