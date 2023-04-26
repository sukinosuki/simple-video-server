package user

import "simple-video-server/common"

type UserQuery struct {
	common.Pager
}

type UserSimple struct {
	ID             uint   `json:"id"`
	Nickname       string `json:"nickname"`
	Avatar         string `json:"avatar"`
	IsFollowing    bool   `json:"is_follow"`
	FollowingCount int64  `json:"following_count"`
	VideoCount     int64  `json:"video_count"`
	FollowersCount int64  `json:"followers"`
}

type RankUsers struct {
	User RankUser `json:"user" gorm:"embedded;embeddedPrefix:user_"`
	//User            LoginResProfileUser `json:"user" gorm:"embedded;embeddedPrefix:user_"`
	IsFollowing    bool  `json:"is_following" gorm:"-"`
	FollowingCount int64 `json:"following_count" gorm:"-"`
	VideoCount     int64 `json:"video_count"`
	FollowersCount int64 `json:"followers" gorm:"-"`
}

type RankUser struct {
	ID       uint   `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}
