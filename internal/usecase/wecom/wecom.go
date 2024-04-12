package wecom

import (
	"context"
	"encoding/xml"
	"fmt"
	"regexp"
	"sync"

	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/util/app"
	"github.com/usual2970/meta-forge/internal/util/dashscope"
	"github.com/usual2970/meta-forge/internal/util/wecom"
	"github.com/usual2970/meta-forge/internal/util/wxbizmsgcrypt"
	"github.com/usual2970/meta-forge/internal/util/xtools"
	"github.com/usual2970/meta-forge/internal/util/zhipu"

	"github.com/panjf2000/ants/v2"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/tools"
	"github.com/tmc/langchaingo/tools/serpapi"
)

var once sync.Once
var instance domain.IWecomUsecase

var _promptPrefix = `Assistant is a large language model trained by OpenAI.

Assistant在大模型的基础上还提供了本地知识库LocalKnowledge工具，如果你不知道如何回答请优先使用LocalKnowledge工具。

Assistant is designed to be able to assist with a wide range of tasks, from answering simple questions to providing in-depth explanations and discussions on a wide range of topics. As a language model, Assistant is able to generate human-like text based on the input it receives, allowing it to engage in natural-sounding conversations and provide responses that are coherent and relevant to the topic at hand.

Assistant is constantly learning and improving, and its capabilities are constantly evolving. It is able to process and understand large amounts of text, and can use this knowledge to provide accurate and informative responses to a wide range of questions. Additionally, Assistant is able to generate its own text based on the input it receives, allowing it to engage in discussions and provide explanations and descriptions on a wide range of topics.

Overall, Assistant is a powerful tool that can help with a wide range of tasks and provide valuable insights and information on a wide range of topics. Whether you need help with a specific question or just want to have a conversation about a particular topic, Assistant is here to assist.

TOOLS:
------

Assistant has access to the following tools:

{{.tool_descriptions}}`

type usecase struct {
	secretRepo domain.ISecretRepository
	wecomRepo  domain.IWecomRepository
	eventChan  chan domain.WecomEvent
	cancel     context.CancelFunc
	executors  map[string]agents.Executor
	sync.RWMutex
}

func NewUsecase(secretRepo domain.ISecretRepository, wecomRepo domain.IWecomRepository) domain.IWecomUsecase {
	once.Do(func() {
		u := &usecase{
			secretRepo: secretRepo,
			wecomRepo:  wecomRepo,
			eventChan:  make(chan domain.WecomEvent, 1),
			executors:  make(map[string]agents.Executor),
		}

		ctx, cancel := context.WithCancel(context.Background())
		u.cancel = cancel

		for i := 0; i < 16; i++ {
			ants.Submit(func() {
				u.process(ctx)
			})
		}
		instance = u
	})

	return instance
}

func (u *usecase) TransServiceState(ctx context.Context, req *wecom.TransServiceStateReq) error {
	config, err := u.secretRepo.Get(ctx, "uri='wecom'")
	if err != nil {
		app.Get().Logger().Info("get wecom config failed", "err", err)
		return err
	}

	wx := wecom.NewWecom(config.Ext["corpId"], config.Ext["secret"])

	return wx.TransServiceState(req)
}

func (u *usecase) GetServiceState(ctx context.Context, req *wecom.GetServiceStateReq) (*wecom.GetServiceStateResp, error) {
	config, err := u.secretRepo.Get(ctx, "uri='wecom'")
	if err != nil {
		app.Get().Logger().Info("get wecom config failed", "err", err)
		return nil, err
	}

	wx := wecom.NewWecom(config.Ext["corpId"], config.Ext["secret"])

	return wx.GetServiceState(req)
}

func (u *usecase) Publish(ctx context.Context, req *domain.WecomEvent) error {
	ants.Submit(func() {
		u.eventChan <- *req
	})
	return nil
}

func (u *usecase) process(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			return
		case event := <-u.eventChan:
			l := app.Get().Logger().With("event", event)
			l.Info("wecom event")

			// 开始处理
			if err := u.toProcess(ctx, event); err != nil {
				l.Info("wecom process error", "err", err)
				continue
			}
			l.Info("wecom process success")
		}
	}
}

