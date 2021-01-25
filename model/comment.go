package model

import (
	"gorm.io/gorm"
	"time"
)

// Comment 指代 帖子相关的内容
type Comment struct {
	ID               int
	Curl             string
	PostID           int
	UserInfoID       int
	CreateAt         time.Time      `gorm:"index"`
	UpdatedAt        time.Time      `gorm:"index"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	Status           state
	CommentID        *int
	Comment          *Comment
	CommentAgreement CommentAgreement `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CommentHistory   CommentHistory   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

//CommentAgreement 记录帖子被赞赏 举报 喜爱的次数
type CommentAgreement struct {
	CommentID  int
	NumLike    int
	NumDislike int
	NumReport  int
	CreateAt   time.Time      `gorm:"index"`
	UpdatedAt  time.Time      `gorm:"index"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

// CommentHistory 历史记录表 记录修改信息
type CommentHistory struct {
	CommentID int
	OUrl      string
	CreateAt  time.Time      `gorm:"index"`
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}