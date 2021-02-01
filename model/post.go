package model

import (
	"gorm.io/gorm"
	"time"
)

// Post 指代 帖子相关的内容
type Post struct {
	ID int `gorm:"primaryKey;autoIncrement;not null"`
	// PostHistory   PostHistory
	UserInfoID int
	TUrl       string
	NumLike    int
	NumDislike int
	NumReport  int
	Status     state
	Comment    []Comment
	CreatedAt  time.Time      `gorm:"index"`
	UpdatedAt  time.Time      `gorm:"index"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

// // PostHistory 历史记录表 记录修改信息
// type PostHistory struct {
// 	PostID    int `gorm:"primaryKey;index"`
// 	OUrl      string
// 	CreatedAt time.Time      `gorm:"index"`
// 	UpdatedAt time.Time      `gorm:"index"`
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }
