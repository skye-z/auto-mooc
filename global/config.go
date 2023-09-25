package global

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const Version = "1.0.1"

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("ini")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			createDefault()
		} else {
			// 配置文件被找到，但产生了另外的错误
			fmt.Println(err)
		}
	}
}

func Set(key string, value interface{}) {
	viper.Set(key, value)
	viper.WriteConfig()
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func createDefault() {
	dir, _ := os.Getwd()
	// 安装状态
	viper.SetDefault("basic.install", "false")
	// 应用目录
	viper.SetDefault("basic.workspace", dir)
	// 服务端口
	viper.SetDefault("basic.port", "80")
	// 用户代理信息
	viper.SetDefault("basic.user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.5.2 Safari/605.1.15")
	// 调试模式
	viper.SetDefault("basic.debug", "false")
	// 慕课服务地址
	viper.SetDefault("mooc.path", "http://cce.org.uooconline.com")
	// 持久化存储位置
	viper.SetDefault("mooc.storage", dir+"/storage.db")
	// 是否已登录
	viper.SetDefault("mooc.login", "false")
	// 是否开启推送
	viper.SetDefault("push.enable", "false")
	// 推送地址
	viper.SetDefault("push.url", "https://api2.pushdeer.com/message/push?pushkey={替换你自己的令牌}&text=")
	viper.SafeWriteConfig()
}
