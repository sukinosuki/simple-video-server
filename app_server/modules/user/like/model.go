package like

type VideoLike struct {
	//VID      uint `json:"vid" form:"vid" binding:"required"`
	LikeType int `json:"like_type" form:"like_type" binding:"required"` // 0点踩1点赞
}
