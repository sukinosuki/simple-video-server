package collection

import (
	"simple-video-server/common"
	"time"
)

type AddCollection struct {
	VID uint `json:"vid" form:"vid" binding:"required"`
}

type UserVideoCollectionRes struct {
	ID        uint                        `json:"id" gorm:"column:id;"`
	Title     string                      `json:"title"`
	Cover     string                      `json:"cover"`
	CreatedAt time.Time                   `json:"created_at"`
	User      *UserVideoCollectionResUser `json:"user" gorm:"embedded;embeddedPrefix:user_"`
	// TODO: 返回失效状态用来展示失效视频
}

type UserVideoCollectionResUser struct {
	ID       uint   `json:"id"`
	Nickname string `json:"nickname"`
}

type CollectionQuery struct {
	common.Pager
}
