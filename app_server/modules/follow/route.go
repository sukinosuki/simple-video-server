package follow

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/core"
)

var _moduleName = "follow"

func SetupRoutes(v1 *gin.RouterGroup) {
	// 关注度排名
	v1.GET("/follow/rank", core.ToHandler(Api.FollowScores, _moduleName))

	// 获取某个用户的粉丝列表
	v1.GET("/user/:id/follower", core.ToHandler(Api.GetUserFollowers, _moduleName))

	// 获取某个用户的关注列表
	v1.GET("/user/:id/following", core.ToHandler(Api.GetUserFollowing, _moduleName))

	// 需要登录
	auth := v1.Group("", middleware.AuthorizeHandler)

	{
		// 登录用户新增关注
		auth.POST("/user/follow", core.ToHandler(Api.Follow, _moduleName))

		// 登录用户删除关注
		auth.DELETE("/user/follow", core.ToHandler(Api.Unfollow, _moduleName))
	}
}
