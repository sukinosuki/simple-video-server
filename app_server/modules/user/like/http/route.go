package http

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	//api := GetApi()

	// 获取全部收藏
	//v1.GET("/user/like/video", common.ToHandler(_Api.GetAll, "collection"))

	// 需要登录
	shouldAuth := v1.Group("", middleware.AuthorizeHandler)

	{
		// 新增收藏
		shouldAuth.POST("/auth/like/video/:id", core.ToHandler(_api.Add, "user_like"))

		// 删除收藏
		shouldAuth.DELETE("/auth/like/video/:id", core.ToHandler(_api.Delete, "user_like"))

	}
}
