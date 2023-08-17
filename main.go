package main

import (
	"log"
	"word-card-app/global"
	"word-card-app/router.go"

	"github.com/gin-gonic/gin"
)

func main() {
	// 第一步，加载设置
	err := global.SetupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	r := gin.Default()
	router.SetupRoutes(r)

	// 初始化数据库连接
	// err := model.InitDB()
	// if err != nil {
	// 	log.Fatal("Failed to initialize database:", err)
	// }

	// 创建数据库表

	r.Run(":38085")
}
