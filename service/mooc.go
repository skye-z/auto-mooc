package service

import (
	"auto-mooc/global"
	"auto-mooc/webkit"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/playwright-community/playwright-go"
)

type MoocService struct {
	WebKitObj *webkit.WebKit
}

// 登录账户
func (ms MoocService) Login(ctx *gin.Context) {
	host := global.GetString("mooc.path")
	// 打开页面
	session, err := webkit.OpenPage(ms.WebKitObj.Engine, host)
	if err != nil {
		log.Fatalf("无法打开页面: %v", err)
	}
	// 检查登录
	next := ms.checkLogin(ctx, host, session.Page)
	if !next {
		return
	}
	ch := make(chan bool)
	// 监听扫码登录跳转
	session.Page.On("framenavigated", func(frame playwright.Frame) {
		if frame.URL() == (host + "/home") {
			ch <- true
			return
		}
	})
	<-ch
	session.Page.Close()
	session.Context.Close()
	session.Browser.Close()
}

func (ms MoocService) checkLogin(ctx *gin.Context, host string, page playwright.Page) bool {
	header, _ := page.Locator(".layout-header-right").TextContent()
	if strings.Contains(header, "登录") {
		if _, err := page.Goto(host + "/oauth/login/weixin"); err != nil {
			log.Fatalf("无法跳转地址: %v", err)
		}
		qrCode, err := page.Locator(".web_qrcode_img").Screenshot()
		if err != nil {
			global.ReturnMessage(ctx, false, "慕课网站登录服务故障")
			return false
		} else {
			ctx.Writer.Write(qrCode)
			return true
		}
	} else {
		global.ReturnMessage(ctx, false, "已登录,请勿重复操作")
		return false
	}
}

// 选课
func (ms MoocService) ClassSelect(ctx *gin.Context) {

}

// 开始上课
func (ms MoocService) StartClass(ctx *gin.Context) {

}

// 结束上课
func (ms MoocService) StopClass(ctx *gin.Context) {

}
