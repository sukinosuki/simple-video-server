package models

import "time"

type Comment struct {
	ID uint `json:"id"`
	// 一级评论的id
	Root *uint `json:"root"`
	// @的用户id
	AtUID *uint `json:"at_uid" gorm:"column:at_uid"`
	// 评论内容(TODO: 图文评论
	Content string `json:"content" gorm:"size:255;not null;"`
	// 可以是video、post的id
	MediaID uint `json:"media" gorm:"index;not null"`
	// 根据media_type判断是评论的video、文章等
	// media_type + media_id 为唯一
	MediaType int `json:"type" gorm:"not null"`
	// 评论人id
	UID uint `json:"uid" gorm:"index;not null"`
	// 回复的评论id(对哪个一级评论进行回复
	ReplyID *uint `json:"reply_id"`
	Like    int   `json:"like" gorm:"not null"`
	Dislike int   `json:"dislike" gorm:"not null"`

	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}
