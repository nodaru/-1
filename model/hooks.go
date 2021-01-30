package model

import (
	"errors"
	"gorm.io/gorm"
)

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
		err = errors.New("email have been regisisted")
	}
	return
}