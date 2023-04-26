package http

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	v1.GET("/log/request-log", core.ToHandler(_api.GetAll, "request_log"))
}
