package http

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/app_server/modules/follow/internal"
	"simple-video-server/core"
)

var _moduleName = "follow"

func SetupRoutes(v1 *gin.RouterGroup) {
	// 关注度排名
	v1.GET("/follow/rank", core.ToHandler(internal.FollowApi.FollowScores, _moduleName))

	// 获取某个用户的粉丝列表
	v1.GET("/user/:uid/follower", core.ToHandler(internal.FollowApi.GetUserFollowers, _moduleName))

	// 获取某个用户的关注列表
	v1.GET("/user/:uid/following", core.ToHandler(internal.FollowApi.GetUserFollowing, _moduleName))

	// 需要登录
	auth := v1.Group("", middleware.AuthorizeHandler)

	{
		// auth用户新增关注
		auth.POST("/auth/following/:id", core.ToHandler(internal.FollowApi.Follow, _moduleName))

		// auth用户删除关注
		auth.DELETE("/auth/following/:id", core.ToHandler(internal.FollowApi.Unfollow, _moduleName))
	}
}
