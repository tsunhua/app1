package db

import (
	"app1/types"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DB_NAME = "app1.db"

var db *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open(DB_NAME), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 自动迁移模式（创建表格）
	db.AutoMigrate(&types.User{})
}

// 根据条件查询单个用户记录
func FindUserByName(name string) (user types.User, err error) {
	db.First(&user, "name = ?", name)
	return user, db.Error
}

// 创建用户记录
func CreateUser(user types.User) error {
	result := db.Create(&user)
	return result.Error
}
