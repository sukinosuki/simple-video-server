package user

import (
	"simple-video-server/app_server/cache"
	"simple-video-server/core"
)

type Service struct {
	dao *Dao
}

var service = &Service{
	dao: dao,
}

func (s *Service) GetAll(c *core.Context, query *UserQuery) ([]*UserSimple, error) {

	users, err := s.dao.GetAll(query)

	for _, user := range users {
		// 自己是否关注了该用户
		if c.Authorized {
			isFollowing, err := cache.Follow.IsFollowingOneUser(*c.AuthUID, user.ID)
			if err != nil {
				panic(err)
			}

			user.IsFollowing = isFollowing
		}

		//用户的粉丝数
		followersCount, err := cache.Follow.OneUserFollowersCount(user.ID)
		if err != nil {
			panic(err)
		}

		user.FollowersCount = followersCount

		followingCount, err := cache.Follow.OneUserFollowingCount(user.ID)
		if err != nil {
			panic(err)
		}

		user.FollowingCount = followingCount
	}

	return users, err
}
