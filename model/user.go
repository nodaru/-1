package model

import (
	"111/util"
	"crypto/rand"
	"encoding/hex"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// User 表 包含着 参与讨论的对象
type User struct {
	ID            int `gorm:"primaryKey;autoIncrement;not null"`
	Privilege     int
	EncryptedPass string
	Salt          string
	NumReport     int
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

//Agreement 记录帖子被赞赏 举报 喜爱的次数
type Agreement struct {
	ID        int `gorm:"primaryKey;index"`
	User      User
	UserID    int
	Post      Post
	PostID    int
	Comment   Comment
	CommentID int
	CreatedAt time.Time      `gorm:"index"`
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

//CreateNormalUser will create a normal user
func CreateNormalUser(name string, email string, passwordEncyrpt string) (user *User, err error) {
	salt := make([]byte, 20)
	_, _err := rand.Read(salt)
	if _err != nil {
		panic(_err)
	}
	passwordEncyrpt = util.HashPass(passwordEncyrpt, hex.EncodeToString(salt))
	user = &User{
		Privilege:     normalUserPrivilege,
		Salt:          hex.EncodeToString(salt),
		EncryptedPass: passwordEncyrpt,
		UserInfo: UserInfo{
			NiceName: name,
			Contact:  email,
		},
	}
	result := db.Create(user)
	err = result.Error
	return
}

//GetUserByID will get a user struct by uid
func GetUserByID(uid int) (user *User, err error) {
	user = &User{}
	result := db.Find(&user, uid)
	if user.ID == 0 || result.RowsAffected == 0 {
		err = ERR_USER_DOES_NOT_EXISTED
	} else {
		err = db.Preload("UserInfo.Post").Preload("UserInfo.Comment").Preload(clause.Associations).Find(&user).Error
	}
	return
}