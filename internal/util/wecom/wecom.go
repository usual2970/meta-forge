package wecom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	xhttp "github.com/usual2970/meta-forge/internal/util/http"

	"github.com/bytedance/gopkg/cache/asynccache"
	jsoniter "github.com/json-iterator/go"
)

var instance *wecom
var once sync.Once

type wecom struct {
	corpId, secret string
	cache          asynccache.AsyncCache
}

const getTokenUrl = "https://qyapi.weixin.qq.com/cgi-bin/gettoken"

const syncMsgUrl = "https://qyapi.weixin.qq.com/cgi-bin/kf/sync_msg"

const sendMsgUrl = "https://qyapi.weixin.qq.com/cgi-bin/kf/send_msg"

const getServiceStateUrl = "https://qyapi.weixin.qq.com/cgi-bin/kf/service_state/get"

const transServiceStateUrl = "https://qyapi.weixin.qq.com/cgi-bin/kf/service_state/trans"

func NewWecom(corpId, secret string) *wecom {
	once.Do(func() {
		instance = &wecom{
			corpId: corpId,
			secret: secret,
		}
		cache := asynccache.NewAsyncCache(asynccache.Options{
			RefreshDuration: time.Hour,
			EnableExpire:    true,
			ExpireDuration:  time.Hour * 2,
			Fetcher: func(key string) (interface{}, error) {
				return instance.getAccessToken(key)
			},
		})

		instance.cache = cache

	})
	return instance
}

func (w *wecom) GetAccessToken() (string, error) {
	rs, err := w.cache.Get("access_token")
	if err != nil {
		return "", fmt.Errorf("wecom getAccessToken error: %v", err)
	}
	return rs.(string), nil
}

type AccesstokenResp struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (w *wecom) getAccessToken(_ string) (string, error) {

	url := fmt.Sprintf("%s?corpid=%s&corpsecret=%s", getTokenUrl, w.corpId, w.secret)

	resp, err := xhttp.Req(url, http.MethodGet, nil, nil)

	if err != nil {
		return "", fmt.Errorf("wecom getAccessToken error: %v", err)
	}

	var respData AccesstokenResp
	if err := json.Unmarshal(resp, &respData); err != nil {
		return "", fmt.Errorf("wecom getAccessToken error: %v", err)
	}
	if respData.Errcode != 0 {
		return "", fmt.Errorf("wecom getAccessToken error: %v", respData.Errmsg)
	}

	return respData.AccessToken, nil
}

type GetServiceStateReq struct {
	OpenKfid       string `json:"open_kfid" query:"open_kfid"`
	ExternalUserid string `json:"external_userid" query:"external_userid"`
}

type GetServiceStateResp struct {
	Errcode        int    `json:"errcode"`
	Errmsg         string `json:"errmsg"`
	ServiceState   int    `json:"service_state"`
	ServicerUserid string `json:"servicer_userid"`
}

func (w *wecom) GetServiceState(req *GetServiceStateReq) (*GetServiceStateResp, error) {
	accessToken, err := w.GetAccessToken()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s?access_token=%s", getServiceStateUrl, accessToken)

	body, _ := json.Marshal(req)
	resp, err := xhttp.Req(url, http.MethodPost, bytes.NewBuffer(body), map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, err
	}
	var respData *GetServiceStateResp
	if err := json.Unmarshal(resp, &respData); err != nil {
		return nil, err
	}
	if respData.Errcode != 0 {
		return nil, fmt.Errorf("wecom get service state error: %v", respData.Errmsg)
	}
	return respData, nil
}

type TransServiceStateReq struct {
	OpenKfid       string `json:"open_kfid"`
	ExternalUserid string `json:"external_userid"`
	ServiceState   int    `json:"service_state"`
	ServicerUserid string `json:"servicer_userid"`
}

type TransServiceStateResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	MsgCode string `json:"msg_code"`
}

func (w *wecom) TransServiceState(req *TransServiceStateReq) error {
	accessToken, err := w.GetAccessToken()
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s?access_token=%s", transServiceStateUrl, accessToken)

	body, _ := jsoniter.MarshalToString(req)
	resp, err := xhttp.Req(url, http.MethodPost, bytes.NewBuffer([]byte(body)), map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return err
	}
	var respData TransServiceStateResp
	if err := json.Unmarshal(resp, &respData); err != nil {
		return err
	}
	if respData.Errcode != 0 {
		return fmt.Errorf("wecom trans service state error: %v", respData.Errmsg)
	}
	return nil
}

