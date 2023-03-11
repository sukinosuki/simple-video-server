package models

import "time"

type Comment struct {
	ID uint `json:"id"`
	// 一级评论的id
	PID uint `json:"pid" gorm:"not null"`
	// @的用户id
	AtUID uint `json:"at_uid"`
	// 评论内容(TODO: 图文评论
	Content   string    `json:"content" gorm:"size:255;not null;"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	// 可以是video、post的id
	MediaID uint `json:"media" gorm:"index"`
	// 根据media_type判断是评论的video、文章等
	// media_type + media_id 为唯一
	MediaType int `json:"type"`
	// 评论人id
	UID uint `json:"uid" gorm:"index"`
	// 回复的评论id
	ReplyID uint `json:"reply_id"`
}
