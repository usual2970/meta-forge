package domain

import (
	"context"

	"github.com/usual2970/meta-forge/internal/util/wecom"
)

type WecomNoticeReq struct {
	MsgSignature string `json:"msg_signature" query:"msg_signature"`
	Timestamp    string `json:"timestamp" query:"timestamp"`
	Nonce        string `json:"nonce" query:"nonce"`
	Echostr      string `json:"echostr" query:"echostr"`

	Post []byte `json:"post"`
}

type WecomEvent struct {
	ToUserName string `xml:"ToUserName"`
	CreateTime string `xml:"CreateTime"`
	MsgType    string `xml:"MsgType"`
	Event      string `xml:"Event"`
	Token      string `xml:"Token"`
	OpenKfId   string `xml:"OpenKfId"`
}

type WecomPublishReq struct {
}
type IWecomUsecase interface {
	Notify(ctx context.Context, req *WecomNoticeReq) (string, error)

	// 获取会话状态
	GetServiceState(ctx context.Context, req *wecom.GetServiceStateReq) (*wecom.GetServiceStateResp, error)

	// 更新会话状态
	TransServiceState(ctx context.Context, req *wecom.TransServiceStateReq) error

	// 投递任务
	Publish(ctx context.Context, req *WecomEvent) error

	ClearExecutor(ctx context.Context)

	Accesstoken(ctx context.Context) (string, error)

	Exit() error
}

type RxSupportStuff struct {
	Meta
	OpenKfId   string `json:"openKfId"`
	NextCursor string `json:"nextCursor"`
}

type RxMedia struct {
	Meta
	MediaId string `json:"media_id"`
	File    string `json:"file"`
}

type IWecomRepository interface {
	Get(ctx context.Context, filter string) (*RxSupportStuff, error)
	SetNextCursor(ctx context.Context, data *RxSupportStuff) error

	GetMedia(ctx context.Context, filter string) (*RxMedia, error)
	SaveMedia(ctx context.Context, data *RxMedia) error
}
