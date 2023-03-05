package upload

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	// 需要登录
	shouldAuth := v1.Group("", middleware.AuthorizeHandler)

	shouldAuth.POST("/upload", core.ToHandler(Api.Upload, "upload"))
}
