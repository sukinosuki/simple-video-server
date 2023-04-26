package app_server

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	authHttp "simple-video-server/app_server/modules/auth/http"
	commentHttp "simple-video-server/app_server/modules/comment/http"
	emailHttp "simple-video-server/app_server/modules/email/http"
	followHttp "simple-video-server/app_server/modules/follow/http"
	operationLogHttp "simple-video-server/app_server/modules/log/operation_log/http"
	requestLogHttp "simple-video-server/app_server/modules/log/request_log/http"
	"simple-video-server/app_server/modules/test_limit"
	uploadHttp "simple-video-server/app_server/modules/upload/http"
	userCollection "simple-video-server/app_server/modules/user/collection/http"
	userLikeHttp "simple-video-server/app_server/modules/user/like/http"
	userHttp "simple-video-server/app_server/modules/user/user/http"
	videoHttp "simple-video-server/app_server/modules/video/http"
)

//var V1 *gin.RouterGroup

func SetupRoutes(r *gin.RouterGroup) {

	v1 := r.Group("/api/v1",
		middleware.TraceHandler,
		middleware.PreAuthorizeHandler,
		middleware.RequestLogHandler,
		middleware.ErrorHandler,
	)

	videoHttp.SetupRoutes(v1)

	authHttp.SetupRoutes(v1)

	userHttp.SetupRoutes(v1)

	//test_student.SetupRoutes(v1)

	userCollection.SetupRoutes(v1)

	userLikeHttp.SetupRoutes(v1)

	uploadHttp.SetupRoutes(v1)

	followHttp.SetupRoutes(v1)

	emailHttp.SetupRoutes(v1)

	commentHttp.SetupRoutes(v1)

	test_limit.SetupRoutes(v1)

	requestLogHttp.SetupRoutes(v1)

	operationLogHttp.SetupRoutes(v1)

}
