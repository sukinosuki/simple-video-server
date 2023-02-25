package video

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/common"
)

func SetupRoutes(group *gin.RouterGroup) {

	//group.POST("/user/register", Controller.Register)
	// 优化toHandler
	group.POST("/video", middleware.PreAuthorizeHandler, middleware.AuthorizeHandler, common.ToHandler(Controller.AddHandler))

	group.GET("/video/:id", middleware.PreAuthorizeHandler, common.ToHandler(Controller.GetById))

}
