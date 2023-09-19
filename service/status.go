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
	if ss.WebKitObj.Running {
		global.ReturnMessage(ctx, true, "服务运行中")
	} else {
		global.ReturnMessage(ctx, true, "服务尚未启动")
	}
}
