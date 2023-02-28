package user

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/middleware"
	"simple-video-server/common"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	//v1.POST("/user/register", Api.Register)
	// 优化toHandler
	v1.POST("/user/register", common.ToHandler(Api.RegisterHandler, "user"))

	v1.POST("/user/login", common.ToHandler(Api.Login, "user"))

	//v1.GET("/user/profile", middleware2.PreAuthorizeHandler, middleware2.AuthorizeHandler, common.ToHandler(Api.Profile))
	// 需要登录
	shouldAuth := v1.Group("", middleware.AuthorizeHandler)

	{
		shouldAuth.GET("/user/profile", common.ToHandler(Api.Profile, "user"))
	}
}
