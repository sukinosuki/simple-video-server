package internal

import (
	"simple-video-server/app_server/modules/follow"
	"simple-video-server/core"
)

type followApi struct {
	service *_service
}

var FollowApi = &followApi{
	service: Service,
}

func (api *followApi) Follow(c *core.Context) (bool, error) {
	//var followAdd *FollowAdd
	//err := c.ShouldBind(&followAdd)
	//if err != nil {
	//	panic(err)
	//}

	//targetUID, err := strconv.Atoi(c.Param("target_uid"))
	//if err != nil {
	//	panic(err)
	//}

	err := api.service.Follow(c, c.GetParamId())
	if err != nil {
		panic(err)
	}

	return true, err
}

func (api *followApi) Unfollow(c *core.Context) (bool, error) {
	//var followAdd *FollowAdd
	//err := c.ShouldBind(&followAdd)
	//if err != nil {
	//	panic(err)
	//}
	//targetUID, err := strconv.Atoi(c.Param("target_uid"))
	//if err != nil {
	//	panic(err)
	//}
	err := api.service.Unfollow(c, c.GetParamId())
	if err != nil {
		panic(err)
	}

	return true, err
}

// FollowScores 获取follow排名
func (api *followApi) FollowScores(c *core.Context) ([]*follow.UserFollowerRankResSimple, error) {
	var query follow.UserFollowRankQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}

	users, err := api.service.FollowScores(c, &query)
	if err != nil {
		panic(err)
	}

	if users == nil {
		users = make([]*follow.UserFollowerRankResSimple, 0)
	}

	return users, err
}

func (api *followApi) GetUserFollowers(c *core.Context) ([]*follow.UserFollowerResSimple, error) {

	var query follow.UserFollowerQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}

	targetUid := c.GetParamUID()
	followers, err := api.service.GetUserFollower(c, targetUid, &query)

	if err != nil {
		panic(err)
	}

	return followers, err
}

func (api *followApi) GetUserFollowing(c *core.Context) ([]*follow.UserFollowerResSimple, error) {
	var query follow.UserFollowerQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}

	targetUid := c.GetParamUID()

	following, err := api.service.GetUserFollowing(c, targetUid, &query)

	if err != nil {
		panic(err)
	}

	return following, err
}
