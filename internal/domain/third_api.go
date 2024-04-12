package domain

type ThirdApiReq struct {
	ThirdApiUriReq
	Param map[string]interface{} `json:"param"` // 参数
}

type ThirdApiUriReq struct {
	Project string `json:"project"`                        // 接口项目
	Uri     string `json:"uri"`                            // 接口标识
	AppKey  string `json:"appKey" v:"required#请输入 appkey"` // appkey
	Sign    string `json:"sign"  v:"required#请输入 sign"`    // 签名
	Ts      string `json:"ts"  v:"required#请输入 ts"`        // 时间戳单位秒
}

type ThirdApiResp struct {
	Data interface{} `json:"data"` //返回的原始数据
}
