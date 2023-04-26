package http

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	// TODO:多种类型媒体收藏(video、post...)
	// TODO:重复的toHandler方法

	// 注册
	v1.POST("/auth/register", core.ToHandler(_api.Register, "user"))

	// 登录
	v1.POST("/auth/login", core.ToHandler(_api.Login, "user"))

	// 需要登录
	auth := v1.Group("", middleware.AuthorizeHandler)

	{
		// 更新用户信息
		auth.PUT("/auth/profile", core.ToHandler(_api.UpdateProfile, "user"))

		// 获取用户信息
		auth.GET("/auth/profile", core.ToHandler(_api.AuthProfile, "user"))

		// 重置密码
		auth.POST("/auth/reset-password", core.ToHandler(_api.ResetPassword, "user"))
		// TODO: 路由设计
		auth.POST("/auth/password/reset", core.ToHandler(_api.ResetPassword, "user"))

		// 注销
		auth.POST("/auth/logoff", core.ToHandler(_api.Logoff, "user"))
	}
}
