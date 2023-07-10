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

func InitDB() {
	//db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{})
	//if err != nil {
	//	fmt.Println("数据库连接失败", err)
	//	return
	//}
}
