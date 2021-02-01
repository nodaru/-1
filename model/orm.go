package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
//InitTable will init the table
func InitTable(){
	usr,err := CreateNormalUser("rux","lux@luxru.top","000114")
	switch err{
	case ERR_EMAIL_HAVE_BEEN_REGISTED:
		err = nil
		break
	case nil:
		break
	default:
		panic(err)
	}
	usr.UserInfo.Comment = []Comment{{
		ID: 1,
		Curl: "/null",
		UserInfoID: usr.UserInfo.ID,
		Status: state_block,
	},
	}
	usr.UserInfo.Post = []Post{{
		ID: 1,
		UserInfoID: usr.UserInfo.ID,
		TUrl: "/null",
		Status: state_block,
		Comment: usr.UserInfo.Comment,
	},
	}
	err = db.Save(usr).Error
	return
}

//MigrateDB will migrate DB
func MigrateDB() {
	// 迁移 schema
	db.AutoMigrate(&User{}, &UserInfo{}, &Post{}, &Agreement{},
		&Comment{}, &Report{}, &Edit{}, &Ban{}, &SysNotice{})
}



