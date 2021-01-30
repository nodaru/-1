package model

import (
	"gorm.io/gorm"
	"time"
)

//SysNoticeType 是系统通知的类型
var SysNoticeType map[int]string = map[int]string{
	1:"新增回复",
	2:"系统通知",
}

//SysNotice 是系统推送的通知
type SysNotice struct{
	ID int
	User User
	UserID int
	SysNoticeTypeKey int
	SysNoticeContent string
	IsRead bool
	CreateAt     time.Time      `gorm:"index"`
	UpdatedAt    time.Time      `gorm:"index"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
