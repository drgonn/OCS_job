package main

import (
	"log"
	"word-card-app/db"
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
	// fmt.Println(global.ServerSetting)
	// fmt.Println(global.DatabaseSetting)
	// fmt.Println(global.AppSetting)

	// 第二步，初始化日志

	// 第三步，初始化数据库
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
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

func setupDBEngine() error {
	var err error
	global.GormEngine, err = db.NewGormClient(global.DatabaseSetting)
	if err != nil {
		return err
	}

	// model.Migrate(global.GormEngine)

	return nil
}