func (w *wecom) SyncMsg(req *SyncMsgReq) (*SyncMsgResp, error) {
	accessToken, err := w.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s?access_token=%s", syncMsgUrl, accessToken)

	body, _ := json.Marshal(req)

	resp, err := xhttp.Req(url, http.MethodPost, bytes.NewBuffer(body), map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, err
	}

	var respData SyncMsgResp
	if err := json.Unmarshal(resp, &respData); err != nil {
		return nil, err
	}
	if respData.Errcode != 0 {
		return nil, fmt.Errorf("wecom sync msg error: %v", respData.Errmsg)
	}

	return &respData, nil
}

func (w *wecom) SendMsg(req *Msg) (*SendMsgResp, error) {
	accessToken, err := w.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s?access_token=%s", sendMsgUrl, accessToken)

	resp, err := xhttp.Req(url, http.MethodPost, bytes.NewBuffer([]byte(req.String())), map[string]string{"Content-Type": "application/json"})

	if err != nil {
		return nil, err
	}

	var respData SendMsgResp
	if err := json.Unmarshal(resp, &respData); err != nil {
		return nil, err
	}
	if respData.Errcode != 0 {
		return nil, fmt.Errorf("wecom send msg error: %v", respData.Errmsg)
	}

	return &respData, nil
}

type SendMsgResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Msgid   string `json:"msgid"`
}

type SyncMsgReq struct {
	Cursor string `json:"cursor"`
	Token  string `json:"token"`
	Limit  uint32 `json:"limit"`

	VoiceFormat uint32 `json:"voice_format"`
	OpenKfid    string `json:"open_kfid"`
}

type SyncMsgResp struct {
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	NextCursor string `json:"next_cursor"`
	HasMore    int    `json:"has_more"`
	MsgList    []Msg  `json:"msg_list"`
}

type Msg struct {
	Msgid          string `json:"msgid,omitempty"`
	OpenKfid       string `json:"open_kfid,omitempty"`
	ExternalUserid string `json:"external_userid,omitempty"`
	SendTime       int    `json:"send_time,omitempty"`
	Origin         int    `json:"origin,omitempty"`
	ServicerUserid string `json:"servicer_userid,omitempty"`
	Touser         string `json:"touser,omitempty"`
	Msgtype        string `json:"msgtype"`

	Text MsgText `json:"text"`

	Image MsgImage `json:"image"`

	Voice MsgVoice `json:"voice"`

	Event MsgEvent `json:"event"`
}

func (m *Msg) String() string {
	rs, _ := jsoniter.MarshalToString(m)
	return rs
}

const MsgtypeText = "text"
const MsgtypeImage = "image"
const MsgtypeVoice = "voice"

type MsgText struct {
	Content string `json:"content"`
	MenuId  string `json:"menu_id"`
}

type MsgImage struct {
	MediaId string `json:"media_id"`
}

type MsgVoice struct {
	MediaId string `json:"media_id"`
}

type MsgEvent struct {
	EventType      string `json:"event_type"`
	OpenKfid       string `json:"open_kfid"`
	ExternalUserid string `json:"external_userid"`
	Scene          string `json:"scene"`
	SceneParam     string `json:"scene_param"`
	WelcomeCode    string `json:"welcome_code"`
	WechatChannels struct {
		Nickname string `json:"nickname"`
		Scene    int    `json:"scene"`
	} `json:"wechat_channels"`
}

type Option func(*Msg)

func WithMsgid(msgid string) Option {
	return func(m *Msg) {
		m.Msgid = msgid
	}
}

func WithOpenKfid(openKfid string) Option {
	return func(m *Msg) {
		m.OpenKfid = openKfid
	}
}

func WithExternalUserid(externalUserid string) Option {
	return func(m *Msg) {
		m.ExternalUserid = externalUserid
	}
}

func WithSendTime(sendTime int) Option {
	return func(m *Msg) {
		m.SendTime = sendTime
	}
}

func WithOrigin(origin int) Option {
	return func(m *Msg) {
		m.Origin = origin
	}
}

func WithServicerUserid(servicerUserid string) Option {
	return func(m *Msg) {
		m.ServicerUserid = servicerUserid
	}
}

func NewTextMsg(text string, to string, opts ...Option) Msg {
	msg := &Msg{
		Msgtype: MsgtypeText,
		Touser:  to,
		Text: MsgText{
			Content: text,
		},
	}

	for _, opt := range opts {
		opt(msg)
	}

	return *msg
}
