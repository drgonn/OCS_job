package db

import (
	"errors"
	"fmt"
	"ocs-app/global"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewGormClient(config global.DatabaseSettingS) (*gorm.DB, error) {
	switch config.Source {
	case "postgres":
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Shanghai",
			config.Host, config.Port, config.User, config.Dbname,
			config.Password,
		)

		gdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: logger.Default.LogMode(logger.Info),
		},
		)

		return gdb, err
	default:
		return nil, errors.New("没有匹配的数据库")
	}
}
