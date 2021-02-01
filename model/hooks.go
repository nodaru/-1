package model

import (
	"gorm.io/gorm"
)

// TODO: 增加举报、喜欢、不喜欢后的同步更新

//AfterCreate -> User, Set the first registered user to be superuser
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	if u.ID == 1 {
		err = tx.Model(u).Update("Privilege", adminPrivilege).Error
	}
	return
}

//BeforeCreate -> User, Check wether there have a registed email
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	result :=db.Where("Contact=?",u.UserInfo.Contact).Find(&u.UserInfo)
	if u.UserInfo.ID == 0 ||result.RowsAffected == 0 {
		err = nil
	}else{
		err = ERR_EMAIL_HAVE_BEEN_REGISTED
	}
	return
}

//AfterCreate -> Comment, init the first comment
func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	if c.RefComment != nil &&c.RefComment.ID == 0{
	err = tx.Model(c).Update("RefCommentID", 1).Error
	}
	return
}

//BeforeCreate -> Post, init the first Post
func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

//AfterCreate -> post, handle the owner of the post
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	return
}