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
	storagePath := global.GetString("mooc.storage")
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
	if err != nil {
		return nil, err
	}
	page, err := context.NewPage()
	if err != nil {
		return nil, err
	}
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
