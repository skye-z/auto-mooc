package service

import (
	"auto-mooc/global"
	"auto-mooc/webkit"
	"log"
	"strings"
	"time"

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
	session.Page.Context().StorageState(global.GetString("mooc.storage"))
	// 检查登录
	next := ms.checkLogin(ctx, host, session.Page)
	if !next {
		session.Page.Context().StorageState(global.GetString("mooc.storage"))
		ms.Close(session)
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
	session.Page.Context().StorageState(global.GetString("mooc.storage"))
	log.Println("[Mooc] Login successful")
	global.Set("mooc.login", "true")
	ms.Close(session)
}

// 检查登录
func (ms MoocService) checkLogin(ctx *gin.Context, host string, page playwright.Page) bool {
	header, _ := page.Locator(".layout-header-right").TextContent()
	if strings.Contains(header, "登录") {
		global.Set("mooc.login", "false")
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
		global.Set("mooc.login", "true")
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

// 关闭会话
func (ms MoocService) Close(session *webkit.Session) {
	time.Sleep(1 * time.Second)
	session.Page.Close()
	session.Context.Close()
	session.Browser.Close()
}
