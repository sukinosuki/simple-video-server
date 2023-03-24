package like

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	// 需要登录
	shouldAuth := v1.Group("", middleware.AuthorizeHandler)

	{
		// 新增收藏
		shouldAuth.POST("/auth/like/video", core.ToHandler(Api.Add, "collection"))

		// 删除收藏
		shouldAuth.DELETE("/auth/like/video", core.ToHandler(Api.Delete, "collection"))

		// 获取全部收藏
		//shouldAuth.GET("/user/like/video", common.ToHandler(Api.GetAll, "collection"))
	}
}
