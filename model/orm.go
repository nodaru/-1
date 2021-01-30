package model

import (
	"111/util"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init(){
	InitDB()
	MigrateDB()
}

//InitDB will init db
func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
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
func CreateNormalUser(name string, email string, passwordEncyrpt string) (user User, err error) {
	if isEmailExsisted(email){
		return User{}, errors.New("email have been regisisted")
	}
	salt := make([]byte, 20)
	_, _err := rand.Read(salt)
	if _err != nil {
		panic(_err)
	}
	passwordEncyrpt = util.HashPass(passwordEncyrpt, hex.EncodeToString(salt))
	user = User{
		Privilege:     normalUserPrivilege,
		Salt:          hex.EncodeToString(salt),
		EncryptedPass: passwordEncyrpt,
		UserInfo: UserInfo{
			NiceName: name,
			Contact:  email,
		},
	}
	db.Create(&user)  //TODO email conflict logic
	err = db.Error
	return
}
func isEmailExsisted(email string) bool {
	userinfo := UserInfo{}
	result :=db.Where("Contact=?",email).Find(&userinfo)
	if userinfo.ID == 0 || errors.Is(result.Error,gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

//GetUserByID will get a user struct by uid
func GetUserByID(uid int) (user User, err error) {
	user = User{}
	db.Find(&user,uid)
	fmt.Println(user.ID)
	if user.ID == 0{
		panic("User does not exist")
	}
	err = db.Model(&user).Association("UserInfo").Find(&user.UserInfo)
	return 
}


//CreatPost will create a post
func CreatPost(uid int,path string) (Post, error){
	user,err:= GetUserByID(uid)
	if err != nil {
		panic(err)
	}
	post := Post{
		UserInfoID: user.UserInfo.ID,
		TUrl: path,
		Status: normal,
	}
	db.Create(&post)
	err = db.Error
	return post,err
}