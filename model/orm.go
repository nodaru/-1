package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//CreateDB will Create DB
func CreateDB() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&User{}, &UserInfo{}, &Post{}, &PostAgreement{}, &PostHistory{},
		&Comment{}, &CommentAgreement{}, &CommentHistory{})
}
