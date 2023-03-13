package user

import (
	"time"
)

// UserRegister 注册form
type UserRegister struct {
	Email    string `json:"email" form:"email" binding:"required,max=50,min=6,email" label:"邮箱"`
	Password string `json:"password" form:"password" binding:"required,max=12,min=6" label:"密码"`
	Code     string `json:"code" form:"code" binding:"required,len=5"`
}

// UserLogin 登录form
type UserLogin struct {
	Email    string `json:"email" form:"email" binding:"required,email" label:"邮箱"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=12" label:"密码"`
}

// Profile 个人综合信息
type Profile struct {
	User            ProfileUser `json:"user" gorm:"embedded;embeddedPrefix:user_"`
	LikeCount       int64       `json:"like_count"`       // 所有视频点赞数
	DislikeCount    int64       `json:"dislike_count"`    // 所有视频点踩数
	CollectionCount int64       `json:"collection_count"` // 收藏数
	VideoCount      int64       `json:"video_count"`      // 发布的视频数
	FlowerCount     int64       `json:"fans_count"`       // 粉丝数
}

// ProfileUser 个人信息
type ProfileUser struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Enabled   bool      `json:"enabled"`
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
}

// LoginRes 登录响应
type LoginRes struct {
	Profile *Profile `json:"profile"`
	Token   string   `json:"token"`
}

// UserResetPassword 重置密码form
type UserResetPassword struct {
	//Email string `json:"email" form:"email" binding:"email,max=50"` // 不需要提交email字段, 由uid从数据库获取到email
	Password string `json:"password" form:"password" binding:"required,min=6,max=12"`
	Code     string `json:"code" form:"code" binding:"required,len=5"`
	*ResetPasswordMethod
}

type ResetPasswordMethod struct {
	Method string `json:"method" form:"method" binding:"required,oneof=email code"`
}
