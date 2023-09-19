package webkit

import (
	"auto-mooc/global"

	"github.com/playwright-community/playwright-go"
)

type Session struct {
	Browser playwright.Browser
	Context playwright.BrowserContext
	Page    playwright.Page
}

func OpenPage(engine *playwright.Playwright, url string) (*Session, error) {
	storagePath := global.GetString("basic.workspace") + "/storage.db"
	browser, err := engine.WebKit.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		return nil, err
	}
	// 创建上下文
	context, err := browser.NewContext(playwright.BrowserNewContextOptions{
		StorageStatePath: playwright.String(storagePath),
	})
	// 配置持久化存储
	context.StorageState(storagePath)
	if err != nil {
		return nil, err
	}
	page, err := context.NewPage()
	if err != nil {
		return nil, err
	}
	page.Context().AddCookies([]playwright.OptionalCookie{})
	if _, err = page.Goto(url); err != nil {
		return nil, err
	}
	session := &Session{
		Browser: browser,
		Context: context,
		Page:    page,
	}
	return session, nil
}