func (u *usecase) toProcess(ctx context.Context, event domain.WecomEvent) error {
	l := app.Get().Logger().With("event", event)
	// 获取消息

	config, err := u.secretRepo.Get(ctx, "uri='wecom'")
	if err != nil {
		app.Get().Logger().Info("get wecom config failed", "err", err)
		return err
	}

	// 获取cursor
	nextCursor := ""
	stuff, err := u.wecomRepo.Get(ctx, "open_kf_id='"+event.OpenKfId+"'")
	if err == nil {
		nextCursor = stuff.NextCursor
	}

	wx := wecom.NewWecom(config.Ext["corpId"], config.Ext["secret"])

	resp, err := wx.SyncMsg(&wecom.SyncMsgReq{
		Cursor:      nextCursor,
		Token:       event.Token,
		VoiceFormat: 0,
		OpenKfid:    event.OpenKfId,
		Limit:       1000,
	})
	if err != nil {
		l.Info("wecom sync msg error", "err", err)
		return err
	}

	l.Info("wecom sync msg success", "resp", resp)

	// 将cursor写入数据库
	stuffId := ""
	if stuff != nil {
		stuffId = stuff.Id
	}

	if err := u.wecomRepo.SetNextCursor(ctx, &domain.RxSupportStuff{
		Meta: domain.Meta{
			Id: stuffId,
		},
		OpenKfId:   event.OpenKfId,
		NextCursor: resp.NextCursor,
	}); err != nil {
		l.Info("wecom set next cursor error", "err", err)
	}

	for _, msg := range resp.MsgList {
		// 处理消息
		text := ""
		switch msg.Msgtype {
		case wecom.MsgtypeText:
			text = msg.Text.Content
		case wecom.MsgtypeVoice:
			text = u.getVoiceText(ctx, msg.Voice.MediaId)
		}

		l.Info("wecom msg", "msg", text)
		// 调用模型
		rs, err := u.completion(ctx, text, msg.OpenKfid, msg.ExternalUserid)
		if err != nil {
			l.Info("wecom completion error", "err", err)
			continue
		}
		l.Info("wecom completion success", "resp", rs)
		// 发送消息
		rsMsg := wecom.NewTextMsg(rs, msg.ExternalUserid, wecom.WithOpenKfid(msg.OpenKfid))
		if _, err := wx.SendMsg(&rsMsg); err != nil {
			l.Info("wecom send msg error", "err", err)
			continue
		}
		l.Info("wecom send msg success")

	}

	return nil
}

func (u *usecase) getVoiceText(ctx context.Context, mediaId string) string {
	l := app.Get().Logger().With("media_id", mediaId).With("module", "getVoiceText")
	accessToken, err := u.Accesstoken(ctx)
	if err != nil {
		l.Info("wecom get accesstoken error", "err", err)
		return ""
	}

	media := &domain.RxMedia{
		MediaId: mediaId,
		File:    fmt.Sprintf("%s?access_token=%s&media_id=%s", "https://qyapi.weixin.qq.com/cgi-bin/media/get", accessToken, mediaId),
	}

	if err := u.wecomRepo.SaveMedia(ctx, media); err != nil {
		l.Info("wecom save media error", "err", err)
		return ""
	}

	rs, err := u.wecomRepo.GetMedia(ctx, "media_id='"+mediaId+"'")
	if err != nil {
		l.Info("wecom get media error", "err", err)
		return ""
	}

	d := dashscope.New(dashscope.WithApiKey("sk-b4efe5ac3aa646c085dbe78a4f32c689"))
	req := &dashscope.MultimodalGenerationReq{
		Model: "qwen-audio-turbo",
		Input: dashscope.MultimodalGenerationInput{
			Messages: []dashscope.MultimodalGenerationMessage{
				{
					Role: "user",
					Content: []dashscope.MultimodalGenerationMessageContent{
						{
							Audio: rs.File,
						},
						{
							Text: "这段音频在说什么?",
						},
					},
				},
			},
		},
		Parameters: map[string]interface{}{},
	}
	resp, err := d.MultimodalGeneration(req)

	if err != nil {
		l.Info("wecom multimodal generation error", "err", err, "req", req)
		return ""
	}

	if resp.Code != "" {
		l.Info("wecom multimodal generation error", "err", resp.Message, "req", req)
		return ""
	}

	text := resp.Output.Choices[0].Message.Content[0].Text

	reg := regexp.MustCompile(`.*："(.*?)"`)

	matches := reg.FindStringSubmatch(text)

	if len(matches) == 0 {
		l.Info("wecom get voice text error", "text", text)
		return ""
	}
	l.Info("wecom get voice text success", "text", matches[1])
	return matches[1]

}

