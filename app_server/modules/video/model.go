package video

import (
	"simple-video-server/models"
)

type VideoAdd struct {
	Title string `json:"title" form:"title" binding:"required,max=50"`
	Cover string `json:"cover" form:"cover" binding:"required,max=255"`
	Url   string `json:"url" form:"url" binding:"required,max=255"`
}

type VideoUpdate struct {
	Title string `json:"title" form:"title" binding:"required,max=50"`
	Cover string `json:"cover" form:"cover" binding:"required,max=50"`

	//VID       string    `db:"vid"`
	//CreatedAt time.Time `db:"created_at"`
}

type VideoRes struct {
	Video *models.Video `json:"video"`
	//点赞数
	LikeCount int `json:"like_count"`
	//点踩数
	DislikeCount int `json:"dislike_count"`
	//是否点赞
	IsLike bool `json:"is_like"`
	// 是否点踩
	IsDislike bool `json:"is_dislike"`
	//是否已收藏
	IsCollect bool `json:"is_collect"`
	// 收藏数
	CollectionCount int64 `json:"collection_count"`
	//评论数
	CommentCount int64 `json:"comment_count"`
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
