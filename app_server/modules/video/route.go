package video

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	// 可能登录了
	v1.GET("/video/:id", core.ToHandler(Api.GetById, "video"))

	v1.GET("/video", core.ToHandler(Api.GetAll, "video"))

	v1.GET("/user/:uid/video", core.ToHandler(Api.GetAll, "video"))

	v1.GET("/feed", core.ToHandler(Api.Feed, "video"))

	// 需要登录
	shouldAuth := v1.Group("", middleware.AuthorizeHandler)
	// 优化toHandler
	{
		shouldAuth.GET("/auth/video", core.ToHandler(Api.GetAuthAll, "video"))

		shouldAuth.POST("/auth/video", core.ToHandler(Api.Add, "video"))

		shouldAuth.PUT("/auth/video/:id", core.ToHandler(Api.Update, "video"))

		shouldAuth.DELETE("/auth/video/:id", core.ToHandler(Api.Delete, "video"))
	}
}
