package user

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	//TODO:多种类型媒体收藏(video、post...)
	//TODO:重复的toHandler方法

	//注册
	v1.POST("/user/register", core.ToHandler(Api.RegisterHandler, "user"))

	// 登录
	v1.POST("/user/login", core.ToHandler(Api.Login, "user"))

	// 需要登录
	auth := v1.Group("", middleware.AuthorizeHandler)

	{
		//获取用户信息
		auth.GET("/user/profile", core.ToHandler(Api.Profile, "user"))

		auth.POST("/user/reset-password", core.ToHandler(Api.ResetPassword, "user"))

		auth.POST("/user/logoff", core.ToHandler(Api.Logoff, "user"))
	}
}
