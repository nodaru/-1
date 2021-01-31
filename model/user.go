package model

import (
	"gorm.io/gorm"
	"time"
)

// User 表 包含着 参与讨论的对象
type User struct {
	ID            int `gorm:"primaryKey;autoIncrement;not null"`
	Privilege     int
	EncryptedPass string
	Salt          string
	UserInfo      UserInfo
	CreatedAt     time.Time      `gorm:"index"`
	UpdatedAt     time.Time      `gorm:"index"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

// UserInfo 表 包含着 User相关的信息
type UserInfo struct {
	ID         int `gorm:"primaryKey;autoIncrement;not null"`
	UserID     int
	NiceName   string
	AvatarPath string
	Contact    string
	Post       []Post
	Comment    []Comment
	Birthday   time.Time
	CreatedAt  time.Time      `gorm:"index"`
	UpdatedAt  time.Time      `gorm:"index"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
