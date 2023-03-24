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
