package global

import "github.com/playwright-community/playwright-go"

type RunPKG struct {
	Engine  *playwright.Playwright
	Running bool
	Error   string
}
