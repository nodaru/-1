package model

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Comment 指代 帖子相关的内容
type Comment struct {
	ID         int `gorm:"primaryKey;autoIncrement;not null"`
	Curl       string
	PostID     int
	UserInfoID int
	NumLike    int
	NumDislike int
	NumReport  int
	Status     state
	//自引用的Comment
	RefCommentID *int
	RefComment   *Comment
	// CommentHistory   CommentHistory   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time      `gorm:"index"`
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// // CommentHistory 历史记录表 记录修改信息
// type CommentHistory struct {
// 	CommentID int `gorm:"primaryKey;index"`
// 	OUrl      string
// 	CreatedAt time.Time      `gorm:"index"`
// 	UpdatedAt time.Time      `gorm:"index"`
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }


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