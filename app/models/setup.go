package models

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DbEIS *gorm.DB

func ConnectDatabase() *gorm.DB {
	dns := viper.GetString("DB_URL")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	DbEIS = db
	return db
}
