package bot

import (
	"context"
	"errors"
	"sync"

	"github.com/usual2970/meta-forge/internal/util/app"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	KindAddEassay = "eassay"
)

const (
	StateWaitTitle   = "wait_title"
	SteteWaitContent = "wait_content"
)

type Eassay struct {
	Title   string
	Content string
}

type Session struct {
	ChatID int64
	Kind   string
	State  string // 添加文章

	Eassay *Eassay
}

func NewSession(chatID int64) *Session {
	return &Session{
		ChatID: chatID,
	}
}

func (s *Session) Process(ctx context.Context, update tgbotapi.Update) (tgbotapi.Chattable, error) {

	return s.processUpdate(ctx, update)
}

func (s *Session) processUpdate(ctx context.Context, update tgbotapi.Update) (tgbotapi.Chattable, error) {
	// 处理命令
	if update.CallbackQuery != nil {
		return s.processCallback(ctx, update)
	}

	if update.Message.IsCommand() {
		return s.processCommand(ctx, update)
	}

	return s.processText(ctx, update)
}

func (s *Session) processCommand(ctx context.Context, update tgbotapi.Update) (tgbotapi.Chattable, error) {

	msg := update.Message
	switch update.Message.Command() {
	case "start", "menu":
		// 发送欢迎消息
		reply := tgbotapi.NewMessage(msg.From.ID, "欢迎使用英语文章背诵机器人")
		reply.ReplyMarkup = getKeyBoards()

		return reply, nil

	}

	return nil, errors.New("unknown command")
}

func (s *Session) processCallback(ctx context.Context, update tgbotapi.Update) (tgbotapi.Chattable, error) {

	switch update.CallbackData() {
	case "add":
		// 开始添加文章
		s.Kind = KindAddEassay
		s.State = StateWaitTitle

		reply := tgbotapi.NewMessage(update.CallbackQuery.From.ID, "请输入文章标题")

		return reply, nil
	case "return2menu":
		reply := tgbotapi.NewMessage(update.CallbackQuery.From.ID, "欢迎使用英语文章背诵机器人")
		reply.ReplyMarkup = getKeyBoards()

		return reply, nil

	}
	app.Get().Logger().Info("process callback", "data", update.CallbackData(), "query", *update.CallbackQuery)

	return nil, errors.New("unknown command")
}

func (s *Session) processText(ctx context.Context, update tgbotapi.Update) (tgbotapi.Chattable, error) {
	switch s.Kind {
	case KindAddEassay:
		return s.processEassay(ctx, update)
	}

	// 将文字转换成语音

	// 保存到数据库

	// 返回成功

	return nil, errors.New("unknown command")

}

func (s *Session) processEassay(ctx context.Context, update tgbotapi.Update) (tgbotapi.Chattable, error) {

	switch s.State {
	case StateWaitTitle:
		s.Eassay = &Eassay{
			Title: update.Message.Text,
		}
		s.State = SteteWaitContent
		reply := tgbotapi.NewMessage(update.Message.From.ID, "请输入文章内容")
		return reply, nil
	case SteteWaitContent:
		s.Eassay.Content = update.Message.Text
		reply := tgbotapi.NewMessage(update.Message.From.ID, "文章已保存")

		reply.ReplyMarkup = getReturnKeyBoards()

		app.Get().Logger().Info("Eassay", "eassay", *s.Eassay)
		s.clearState()

		return reply, nil
	}

	return nil, errors.New("unknown command")
}

func (s *Session) clearState() {
	s.State = ""
	s.Kind = ""
	s.Eassay = nil
}

var sessionMap *sessionList
var once sync.Once

func GetSessions() *sessionList {
	once.Do(func() {
		sessionMap = NewSessionList()
	})

	return sessionMap
}

func GetSession(update tgbotapi.Update) *Session {
	var chatID int64
	if update.CallbackQuery != nil {
		chatID = update.CallbackQuery.From.ID
	} else {
		chatID = update.Message.From.ID
	}
	session, ok := GetSessions().GetSession(chatID)
	if !ok {
		session = &Session{
			ChatID: chatID,
		}
		AddSession(session)
	}

	return session
}

func AddSession(session *Session) {
	GetSessions().AddSession(session)
}

type sessionList struct {
	sessions map[int64]*Session
	sync.RWMutex
}

func NewSessionList() *sessionList {
	return &sessionList{
		sessions: make(map[int64]*Session),
	}
}

func (s *sessionList) AddSession(session *Session) {
	s.Lock()
	defer s.Unlock()
	s.sessions[session.ChatID] = session
}

func (s *sessionList) GetSession(chatID int64) (*Session, bool) {
	s.RLock()
	defer s.RUnlock()
	rs, ok := s.sessions[chatID]

	return rs, ok
}

func getKeyBoards() tgbotapi.InlineKeyboardMarkup {

	return tgbotapi.NewInlineKeyboardMarkup([][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("添加文章", "add"),
			tgbotapi.NewInlineKeyboardButtonData("文章列表", "list"),
		},
	}...)

}

func getReturnKeyBoards() tgbotapi.InlineKeyboardMarkup {

	return tgbotapi.NewInlineKeyboardMarkup([][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("返回到菜单", "return2menu"),
		},
	}...)

}
