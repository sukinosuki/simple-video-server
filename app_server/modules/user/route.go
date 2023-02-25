package user

import (
	"github.com/gin-gonic/gin"
	middleware2 "simple-video-server/app_server/middleware"
	"simple-video-server/common"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	//v1.POST("/user/register", Controller.Register)
	// 优化toHandler
	v1.POST("/user/register", common.ToHandler(Controller.RegisterHandler))

	v1.POST("/user/login", common.ToHandler(Controller.Login))

	v1.GET("/user/profile", middleware2.PreAuthorizeHandler, middleware2.AuthorizeHandler, common.ToHandler(Controller.Info))
}
