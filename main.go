package main

import (
	"embed"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/skye-z/auto-mooc/global"
	"github.com/skye-z/auto-mooc/service"
	"github.com/skye-z/auto-mooc/webkit"

	"github.com/gin-gonic/gin"
)

//go:embed page/*
var page embed.FS

func main() {
	// 初始化配置
	global.InitConfig()
	// 初始化WebKit
	WebKit := webkit.InitWebKit()
	if len(os.Args) == 1 {
		// 启动Http服务
		RunHttp(WebKit)
	} else {
		os.Exit(0)
	}
}

// 启动Http服务
func RunHttp(obj *webkit.WebKit) {
	// 关闭调试
	gin.SetMode(gin.ReleaseMode)
	// 禁用路由日志
	if !global.GetBool("basic.debug") {
		gin.DefaultWriter = io.Discard
	}
	log.Println("[Http] Route registration")
	// 创建路由
	route := gin.Default()
	// 配置静态页面
	route.NoRoute(func(ctx *gin.Context) {
		data, err := page.ReadFile("page/index.html")
		if err != nil {
			ctx.HTML(http.StatusOK, "404.html", gin.H{
				"title": "404",
			})
			return
		}
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})
	// 创建状态服务
	statusService := &service.StatusService{
		WebKitObj: obj,
	}
	// 创建慕课服务
	moocService := &service.MoocService{
		WebKitObj: obj,
	}
	// 接口 查询状态
	route.GET("/status", statusService.GetStatus)
	// 接口 获取任务截图
	route.GET("/screenshot", statusService.GetScreenshot)
	// 接口 登录账户
	route.GET("/login", moocService.Login)
	// 接口 课程列表
	route.GET("/class/list", moocService.ClassList)
	// 接口 提交选课
	route.GET("/class/select", moocService.ClassSelect)
	// 接口 开始上课
	route.GET("/class/start", moocService.StartClass)
	// 接口 停止上课
	route.GET("/class/stop", moocService.StopClass)

	// 获取端口号
	port := getPort()
	log.Println("[Http] Service started, port is", port)
	route.Run(":" + port)
}

// 获取端口号
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = global.GetString("basic.port")
	}
	if port == "" {
		port = "80"
	}
	return port
}
