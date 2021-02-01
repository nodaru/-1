package model

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Post 指代 帖子相关的内容
type Post struct {
	ID int `gorm:"primaryKey;autoIncrement;not null"`
	// PostHistory   PostHistory
	UserInfoID int
	TUrl       string
	NumLike    int
	NumDislike int
	NumReport  int
	Status     state
	Comment    []Comment
	CreatedAt  time.Time      `gorm:"index"`
	UpdatedAt  time.Time      `gorm:"index"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

// // PostHistory 历史记录表 记录修改信息
// type PostHistory struct {
// 	PostID    int `gorm:"primaryKey;index"`
// 	OUrl      string
// 	CreatedAt time.Time      `gorm:"index"`
// 	UpdatedAt time.Time      `gorm:"index"`
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }

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

//GetAllComment will get all comment belong to a post
func (post *Post) GetAllComment()(c *[]Comment,err error){
	err  = db.Preload(clause.Associations).Find(post).Error
	c = &post.Comment
	return
	}