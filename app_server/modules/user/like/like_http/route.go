package like_http

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/app_server/modules/user/like/internal"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	api := internal.GetApi()

	// 需要登录
	shouldAuth := v1.Group("", middleware.AuthorizeHandler)

	{
		// 新增收藏
		shouldAuth.POST("/auth/like/video", core.ToHandler(api.Add, "user_like"))

		// 删除收藏
		shouldAuth.DELETE("/auth/like/video", core.ToHandler(api.Delete, "user_like"))

		// 获取全部收藏
		//shouldAuth.GET("/user/like/video", common.ToHandler(_Api.GetAll, "collection"))
	}
}
