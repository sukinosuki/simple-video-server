package like

import "simple-video-server/constants/like_type"

type VideoLike struct {
	//VID      uint `json:"vid" form:"vid" binding:"required"`
	LikeType like_type.LikeType `json:"like_type" form:"like_type" binding:"required"` // 0点踩1点赞
}
