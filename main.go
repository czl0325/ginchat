package main

import (
	"fmt"
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitSystemConfig()
	utils.InitDB()
	r := router.Router()
	err := r.Run()
	if err != nil {
		fmt.Println("服务器启动失败", err)
	}
}
