package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBSqlite *gorm.DB

func InitDBSqlite() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移模式
	DB.AutoMigrate(&User{})
}

type UserSqlite struct {
	gorm.Model
	Name  string
	Email string
}
