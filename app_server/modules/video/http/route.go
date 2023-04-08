package http

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	// 可能登录了
	v1.GET("/video/:id", core.ToHandler(_api.GetById, "video"))

	v1.GET("/video", core.ToHandler(_api.GetAll, "video"))

	v1.GET("/user/:uid/video", core.ToHandler(_api.GetAll, "video"))

	v1.GET("/feed", core.ToHandler(_api.Feed, "video"))

	// 需要登录
	shouldAuth := v1.Group("", middleware.AuthorizeHandler)
	// 优化toHandler
	{
		shouldAuth.GET("/auth/video", core.ToHandler(_api.GetAuthAll, "video"))

		shouldAuth.POST("/auth/video", core.ToHandler(_api.Add, "video"))

		shouldAuth.PUT("/auth/video/:id", core.ToHandler(_api.Update, "video"))

		shouldAuth.DELETE("/auth/video/:id", core.ToHandler(_api.Delete, "video"))
	}
}
