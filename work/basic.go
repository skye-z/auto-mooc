package work

import (
	"github.com/skye-z/auto-mooc/global"

	"github.com/playwright-community/playwright-go"
)

type Session struct {
	Browser playwright.Browser
	Context playwright.BrowserContext
	Page    playwright.Page
}

func OpenPage(engine *playwright.Playwright, url string) (*Session, error) {
	storagePath := global.GetString("mooc.storage")
	var (
		browser playwright.Browser
		err     error
	)
	// 启动浏览器
	if global.GetString("basic.engine") == "webkit" {
		browser, err = engine.WebKit.Launch(playwright.BrowserTypeLaunchOptions{
			Headless: playwright.Bool(!global.GetBool("basic.debug")),
		})
	} else {
		browser, err = engine.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
			Headless: playwright.Bool(!global.GetBool("basic.debug")),
		})
	}
	if err != nil {
		return nil, err
	}
	// 创建上下文
	context, err := browser.NewContext(playwright.BrowserNewContextOptions{
		StorageStatePath: playwright.String(storagePath),
		UserAgent:        playwright.String(global.GetString("basic.user-agent")),
	})
	if err != nil {
		return nil, err
	}
	// 创建新页面
	page, err := context.NewPage()
	if err != nil {
		return nil, err
	}
	// 页面跳转
	if _, err = page.Goto(url); err != nil {
		return nil, err
	}
	// 构建会话
	session := &Session{
		Browser: browser,
		Context: context,
		Page:    page,
	}
	return session, nil
}
