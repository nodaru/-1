package model

import (
	"gorm.io/gorm"
	"time"
)

//SysNotice 是系统推送的通知
type SysNotice struct {
	ID               int
	User             User
	UserID           int
	SysNoticeTypeKey int
	SysNoticeContent string
	IsRead           bool
	CreatedAt        time.Time      `gorm:"index"`
	UpdatedAt        time.Time      `gorm:"index"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

//TODO 优化消息推送

//CreateSysNotice will create a sys notice
func CreateSysNotice()(sys *SysNotice,err error){
	sys = &SysNotice{}
return
}