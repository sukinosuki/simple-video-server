package auth

import (
	"simple-video-server/constants/gender"
	"time"
)

// RegisterForm AuthRegisterForm RegisterDTO 注册form
type RegisterForm struct {
	Email    string `json:"email" form:"email" binding:"required,max=50,min=6,email" label:"邮箱"`
	Password string `json:"password" form:"password" binding:"required,max=12,min=6" label:"密码"`
	Code     string `json:"code" form:"code" binding:"required,len=5"`

	//	TODO: 多种注册方式
}

// LoginForm 登录form
type LoginForm struct {
	Email    string `json:"email" form:"email" binding:"required,email" label:"邮箱"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=12" label:"密码"`
}

// LoginRes AuthLoginRes 登录响应
type LoginRes struct {
	Profile *LoginResProfile `json:"profile"`
	Token   string           `json:"token"`
}

// LoginResProfile AuthLoginResProfile 个人综合信息
type LoginResProfile struct {
	User            LoginResProfileUser `json:"user" gorm:"embedded;embeddedPrefix:user_"`
	LikeCount       int64               `json:"like_count"`       // 所有视频点赞数
	DislikeCount    int64               `json:"dislike_count"`    // 所有视频点踩数
	CollectionCount int64               `json:"collection_count"` // 收藏数
	VideoCount      int64               `json:"video_count"`      // 发布的视频数
	FollowerCount   int64               `json:"follower_count"`   // 粉丝数
	FollowingCount  int64               `json:"following_count"`  // 关注(其他用户)数
}

// LoginResProfileUser AuthLoginResProfileUser 个人信息
type LoginResProfileUser struct {
	ID        uint          `json:"id"`
	CreatedAt time.Time     `json:"created_at"`
	Enabled   bool          `json:"enabled"`
	Nickname  string        `json:"nickname"`
	Email     string        `json:"email"`
	Avatar    string        `json:"avatar"`
	Birthday  *time.Time    `json:"birthday"`
	Gender    gender.Gender `json:"gender"`
}

// UpdateForm AuthUpdateDTO 更新
type UpdateForm struct {
	Nickname string     `json:"nickname" form:"nickname" binding:"required"`
	Birthday *time.Time `json:"birthday" form:"birthday"`
	//Gender   *gender.Gender `json:"gender" form:"gender" binding:"required"`
	Gender *gender.Gender `json:"gender" form:"gender" binding:"required"`
	Avatar string         `json:"avatar" form:"avatar" binding:"required"`
}

// ResetPasswordForm AuthResetPasswordForm 重置密码
type ResetPasswordForm struct {
	//Email string `json:"email" form:"email" binding:"email,max=50"` // 不需要提交email字段, 由uid从数据库获取到email
	Password string `json:"password" form:"password" binding:"required,min=6,max=12"`
	Code     string `json:"code" form:"code" binding:"required,len=5"`
	Method   string `json:"method" form:"method" binding:"required,oneof=email code"`
}