func (u *usecase) getExecutor(ctx context.Context, openKfid, externalUserid string) (*agents.Executor, error) {

	key := fmt.Sprintf("%s_%s", openKfid, externalUserid)

	u.RLock()
	e, ok := u.executors[key]
	u.RUnlock()

	if !ok {
		zhipuConfig, err := u.secretRepo.Get(ctx, "uri='zhipu'")
		if err != nil {
			return nil, err
		}
		llm := zhipu.NewZhipu(zhipuConfig.SecretKey, llms.WithTemperature(0.95), llms.WithTopP(0.70))

		searchTool, err := serpapi.New()
		if err != nil {
			return nil, err
		}

		knowledge := xtools.NewKnowledge(zhipuConfig.SecretKey)

		toolList := []tools.Tool{knowledge, xtools.GetOrderNum{}, searchTool}
		e, err = agents.Initialize(
			llm,
			toolList,
			agents.ConversationalReactDescription,
			agents.WithParserErrorHandler(agents.NewParserErrorHandler(func(s string) string {
				return "Check your output and make sure it conforms, use the Action/Action Input syntax"
			})),
			agents.WithMemory(memory.NewConversationWindowBuffer(8)),
			agents.WithPromptPrefix(_promptPrefix),
		)
		if err != nil {
			app.Get().Logger().Error("Initialize angents failed", "err", err)
			return nil, err
		}
		u.Lock()
		u.executors[key] = e
		u.Unlock()
	}

	return &e, nil
}

func (u *usecase) completion(ctx context.Context, text, openKfid, externalUserid string) (string, error) {
	e, err := u.getExecutor(ctx, openKfid, externalUserid)
	if err != nil {
		app.Get().Logger().Error("Get executor failed", "err", err)
		return "", err
	}
	rs, err := chains.Run(context.Background(), e, text)
	if err != nil {
		app.Get().Logger().Error("Get result failed", "err", err)
		return "", err
	}

	return rs, nil
}

func (u *usecase) Exit() error {
	u.cancel()
	return nil
}

func (u *usecase) ClearExecutor(ctx context.Context) {
	u.Lock()
	defer u.Unlock()

	u.executors = make(map[string]agents.Executor)
}

func (u *usecase) Notify(ctx context.Context, req *domain.WecomNoticeReq) (string, error) {
	app.Get().Logger().Info("accept wecom notify", "req", req)
	config, err := u.secretRepo.Get(ctx, "uri='wecom'")
	if err != nil {
		app.Get().Logger().Info("get wecom config failed", "err", err)
		return "", err
	}

	wx := wxbizmsgcrypt.NewWXBizMsgCrypt(config.Ext["token"], config.Ext["encodingAESKey"], "", wxbizmsgcrypt.XmlType)

	if req.Echostr != "" {
		rs, err := wx.VerifyURL(req.MsgSignature, req.Timestamp, req.Nonce, req.Echostr)
		if err != nil {
			app.Get().Logger().Info("wecom verify url failed", "err", err)
			return "", err
		}

		return string(rs), nil
	}

	msg, err := wx.DecryptMsg(req.MsgSignature, req.Timestamp, req.Nonce, req.Post)
	if err != nil {
		app.Get().Logger().Info("wecom verify url failed", "err", err)
		return "", err
	}

	var event domain.WecomEvent
	err = xml.Unmarshal(msg, &event)
	if err != nil {
		app.Get().Logger().Info("wecom parse xml failed", "err", err)
		return "", err
	}

	app.Get().Logger().Info("get msg from wecom", "msg", string(msg))

	if err := u.Publish(ctx, &event); err != nil {
		app.Get().Logger().Info("wecom publish failed", "err", err)
		return "", err
	}

	return "", nil
}

func (u *usecase) Accesstoken(ctx context.Context) (string, error) {
	config, err := u.secretRepo.Get(ctx, "uri='wecom'")
	if err != nil {
		app.Get().Logger().Info("get wecom config failed", "err", err)
		return "", err
	}

	wx := wecom.NewWecom(config.Ext["corpId"], config.Ext["secret"])
	return wx.GetAccessToken()
}
