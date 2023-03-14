package comment

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	//
	v1.GET("/:media/:media_id/comment", core.ToHandler(Api.Get, "comment"))

	auth := v1.Group("", middleware.AuthorizeHandler)

	{
		// 新增评论
		auth.POST("/:media/:media_id/comment", core.ToHandler(Api.Add, "comment"))

		auth.GET("/:media/:media_id/comment/:id")

		auth.PUT("/:media/:media_id/comment/:id")

		auth.DELETE("/:media/:media_id/comment/:id", core.ToHandler(Api.Delete, "comment"))
	}
}
