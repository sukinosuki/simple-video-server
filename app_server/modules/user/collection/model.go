package collection

import (
	"simple-video-server/common"
	"time"
)

type AddCollection struct {
	VID uint `json:"vid" form:"vid" binding:"required"`
}

type UserVideoCollectionRes struct {
	ID        uint      `json:"id" gorm:"column:id;"`
	Title     string    `json:"title"`
	Cover     string    `json:"cover"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	// TODO: 返回失效状态用来展示失效视频
}

type CollectionQuery struct {
	common.Pager
}
