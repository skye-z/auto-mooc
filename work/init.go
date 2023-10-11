package work

import (
	"io"
	"log"
	"os"

	"github.com/skye-z/auto-mooc/global"

	"github.com/playwright-community/playwright-go"
)

// 启动参数
var RunOptions = &playwright.RunOptions{
	Browsers:            []string{"webkit"},
	SkipInstallBrowsers: false,
	Verbose:             true,
}

func InitWork() *global.RunPKG {
	engine := global.GetString("basic.engine")
	if global.GetString("basic.engine") != "webkit" {
		RunOptions.Browsers = []string{"chromium"}
	}
	// 检查是否已安装环境
	installed := global.GetBool("basic.install")
	// 未安装
	if !installed {
		log.Println("[Work] No " + engine + " detected, ready to install...")
		// 开始安装浏览器
		err := playwright.Install(RunOptions)
		// 安装出错
		if err != nil {
			log.Fatalf("[Work] Error installing: %v", err)
			return nil
		}
		// 创建持久化文件
		CreateStorage()
		global.Set("basic.install", true)
	}
	// 启动浏览器
	pw, err := playwright.Run(RunOptions)
	// 启动出错
	if err != nil {
		log.Fatalf("[Work] Error launching: %v", err)
		return nil
	}
	// 输出启动脚步位置
	if global.GetString("basic.engine") == "webkit" {
		log.Printf("[Work] Launches from: %s", pw.WebKit.ExecutablePath())
	} else {
		log.Printf("[Work] Launches from: %s", pw.Chromium.ExecutablePath())
	}
	return &global.RunPKG{
		Engine:  pw,
		Running: false,
	}
}

func CreateStorage() {
	path := global.GetString("basic.workspace") + "/storage.db"
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("[Work] Create storage errir: %v", err)
	}
	data := "{\"cookies\":[],\"origins\":[]}"
	_, err = io.Writer.Write(file, []byte(data))
	if err != nil {
		log.Fatalf("[Work] Write storage errir: %v", err)
	}
	log.Printf("[Work] Create storage from: %s", path)
}
