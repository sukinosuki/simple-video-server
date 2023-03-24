package follow

type FollowAdd struct {
	UID uint `json:"uid" form:"uid"`
}

type UserFollowRankQuery struct {
	Range []int64 `json:"range" form:"range"` // TODO binding 限制range里的取值范围
}

type UserFollowerRankResSimple struct {
	ID       uint    `json:"id"`
	Nickname string  `json:"nickname"`
	Score    float64 `json:"score" gorm:"-"`
	Avatar   string  `json:"avatar"`
}

type UserFollowerResSimple struct {
	ID       uint   `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type UserFollowerQuery struct {
	IsInter bool `json:"is_inter" form:"is_inter"` // 是否获取共同关注/粉丝
}
