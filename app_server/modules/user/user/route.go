package user

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	//v1.POST("/user/register", Api.Register)
	// 优化toHandler
	//注册
	v1.POST("/user/register", core.ToHandler(Api.RegisterHandler, "user"))

	// 登录
	v1.POST("/user/login", core.ToHandler(Api.Login, "user"))

	//v1.GET("/user/profile", middleware2.PreAuthorizeHandler, middleware2.AuthorizeHandler, common.ToHandler(Api.Profile))
	// 需要登录
	shouldAuth := v1.Group("", middleware.AuthorizeHandler)

	{
		//获取用户信息
		shouldAuth.GET("/user/profile", core.ToHandler(Api.Profile, "user"))
		//
		////TODO:多种类型媒体收藏(video、post...)
		////TODO:重复的toHandler方法
		//
		//// 新增收藏
		//shouldAuth.POST("/user/collection/video", common.ToHandler(Api.Add, "user"))
		//
		//// 删除收藏
		//shouldAuth.DELETE("/user/collection/video/:id", common.ToHandler(Api.Delete, "user"))
		//
		//// 获取全部收藏
		//shouldAuth.GET("/user/collection/video", common.ToHandler(Api.GetAll, "user"))
	}
}
