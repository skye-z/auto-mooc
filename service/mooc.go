package service

import (
	"auto-mooc/global"
	"auto-mooc/webkit"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/playwright-community/playwright-go"
)

type MoocService struct {
	WebKitObj *webkit.WebKit
}

type Class struct {
	Id   int64
	Name string
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
	if !ms.getLoginStatus(page) {
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

// 获取课程列表
func (ms MoocService) ClassList(ctx *gin.Context) {
	host := global.GetString("mooc.path")
	// 打开页面
	session, err := webkit.OpenPage(ms.WebKitObj.Engine, host+"/home")
	if err != nil {
		log.Fatalf("无法打开页面: %v", err)
	}
	// 检查登录
	if !ms.getLoginStatus(session.Page) {
		global.ReturnMessage(ctx, false, "请先登录")
		return
	}
	// 获取课程列表
	classList, err := session.Page.Locator(".course-item").All()
	if err != nil || len(classList) == 0 {
		global.ReturnMessage(ctx, false, "没有可选课程")
		return
	}
	var list []Class
	for i := 0; i < len(classList); i++ {
		item := classList[i].Locator(".course-name")
		href, _ := item.GetAttribute("href")
		classId, _ := strconv.ParseInt(href[strings.Index(href, "cycleid=")+8:], 10, 64)
		className, _ := item.TextContent()
		classInfo := &Class{
			Id:   classId,
			Name: className,
		}
		list = append(list, *classInfo)
	}
	global.ReturnData(ctx, true, "请选择课程", list)
}

// 选课
func (ms MoocService) ClassSelect(ctx *gin.Context) {
	id := ctx.Query("id")
	if len(id) == 0 {
		global.ReturnMessage(ctx, false, "请传入课程编号")
	} else {
		clssId, _ := strconv.ParseInt(id, 10, 64)
		global.Set("mooc.class", clssId)
		global.ReturnMessage(ctx, true, "选课已登记")
	}
}

// 开始上课
func (ms MoocService) StartClass(ctx *gin.Context) {
	// 检查登录
	if !global.GetBool("mooc.login") {
		global.ReturnMessage(ctx, false, "请先登录")
		return
	}
	// 检查选课
	classId := global.GetString("mooc.class")
	if len(classId) == 0 {
		global.ReturnMessage(ctx, false, "请先完成选课")
		return
	}
	host := global.GetString("mooc.path")
	// 打开页面
	session, err := webkit.OpenPage(ms.WebKitObj.Engine, host+"/home/learn/index#/"+classId+"/go")
	if err != nil {
		log.Fatalf("无法打开页面: %v", err)
	}
	state := webkit.CreateWork(session, ms.WebKitObj)
	if !state {
		global.ReturnMessage(ctx, false, "正在上课中")
		return
	}
	global.ReturnMessage(ctx, true, "上课开始")
}

// 结束上课
func (ms MoocService) StopClass(ctx *gin.Context) {
	state := webkit.CloseWork(ms.WebKitObj)
	if !state {
		global.ReturnMessage(ctx, false, "未在上课中")
		return
	}
	global.ReturnMessage(ctx, true, "上课结束")
}

// 获取登录状态
func (ms MoocService) getLoginStatus(page playwright.Page) bool {
	header, _ := page.Locator(".layout-header-right").TextContent()
	return !strings.Contains(header, "登录")
}

// 关闭会话
func (ms MoocService) Close(session *webkit.Session) {
	time.Sleep(1 * time.Second)
	session.Page.Close()
	session.Context.Close()
	session.Browser.Close()
}
