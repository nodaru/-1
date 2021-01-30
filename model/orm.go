package model

import (
	"fmt"
	"111/util"
	"crypto/rand"
	"encoding/hex"
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

//MigrateDB will Create DB
func MigrateDB() {
	// 迁移 schema
	db.AutoMigrate(&User{}, &UserInfo{}, &Post{}, &PostAgreement{}, &PostHistory{},
		&Comment{}, &CommentAgreement{}, &CommentHistory{}, &Report{}, &Edit{}, &Ban{}, &SysNotice{})
}

//TODO： 完成业务逻辑

//CreateNormalUser will create a normal user
func CreateNormalUser(name string, email string, passwordEncyrpt string) (user User, err error) {
	salt := make([]byte, 20)
	_, _err := rand.Read(salt)
	if _err != nil {
		panic(_err)
	}
	passwordEncyrpt = util.HashPass(passwordEncyrpt, hex.EncodeToString(salt))
	fmt.Println(passwordEncyrpt)
	user = User{
		Privilege:     normalUserPrivilege,
		Salt:          hex.EncodeToString(salt),
		EncryptedPass: passwordEncyrpt,
	}
	db.Create(&user)
	userInfo := UserInfo{
		UserID: user.ID,
		NiceName: name,
		Contact:  email,
	}
	db.Create(&userInfo)
	fmt.Println(db.Error)
	err = nil
	return
}

//GetUserByID will get a user struct by uid
func GetUserByID() (user User, err error) {
	user = User{}
	err = nil
	return
}
