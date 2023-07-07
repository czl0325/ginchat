package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitSystemConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置文件失败", err)
		return
	}
}
