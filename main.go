package main

import (
	"log"
	"word-card-app/models"
	"word-card-app/router.go"

	"github.com/gin-gonic/gin"
)

// 全局变量，存储已撤销的 JWT 标识
// 这里是为了简化，正规项目当然是放到 Redis 或者数据库中
var revokedTokens = make(map[string]bool)

func RevokeToken(tokenID string) {
    revokedTokens[tokenID] = true
}

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
