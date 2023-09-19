package webkit

import (
	"auto-mooc/global"
	"io"
	"log"
	"os"

	"github.com/playwright-community/playwright-go"
)

// 启动参数
var RunOptions = &playwright.RunOptions{
	Browsers:            []string{"webkit"},
	SkipInstallBrowsers: false,
	Verbose:             true,
}

type WebKit struct {
	Engine  *playwright.Playwright
	Running bool
}

func InitWebKit() *WebKit {
	// 检查是否已安装环境
	installed := global.GetBool("basic.install")
	// 未安装
	if !installed {
		log.Println("[WebKit] No webkit detected, ready to install...")
		// 开始安装WebKit
		err := playwright.Install(RunOptions)
		// 安装出错
		if err != nil {
			log.Fatalf("[WebKit] Error installing: %v", err)
			return nil
		}
		global.Set("basic.install", true)
	}
	// 启动WebKit
	pw, err := playwright.Run(RunOptions)
	// 启动出错
	if err != nil {
		log.Fatalf("[WebKit] Error launching: %v", err)
		return nil
	}
	// 输出启动脚步位置
	log.Printf("[WebKit] Launches from: %s", pw.WebKit.ExecutablePath())
	// 创建持久化文件
	CreateStorage()
	return &WebKit{
		Engine:  pw,
		Running: false,
	}
}

func CreateStorage() {
	path := global.GetString("basic.workspace") + "/storage.db"
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("[WebKit] Create storage errir: %v", err)
	}
	data := "{\"cookies\":[],\"origins\":[]}"
	_, err = io.Writer.Write(file, []byte(data))
	if err != nil {
		log.Fatalf("[WebKit] Write storage errir: %v", err)
	}
	log.Printf("[WebKit] Create storage from: %s", path)
}
