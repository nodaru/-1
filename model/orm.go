package model

import (
	"111/util"
	"crypto/rand"
	"encoding/hex"
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var db *gorm.DB

func init(){
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

//TODO： 完成业务逻辑

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
	result := db.Find(&user,uid)
	if user.ID == 0 || result.RowsAffected == 0{
		err = errors.New("user does not existed")
	}else{
		err = db.Preload("UserInfo.Post").Preload("UserInfo.Comment").Preload(clause.Associations).Find(&user).Error
	}
	return
}

//GetPostByID will get a post by pid
func GetPostByID(pid int) (post *Post,err error){
post = &Post{}
result := db.Find(&post,pid)
if post.ID == 0 || result.RowsAffected == 0{
	err = errors.New("post does not existed")
}else{
	err = db.Preload(clause.Associations).Find(post,pid).Error
}
return
}

//CreatePost will create a post
func (user *User) CreatePost(pathToFile string) (*Post, error){
	post := &Post{
		UserInfoID: user.UserInfo.ID,
		TUrl: pathToFile,
		Status: normal,
	}
	db.Create(&post)
	err := db.Error
	return post,err
}

//CreateComment will create a comment
func (user *User) CreateComment(pathToFile string,post *Post) (*Comment, error){
	comment := &Comment{
		UserInfoID: user.UserInfo.ID,
		Curl: pathToFile,
		PostID: post.ID,
	}
	//TODO: finins this func
	return comment,nil
}