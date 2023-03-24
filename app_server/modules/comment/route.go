package comment

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	// 获取1级评论及对应的top n 2级评论
	v1.GET("/comment", core.ToHandler(Api.GetAll, "comment"))

	//获取1级评论下的所有二级评论
	v1.GET("/comment/:id/replies", core.ToHandler(Api.Get, "comment"))

	auth := v1.Group("", middleware.AuthorizeHandler)

	{
		// 新增评论
		auth.POST("/comment", core.ToHandler(Api.Add, "comment"))

		//auth.GET("/comment/:id")
		//
		//auth.PUT("/comment/:id")

		auth.DELETE("/comment/:id", core.ToHandler(Api.Delete, "comment"))
	}
}
