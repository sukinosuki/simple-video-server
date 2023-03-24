package video

import (
	"simple-video-server/common"
	"time"
)

type VideoAdd struct {
	Title string `json:"title" form:"title" binding:"required,max=50"`
	Cover string `json:"cover" form:"cover" binding:"required,max=255"`
	Url   string `json:"url" form:"url" binding:"required,max=255"`
}

type VideoUpdate struct {
	Title string `json:"title" form:"title" binding:"required,max=50"`
	Cover string `json:"cover" form:"cover" binding:"required,max=50"`
}

type VideoQuery struct {
	common.Pager
	//UID     *uint  `json:"uid" form:"uid"`
	Exclude []uint `json:"exclude" form:"exclude"`
	random  bool
}

type VideoSimple struct {
	ID        uint             `json:"id"`
	CreatedAt time.Time        `json:"created_at"`
	Cover     string           `json:"cover"`
	Url       string           `json:"url"`
	Title     string           `json:"title"`
	Locked    bool             `json:"locked"`
	User      *VideoSimpleUser `json:"user" gorm:"embedded;embeddedPrefix:user_"`
}

type VideoSimpleUser struct {
	ID       uint   `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type VideoResVideo struct {
	ID        uint              `json:"id"`
	CreatedAt time.Time         `json:"created_at"`
	Cover     string            `json:"cover"`
	Title     string            `json:"title"`
	Url       string            `json:"url"`
	Uid       uint              `json:"uid"`
	User      VideoResVideoUser `json:"user" gorm:"embedded;embeddedPrefix:user_"`
}

type VideoResVideoUser struct {
	ID       uint   `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type VideoRes struct {
	Video           *VideoResVideo `json:"video"`
	LikeCount       int            `json:"like_count"`       //点赞数
	DislikeCount    int            `json:"dislike_count"`    //点踩数
	IsLike          bool           `json:"is_like"`          //是否点赞
	IsDislike       bool           `json:"is_dislike"`       // 是否点踩
	IsCollect       bool           `json:"is_collect"`       //是否已收藏
	CollectionCount int64          `json:"collection_count"` // 收藏数
	CommentCount    int64          `json:"comment_count"`    //评论数
}

//func init() {
//	err := global.MysqlDB.AutoMigrate(
//		&Video{},
//	)
//
//	if err != nil {
//		panic(err)
//	}
//}
