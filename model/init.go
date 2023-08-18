package model

import (
	"word-card-app/global"
)

// func init() {
// 	// if runMode := os.Getenv("RUN_MODE"); runMode == "testing" {
// 	if runMode := global.ServerSetting.RunMode; runMode == "testing" {
// 		// TO-DO
// 	} else {
// 		AutoMigrateAll()
// 	}
// }

// Migrate Model
func AutoMigrateAll() {
	migrateErr := global.GormDb.AutoMigrate(
		&Stock{},
	)
	if migrateErr != nil {
		panic("database migrate failed")
	}
}
