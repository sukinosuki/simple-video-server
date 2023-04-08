package http

import (
	"simple-video-server/app_server/modules/follow"
	"simple-video-server/app_server/modules/user/user"
	"simple-video-server/core"
)

type Service struct {
	dao         *user.Dao
	followCache *follow.Cache
}

var service = &Service{
	dao:         user.GetDao(),
	followCache: follow.GetFollowCache(),
}

func (s *Service) GetAll(c *core.Context, query *user.UserQuery) ([]*user.UserSimple, error) {

	users, err := s.dao.GetAll(query)

	for _, user := range users {
		// 自己是否关注了该用户
		if c.Authorized {
			isFollowing, err := s.followCache.IsFollowingOneUser(*c.AuthUID, user.ID)
			if err != nil {
				panic(err)
			}

			user.IsFollowing = isFollowing
		}

		//用户的粉丝数
		followersCount, err := s.followCache.OneUserFollowersCount(user.ID)
		if err != nil {
			panic(err)
		}

		user.FollowersCount = followersCount

		followingCount, err := s.followCache.OneUserFollowingCount(user.ID)
		if err != nil {
			panic(err)
		}

		user.FollowingCount = followingCount
	}

	return users, err
}
