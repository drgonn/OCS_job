package main

import (
	"log"
	"word-card-app/models"
	"word-card-app/router.go"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.SetupRoutes(r)

	// 初始化数据库连接
	err := models.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// 创建数据库表
	err = models.CreateTables()
	if err != nil {
		log.Fatal("Failed to create tables:", err)
	}
	r.Run(":38085")
}
