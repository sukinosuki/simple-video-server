package app_server

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/app_server/modules/follow"
	"simple-video-server/app_server/modules/test_student"
	"simple-video-server/app_server/modules/upload"
	"simple-video-server/app_server/modules/user/collection"
	"simple-video-server/app_server/modules/user/like"
	"simple-video-server/app_server/modules/user/user"
	"simple-video-server/app_server/modules/video"
)

func SetupRoutes(r *gin.RouterGroup) {

	v1 := r.Group("/api/v1",
		middleware.TraceHandler,
		middleware.PreAuthorizeHandler,
		middleware.ErrorHandler,
		middleware.RequestLogHandler,
	)

	video.SetupRoutes(v1)

	user.SetupRoutes(v1)

	test_student.SetupRoutes(v1)

	collection.SetupRoutes(v1)

	like.SetupRoutes(v1)

	upload.SetupRoutes(v1)

	follow.SetupRoutes(v1)

}
