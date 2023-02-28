package video

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/common"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	// 可能登录了
	//possibleAuth := v1.Group("", middleware.PreAuthorizeHandler)
	//v1.GET("/video/:id", middleware.PreAuthorizeHandler, common.ToHandler(Controller.GetById))
	v1.GET("/video/:id", common.ToHandler(Controller.GetById, "video"))

	// 需要登录
	shouldAuth := v1.Group("", middleware.AuthorizeHandler)
	// 优化toHandler
	//v1.POST("/video",  common.ToHandler(Controller.Add))
	{
		shouldAuth.POST("/video", common.ToHandler(Controller.Add, "video"))
	}

}
