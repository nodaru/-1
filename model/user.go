package model

import (
	"gorm.io/gorm"
	"time"
)

//UserPrivilege 代表着用户的权限
type UserPrivilege int16

/*
baned 指代被禁止访问的用户 其无法注销，无法浏览，也无法发言
muted 指代被禁言的用户 其权限与匿名用户相同
admin 指代最高级的管理员 拥有最高权限，可以恢复被删除的内容
manager 指代一般的管理员 可以处理举报，删除发言，有限次数的封禁用户/每三天，解封被自身封禁的用户
normalUser 指代正常用户 可以正常的使用功能
anonymous 指代未登录的用户 与banned用户不同是其未登录
*/

const (
	banedUserPrivilege     = -100
	mutedUserPrivilege     = -1
	adminPrivilege         = 0
	managerPrivilege       = 1
	normalUserPrivilege    = 10
	anonymousUserPrivilege = 100
)


// User 表 包含着 参与讨论的对象
type User struct {
	ID            int `gorm:"primaryKey;autoIncrement;not null"`
	Privilege     int
	EncryptedPass string
	Salt          string
	UserInfo      UserInfo
	CreatedAt      time.Time      `gorm:"index"`
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
	CreatedAt   time.Time      `gorm:"index"`
	UpdatedAt  time.Time      `gorm:"index"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
