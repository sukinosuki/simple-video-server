package http

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/modules/user/user/internal"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	api := internal.GetApi()

	v1.GET("/user", core.ToHandler(api.GetAll, "up"))
}
