package http

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/app_server/modules/auth/internal"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	// TODO:多种类型媒体收藏(video、post...)
	// TODO:重复的toHandler方法
	//api := GetApi()
	api := internal.GetApi()
	// 注册
	v1.POST("/auth/register", core.ToHandler(api.RegisterHandler, "user"))

	// 登录
	v1.POST("/auth/login", core.ToHandler(api.Login, "user"))

	// 需要登录
	auth := v1.Group("", middleware.AuthorizeHandler)

	{
		// 更新用户信息
		auth.PUT("/auth/profile", core.ToHandler(api.UpdateProfile, "user"))

		// 获取用户信息
		auth.GET("/auth/profile", core.ToHandler(api.AuthProfile, "user"))

		// 重置密码
		auth.POST("/auth/reset-password", core.ToHandler(api.ResetPassword, "user"))

		// 注销
		auth.POST("/auth/logoff", core.ToHandler(api.Logoff, "user"))

		//// 登录用户新增关注
		//auth.POST("/auth/follower/:id", core.ToHandler(api.Follow, "user"))
		//
		//// 登录用户删除关注
		//auth.DELETE("/auth/follower/:id", core.ToHandler(api.Unfollow, "user"))
	}
}
