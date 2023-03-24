package auth

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	// TODO:多种类型媒体收藏(video、post...)
	// TODO:重复的toHandler方法

	// 获取用户信息
	//v1.GET("/user/:id/profile", core.ToHandler(Api.Profile, "user"))

	// 注册
	v1.POST("/auth/register", core.ToHandler(Api.RegisterHandler, "user"))

	// 登录
	v1.POST("/auth/login", core.ToHandler(Api.Login, "user"))

	// 需要登录
	auth := v1.Group("", middleware.AuthorizeHandler)

	// TODO: 可以获取自己和其他人的信息
	{
		//更新用户信息
		auth.PUT("/auth/profile", core.ToHandler(Api.UpdateProfile, "user"))

		//获取用户信息
		auth.GET("/auth/profile", core.ToHandler(Api.AuthProfile, "user"))

		//重置密码
		auth.POST("/auth/reset-password", core.ToHandler(Api.ResetPassword, "user"))

		//注销
		auth.POST("/auth/logoff", core.ToHandler(Api.Logoff, "user"))

		//// 登录用户新增关注
		//auth.POST("/auth/follower/:id", core.ToHandler(Api.Follow, "user"))
		//
		//// 登录用户删除关注
		//auth.DELETE("/auth/follower/:id", core.ToHandler(Api.Unfollow, "user"))
	}
}
