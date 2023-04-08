package http

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	//api := GetApi()

	v1.GET("/user", core.ToHandler(_api.GetAll, "up"))
}
