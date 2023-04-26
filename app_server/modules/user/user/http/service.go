package http

import (
	"simple-video-server/app_server/modules/follow"
	"simple-video-server/app_server/modules/user/user"
	"simple-video-server/core"
	"simple-video-server/pkg/arr"
)

type Service struct {
	dao         *user.Dao
	followCache *follow.Cache
}

var service = &Service{
	dao:         user.GetDao(),
	followCache: follow.GetCache(),
}

func (s *Service) GetRanks(c *core.Context, query *user.UserQuery) ([]*user.RankUsers, error) {
	handlerName := "GetRanks"
	users, err := s.dao.GetRanks(query)

	c.PanicIfErr(err, handlerName, "获取用户失败")

	arr.ForEach(users, func(item *user.RankUsers, index int) {
		isFollowing := false
		// 自己是否关注了该用户
		if c.Authorized {
			isFollowing, err = s.followCache.IsUserFollowingAnotherUser(*c.AuthUID, item.User.ID)
			if err != nil {
				panic(err)
			}
		}

		//用户的粉丝数
		followersCount, err := s.followCache.GetOneUserFollowersCount(item.User.ID)
		if err != nil {
			panic(err)
		}

		//user.FollowersCount = followersCount

		followingCount, err := s.followCache.GetOneUserFollowingCount(item.User.ID)
		if err != nil {
			panic(err)
		}

		item.IsFollowing = isFollowing
		item.FollowersCount = followersCount
		item.FollowingCount = followingCount
	})

	return users, nil
	//_users := arr.Map(users, func(item *user.RankUsers, index int) *user.RankUsers {
	//	isFollowing := false
	//	// 自己是否关注了该用户
	//	if c.Authorized {
	//		isFollowing, err = s.followCache.IsUserFollowingAnotherUser(*c.AuthUID, item.User.ID)
	//		if err != nil {
	//			panic(err)
	//		}
	//
	//		//user.IsFollowing = isFollowing
	//
	//	}
	//
	//	//用户的粉丝数
	//	followersCount, err := s.followCache.GetOneUserFollowersCount(item.User.ID)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	//user.FollowersCount = followersCount
	//
	//	followingCount, err := s.followCache.GetOneUserFollowingCount(item.User.ID)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	//user.FollowingCount = followingCount
	//	return &user.RankUsers{
	//		User: user.RankUser{
	//			ID:       item.User.ID,
	//			Nickname: item.User.Nickname,
	//			Avatar:   item.User.Avatar,
	//		},
	//		FollowingCount: followingCount,
	//		FollowersCount: followersCount,
	//		IsFollowing:    isFollowing,
	//	}
	//})
	//
	//return _users, nil
	//_users := arr.Map(users, func(item *user.RankUsers, index int) *user.RankUsers {
	//	isFollowing := false
	//	// 自己是否关注了该用户
	//	if c.Authorized {
	//		isFollowing, err = s.followCache.IsUserFollowingAnotherUser(*c.AuthUID, item.User.ID)
	//		if err != nil {
	//			panic(err)
	//		}
	//
	//		//user.IsFollowing = isFollowing
	//
	//	}
	//
	//	//用户的粉丝数
	//	followersCount, err := s.followCache.GetOneUserFollowersCount(item.User.ID)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	//user.FollowersCount = followersCount
	//
	//	followingCount, err := s.followCache.GetOneUserFollowingCount(item.User.ID)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	//user.FollowingCount = followingCount
	//	return &user.RankUsers{
	//		User: user.RankUser{
	//			ID:       item.User.ID,
	//			Nickname: item.User.Nickname,
	//			Avatar:   item.User.Avatar,
	//		},
	//		FollowingCount: followingCount,
	//		FollowersCount: followersCount,
	//		IsFollowing:    isFollowing,
	//	}
	//})
	//
	//return _users, nil
	//for _, user := range users {
	//	// 自己是否关注了该用户
	//	if c.Authorized {
	//		isFollowing, err := s.followCache.IsUserFollowingAnotherUser(*c.AuthUID, user.ID)
	//		if err != nil {
	//			panic(err)
	//		}
	//
	//		user.IsFollowing = isFollowing
	//	}
	//
	//	//用户的粉丝数
	//	followersCount, err := s.followCache.GetOneUserFollowersCount(user.ID)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	user.FollowersCount = followersCount
	//
	//	followingCount, err := s.followCache.GetOneUserFollowingCount(user.ID)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	user.FollowingCount = followingCount
	//}
	//
	//return users, err
}
