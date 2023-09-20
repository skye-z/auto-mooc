package global

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const Version = "1.0.0"

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
	// 慕课服务地址
	viper.SetDefault("mooc.path", "http://cce.org.uooconline.com")
	// 持久化存储位置
	viper.SetDefault("mooc.storage", dir+"/storage.db")
	viper.SafeWriteConfig()
}
