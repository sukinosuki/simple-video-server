package collection_http

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/app_server/modules/user/collection/internal"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {
	api := internal.GetApi()

	// 获取收藏列表
	v1.GET("/user/:uid/collection/video", core.ToHandler(api.GetAll, "user_collection"))

	// 需要登录
	auth := v1.Group("", middleware.AuthorizeHandler)

	{
		// 新增收藏
		auth.POST("/auth/collection/video", core.ToHandler(api.Add, "user_collection"))

		// 删除收藏
		auth.DELETE("/auth/collection/video/:id", core.ToHandler(api.Delete, "user_collection"))

	}
}
