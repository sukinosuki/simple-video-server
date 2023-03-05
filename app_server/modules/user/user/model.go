package user

import (
	"simple-video-server/models"
)

type UserRegister struct {
	Email    string `json:"email" form:"email" binding:"required,max=50,min=6,email" label:"邮箱"`
	Password string `json:"password" form:"password" binding:"required,max=12,min=6" label:"密码"`
}

type UserLogin struct {
	Email    string `json:"email" form:"email" binding:"required,email" label:"邮箱"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=12" label:"密码"`
}

type UserProfile struct {
	LikeCount       int //点赞数
	CollectionCount int // 收藏数
}

type LoginRes struct {
	User  *models.User `json:"user"`
	Token string       `json:"token"`
}
