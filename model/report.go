package model

import (
	"gorm.io/gorm"
	"time"
)


//Report 是举报相关的表
type Report struct {
	ID              int  `gorm:"primaryKey;autoIncrement;not null"`
	ReportUser      User `gorm:"foreignKey:ReportUserID"`
	ReportUserID    int
	ReportPost      Post `gorm:"foreignKey:ReportPostID"`
	ReportPostID    int
	ReportComment   Comment `gorm:"foreignKey:ReportCommentID"`
	ReportCommentID int
	ReportedUser    User `gorm:"foreignKey:ReportedUserID"`
	ReportedUserID  int
	ReportType      ReportReasonKeyType
	//ReportReason 是用户提交的附加信息
	ReportReason string
	BanID        int
	CreatedAt    time.Time      `gorm:"index"`
	UpdatedAt    time.Time      `gorm:"index"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

//Edit 是与编辑发言相关的表
type Edit struct {
	// 删除者
	ID              int
	EditUser        User
	EditUserID      int
	EditPost        Post
	EditPostID      int
	EditComment     Comment
	EditCommentID   int
	EditedUser      User
	EditedUserID    int
	HandleReasonKey int
	CreatedAt       time.Time      `gorm:"index"`
	UpdatedAt       time.Time      `gorm:"index"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

//Ban 是与封禁用户相关的表
type Ban struct {
	ID          int
	BanUser     User
	BanUserID   int
	BanedUser   User
	BanedUserID int
	Report      []Report
	BanReason   string
	ExpireAt    time.Time
	CreatedAt   time.Time      `gorm:"index"`
	UpdatedAt   time.Time      `gorm:"index"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
