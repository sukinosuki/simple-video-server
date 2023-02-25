package app_server

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/app_server/modules/user"
	"simple-video-server/app_server/modules/video"
)

func SetupRoutes(r *gin.RouterGroup) {

	v1 := r.Group("/api/v1", middleware.TraceHandler, middleware.ErrorHandler, middleware.RequestLogHandler)

	video.SetupRoutes(v1)

	user.SetupRoutes(v1)
}
