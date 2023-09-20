package service

import (
	"auto-mooc/global"
	"auto-mooc/webkit"

	"github.com/gin-gonic/gin"
)

type StatusService struct {
	WebKitObj *webkit.WebKit
}

func (ss StatusService) GetStatus(ctx *gin.Context) {
	isLogin := global.GetBool("mooc.login")
	loginTips := "已登录"
	if !isLogin {
		loginTips = "未登录"
	}
	if ss.WebKitObj.Running {
		global.ReturnMessage(ctx, true, loginTips+",服务运行中")
	} else {
		global.ReturnMessage(ctx, true, loginTips+",服务尚未启动")
	}
}

func (ss StatusService) GetScreenshot(ctx *gin.Context) {
	if !ss.WebKitObj.Running {
		global.ReturnMessage(ctx, false, "服务尚未启动")
		return
	}
	data, err := webkit.WorkScreenshot()
	if err != nil {
		global.ReturnMessage(ctx, false, "获取任务截图失败")
	}
	ctx.Writer.Write(data)
	ctx.Abort()
}
