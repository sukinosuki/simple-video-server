package http

import (
	"simple-video-server/app_server/modules/follow"
	"simple-video-server/core"
)

type Api struct {
	service *_service
}

var _api = &Api{
	service: Service,
}

func GetApi() *Api {
	return _api
}

func (api *Api) Follow(c *core.Context) (bool, error) {

	err := api.service.Follow(c, c.GetParamId())
	if err != nil {
		panic(err)
	}

	return true, err
}

func (api *Api) Unfollow(c *core.Context) (bool, error) {
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
func (api *Api) FollowScores(c *core.Context) ([]*follow.UserFollowerRankResSimple, error) {
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

func (api *Api) GetUserFollowers(c *core.Context) ([]*follow.UserFollowerResSimple, error) {

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

func (api *Api) GetUserFollowing(c *core.Context) ([]*follow.UserFollowerResSimple, error) {
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
