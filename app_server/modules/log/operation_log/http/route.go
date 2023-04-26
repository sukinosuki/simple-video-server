package http

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	v1.GET("/log/operation-log", core.ToHandler(_api.GetAll, "operation_log"))
}
