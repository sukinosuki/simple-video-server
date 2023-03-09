package collection

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	// 需要登录
	auth := v1.Group("", middleware.AuthorizeHandler)

	{
		// 新增收藏
		auth.POST("/user/collection/video", core.ToHandler(Api.Add, "collection"))

		// 删除收藏
		auth.DELETE("/user/collection/video/:id", core.ToHandler(Api.Delete, "collection"))

		// 获取收藏列表
		auth.GET("/user/collection/video", core.ToHandler(Api.GetAll, "collection"))
	}
}
