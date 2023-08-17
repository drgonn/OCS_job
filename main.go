package main

import (
	"log"
	model "word-card-app/model"
	"word-card-app/router.go"

	"github.com/gin-gonic/gin"
)

func main() {
	// 第一步，加载设置


	r := gin.Default()
	router.SetupRoutes(r)

	// 初始化数据库连接
	err := model.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// 创建数据库表

	r.Run(":38085")
}
