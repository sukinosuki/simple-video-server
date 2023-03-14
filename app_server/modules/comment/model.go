package comment

import (
	"simple-video-server/common"
	"time"
)

type CommentAdd struct {
	Content string `json:"content" form:"content" binding:"required,max=250"`
	//MediaID   uint   `json:"media_id" form:"media_id" binding:"required"`
	//MediaType int    `json:"media_type" form:"media_type" binding:"required"`
	AtUID   *uint `json:"at_uid" form:"at_uid"`
	Root    *uint `json:"root" form:"root"`
	ReplyID *uint `json:"reply_id" form:"reply_id"`
}

type CommentQuery struct {
	//Page int `json:"page" form:"page"`
	//Size int `json:"size" form:"size"`
	common.Pager
}

type CommentResSimple struct {
	ID         uint                `json:"id"`
	Root       *uint               `json:"root"`
	AtUID      *uint               `json:"at_uid"`
	Content    string              `json:"content"`
	MediaID    uint                `json:"media_id"`
	MediaType  int                 `json:"media_type"`
	UID        uint                `json:"uid"`
	CreatedAt  time.Time           `json:"created_at"`
	Like       int                 `json:"like"`
	Dislike    int                 `json:"dislike"`
	ReplyCount int                 `json:"reply_count"`
	RowNum     int                 `json:"row_num"`
	Replies    []*CommentResSimple `json:"replies,omitempty" gorm:"-"`

	User *CommentResSimpleUser `json:"user" gorm:"embedded;embeddedPrefix:user_"`
}

type CommentResSimpleUser struct {
	ID       uint   `json:"id"`
	Nickname string `json:"nickname"`
	Cover    string `json:"cover"`
}
