package model

import (
	"gorm.io/gorm"
	"time"
)

// Comment 指代 帖子相关的内容
type Comment struct {
	ID         int `gorm:"primaryKey;autoIncrement;not null"`
	Curl       string
	PostID     int
	UserInfoID int
	NumLike    int
	NumDislike int
	NumReport  int
	Status     state
	//自引用的Comment
	RefCommentID *int
	RefComment   *Comment
	// CommentHistory   CommentHistory   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time      `gorm:"index"`
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// // CommentHistory 历史记录表 记录修改信息
// type CommentHistory struct {
// 	CommentID int `gorm:"primaryKey;index"`
// 	OUrl      string
// 	CreatedAt time.Time      `gorm:"index"`
// 	UpdatedAt time.Time      `gorm:"index"`
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }
