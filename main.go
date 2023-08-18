package main

import (
	"log"
	"ocs-app/db"
	"ocs-app/global"
	"ocs-app/model"
	"ocs-app/router.go"

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

	r.Run(":" + global.ServerSetting.HttpPort)
}

func setupDBEngine() error {
	var err error
	global.GormDb, err = db.NewGormClient(global.DatabaseSetting)
	if err != nil {
		return err
	}

	// 创建数据库表
	model.AutoMigrateAll()

	return nil
}
