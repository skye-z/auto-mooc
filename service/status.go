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
