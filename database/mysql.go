package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// MySQL 连接字符串格式: "user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:root@tcp(127.0.0.1:3306)/wanworld?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移模式
	DB.AutoMigrate(&User{})
}

type User struct {
	gorm.Model
	Name   string `gorm:"size:20"`
	Email  string `gorm:"size:255"`
	Status uint   `gorm:"type:tinyint"`
	Type   uint   `gorm:"type:smallint"`
}
