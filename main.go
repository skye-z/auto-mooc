package main

import (
	"auto-mooc/global"
	"auto-mooc/service"
	"auto-mooc/webkit"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	global.InitConfig()
	// 初始化WebKit
	WebKit := webkit.InitWebKit()
	// 启动Http服务
	RunHttp(WebKit)
}

// 启动Http服务
func RunHttp(obj *webkit.WebKit) {
	// 关闭调试
	gin.SetMode(gin.ReleaseMode)
	log.Println("[Http] Route registration")
	// 创建路由
	route := gin.Default()

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
	// 接口 登录账户
	route.GET("/login", moocService.Login)
	// 接口 选课
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
