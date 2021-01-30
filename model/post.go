package model

import (
	"gorm.io/gorm"
	"time"
)

type state int

const (
	nice    state = 1
	normal  state = 0
	deleted state = -2
	edited  state = -1
	pined   state = 2
)

// Post 指代 帖子相关的内容
type Post struct {
	ID            int `gorm:"primaryKey;autoIncrement;not null"`
	PostHistory   PostHistory
	PostAgreement PostAgreement
	UserInfoID    int
	TUrl          string
	Status        state
	Comment       []Comment
	CreatedAt     time.Time      `gorm:"index"`
	UpdatedAt     time.Time      `gorm:"index"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

// PostHistory 历史记录表 记录修改信息
type PostHistory struct {
	PostID    int `gorm:"primaryKey;index"`
	OUrl      string
	CreatedAt time.Time      `gorm:"index"`
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

//PostAgreement 记录帖子被赞赏 举报 喜爱的次数
type PostAgreement struct {
	PostID     int `gorm:"primaryKey;index"`
	NumLike    int
	NumDislike int
	NumReport  int
	CreatedAt  time.Time      `gorm:"index"`
	UpdatedAt  time.Time      `gorm:"index"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
