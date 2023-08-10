package main

import (
	"log"
	"word-card-app/models"
)

func main() {
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

}
