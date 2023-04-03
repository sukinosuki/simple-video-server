package app_server

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/app_server/modules/auth/auth_http"
	"simple-video-server/app_server/modules/comment/comment_http"
	"simple-video-server/app_server/modules/email/email_http"
	"simple-video-server/app_server/modules/follow/follow_http"
	"simple-video-server/app_server/modules/test_limit"
	"simple-video-server/app_server/modules/upload/upload_http"
	"simple-video-server/app_server/modules/user/collection/collection_http"
	"simple-video-server/app_server/modules/user/like/like_http"
	"simple-video-server/app_server/modules/user/user/user_http"
	"simple-video-server/app_server/modules/video/video_http"
)

//var V1 *gin.RouterGroup

func SetupRoutes(r *gin.RouterGroup) {

	v1 := r.Group("/api/v1",
		middleware.TraceHandler,
		middleware.PreAuthorizeHandler,
		middleware.ErrorHandler,
		middleware.RequestLogHandler,
	)

	video_http.SetupRoutes(v1)

	auth_http.SetupRoutes(v1)

	user_http.SetupRoutes(v1)

	//test_student.SetupRoutes(v1)

	collection_http.SetupRoutes(v1)

	like_http.SetupRoutes(v1)

	upload_http.SetupRoutes(v1)

	follow_http.SetupRoutes(v1)

	email_http.SetupRoutes(v1)

	comment_http.SetupRoutes(v1)

	test_limit.SetupRoutes(v1)
}
