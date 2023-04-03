package http

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/app_server/modules/upload/internal"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {
	api := internal.GetApi()

	// 需要登录
	shouldAuth := v1.Group("", middleware.AuthorizeHandler)

	shouldAuth.POST("/upload", core.ToHandler(api.Upload, "upload"))
}
