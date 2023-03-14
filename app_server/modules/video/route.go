package video

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	// 可能登录了
	//possibleAuth := v1.Group("", middleware.PreAuthorizeHandler)
	//v1.GET("/video/:id", middleware.PreAuthorizeHandler, common.ToHandler(Api.GetById))
	v1.GET("/video/:id", core.ToHandler(Api.GetById, "video"))

	v1.GET("/video", core.ToHandler(Api.GetAll, "video"))

	//v1.GET("/video/:id/comment", core.ToHandler(Api.GetComment, "video"))

	// 需要登录
	shouldAuth := v1.Group("", middleware.AuthorizeHandler)
	// 优化toHandler
	{
		shouldAuth.POST("/video", core.ToHandler(Api.Add, "video"))

		shouldAuth.PUT("/video/:id", core.ToHandler(Api.Update, "video"))

		shouldAuth.DELETE("/video/:id", core.ToHandler(Api.Delete, "video"))
	}
}
