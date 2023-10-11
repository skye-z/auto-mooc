package service

import (
	"github.com/skye-z/auto-mooc/global"
	"github.com/skye-z/auto-mooc/work"

	"github.com/gin-gonic/gin"
)

type StatusService struct {
	PKG *global.RunPKG
}

func (ss StatusService) GetStatus(ctx *gin.Context) {
	isLogin := global.GetBool("mooc.login")
	loginTips := "已登录"
	if !isLogin {
		loginTips = "未登录"
	}
	classId := global.GetString("mooc.class")
	if len(classId) != 0 {
		loginTips += ", 已选课(" + classId + ")"
	}
	if ss.PKG.Running {
		global.ReturnMessage(ctx, true, loginTips+", 任务执行中")
	} else {
		global.ReturnMessage(ctx, true, loginTips+", 任务未启动")
	}
}

func (ss StatusService) GetScreenshot(ctx *gin.Context) {
	if !ss.PKG.Running {
		global.ReturnMessage(ctx, false, "服务尚未启动")
		return
	}
	data, err := work.WorkScreenshot()
	if err != nil {
		global.ReturnMessage(ctx, false, "获取任务截图失败")
	}
	ctx.Writer.Write(data)
	ctx.Abort()
}
