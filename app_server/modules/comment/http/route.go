package http

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	//api := GetApi()

	// 获取1级评论及对应的top n 2级评论
	v1.GET("/comment", core.ToHandler(_api.GetAll, "comment"))

	//获取1级评论下的所有二级评论
	v1.GET("/comment/:id/replies", core.ToHandler(_api.Get, "comment"))

	auth := v1.Group("", middleware.AuthorizeHandler)

	{
		// 新增评论
		auth.POST("/comment", core.ToHandler(_api.Add, "comment"))

		//auth.GET("/comment/:id")
		//
		//auth.PUT("/comment/:id")

		auth.DELETE("/comment/:id", core.ToHandler(_api.Delete, "comment"))
	}
}
