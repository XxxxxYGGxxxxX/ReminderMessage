package models

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化数据库连接，数据库配置信息
var DB = Init()

func Init() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/message?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("gorm Init Error:", err)
	}

	return db

}
