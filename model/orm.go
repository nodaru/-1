package model

import (
	"111/util"
	"crypto/rand"
	"encoding/hex"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var db *gorm.DB

func init() {
	InitDB()
	MigrateDB()
}

//InitDB will init db
func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("D:/Lux/src/go/-1/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return
}

//MigrateDB will migrate DB
func MigrateDB() {
	// 迁移 schema
	db.AutoMigrate(&User{}, &UserInfo{}, &Post{}, &PostAgreement{}, &PostHistory{},
		&Comment{}, &CommentAgreement{}, &CommentHistory{}, &Report{}, &Edit{}, &Ban{}, &SysNotice{})
}

//TODO： 完成逻辑

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

//GetPostByID will get a post by pid
func GetPostByID(pid int) (post *Post, err error) {
	post = &Post{}
	result := db.Find(&post, pid)
	if post.ID == 0 || result.RowsAffected == 0 {
		err = ERR_POST_DOES_NOT_EXISTED
	} else {
		err = db.Preload("Comment.RefComment").Preload(clause.Associations).Find(post, pid).Error
	}
	return
}

//GetCommentByID will get a comment by cid
func GetCommentByID(cid int) (comment *Comment, err error) {
	comment = &Comment{}
	result := db.Preload(clause.Associations).Find(&comment, cid)
	if comment.ID == 0 || result.RowsAffected == 0 {
		err = ERR_COMMENT_DOES_NOT_EXISTED
	} else {
		err = result.Error
	}
	return
}

//CreatePost will create a post
func (user *User) CreatePost(pathToFile string) (*Post, error) {
	post := &Post{
		UserInfoID: user.UserInfo.ID,
		TUrl:       pathToFile,
		Status:     state_normal,
	}
	result := db.Create(&post)
	err := result.Error
	return post, err
}

//CreateComment will create a comment
func (user *User) CreateComment(pathToFile string, post *Post, referComment *Comment) (*Comment, error) {
	comment := &Comment{
		UserInfoID:   user.UserInfo.ID,
		Curl:         pathToFile,
		PostID:       post.ID,
		RefComment:   referComment,
		RefCommentID: &referComment.ID,
		Status:       state_normal,
	}
	result := db.Create(comment)
	err := result.Error
	return comment, err
}
