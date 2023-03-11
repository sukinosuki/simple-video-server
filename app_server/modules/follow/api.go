package follow

import "simple-video-server/core"

type _api struct {
	service *_service
}

var Api = &_api{
	service: Service,
}

func (api *_api) Follow(c *core.Context) (bool, error) {
	var followAdd *FollowAdd
	err := c.ShouldBind(&followAdd)
	if err != nil {
		panic(err)
	}

	err = api.service.Follow(c, followAdd.UID)
	if err != nil {
		panic(err)
	}

	return true, err
}

func (api *_api) Unfollow(c *core.Context) (bool, error) {
	var followAdd *FollowAdd
	err := c.ShouldBind(&followAdd)
	if err != nil {
		panic(err)
	}

	err = api.service.Unfollow(c, followAdd.UID)
	if err != nil {
		panic(err)
	}

	return true, err
}

// FollowScores 获取follow排名
func (api *_api) FollowScores(c *core.Context) ([]*UserFollowerRankResSimple, error) {
	var query UserFollowRankQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}

	users, err := api.service.FollowScores(c, &query)
	if err != nil {
		panic(err)
	}

	if users == nil {
		users = make([]*UserFollowerRankResSimple, 0)
	}

	return users, err
}

func (api *_api) GetUserFollowers(c *core.Context) ([]*UserFollowerResSimple, error) {

	var query UserFollowerQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}

	targetUid := c.GetId()
	followers, err := api.service.GetUserFollower(c, targetUid, &query)

	if err != nil {
		panic(err)
	}

	return followers, err
}

func (api *_api) GetUserFollowing(c *core.Context) ([]*UserFollowerResSimple, error) {
	var query UserFollowerQuery
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}

	targetUid := c.GetId()

	following, err := api.service.GetUserFollowing(c, targetUid, &query)

	if err != nil {
		panic(err)
	}

	return following, err
}
