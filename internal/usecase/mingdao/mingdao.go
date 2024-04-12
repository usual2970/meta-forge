package mingdao

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/util/app"

	"github.com/playwright-community/playwright-go"
)

var mdPassIdOnce sync.Once
var mdPassIdInstance *mdPassId

type mdPassId struct {
	sync.RWMutex
	passId   string
	expireIn int64
}

func getPassId(ctx context.Context, fn func(ctx context.Context) (string, error)) (string, error) {
	mdPassIdOnce.Do(func() {
		mdPassIdInstance = &mdPassId{}
	})
	return mdPassIdInstance.GetPassId(ctx, fn)
}

func (m *mdPassId) GetPassId(ctx context.Context, fn func(ctx context.Context) (string, error)) (string, error) {
	m.RLock()
	passId := m.passId
	expireIn := m.expireIn
	m.RUnlock()
	if passId == "" || expireIn < time.Now().Unix() {
		var err error
		passId, err = fn(context.Background())
		if err != nil {
			return "", err
		}
		m.Lock()
		m.passId = passId
		m.expireIn = time.Now().Unix() + 86400*3
		m.Unlock()

	}
	return passId, nil
}

type usecase struct{}

func NewUsecase() domain.IMingdaoUsecase {
	return &usecase{}
}

func (u *usecase) GetPassId(ctx context.Context, param *domain.MingdaoGetPassIdReq) (string, error) {

	passid, err := getPassId(ctx, func(ctx context.Context) (string, error) {
		return u.getPassId(ctx, param)
	})

	return passid, err
}

func (u *usecase) getPassId(ctx context.Context, param *domain.MingdaoGetPassIdReq) (string, error) {

	pw, err := playwright.Run()
	if err != nil {
		app.Get().Logger().Error("playwright.Run", err)
		return "", err
	}
	defer pw.Stop()
	browser, err := pw.Chromium.Launch()
	if err != nil {
		app.Get().Logger().Error("pw.Chromium.Launch", err)
		return "", err
	}
	defer browser.Close()

	pwContext, err := browser.NewContext()
	if err != nil {
		app.Get().Logger().Error("browser.NewContext", err)
		return "", err
	}
	defer pwContext.Close()

	page, err := pwContext.NewPage()
	if err != nil {
		app.Get().Logger().Error("pwContext.NewPage", err)
		return "", err
	}
	defer page.Close()
	if _, err = page.Goto("https://z.daopuqifu.com/network?ReturnUrl=https%3A%2F%2Fz.daopuqifu.com%2Fdashboard"); err != nil {
		app.Get().Logger().Error("page.Goto", err)
		return "", err
	}

	page.Locator("#txtMobilePhone").Fill(param.UserName)
	page.Locator(".passwordIcon").Fill(param.Password)
	page.Locator(".btnForLogin").Click()

	if err := page.WaitForURL("https://z.daopuqifu.com/dashboard"); err != nil {
		app.Get().Logger().Error("page.WaitForURL", err)
		return "", err
	}
	cookies, err := pwContext.Cookies()
	if err != nil {
		app.Get().Logger().Error("pwContext.Cookies", err)
		return "", err
	}

	for _, cookie := range cookies {
		if cookie.Name == "md_pss_id" {
			app.Get().Logger().Info("found cookie md_pss_id", cookie.Value)
			return cookie.Value, nil
		}
	}
	app.Get().Logger().Error("not found cookie md_pss_id")
	return "", errors.New("not found cookie md_pss_id")
}
