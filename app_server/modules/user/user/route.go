package user

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	v1.GET("/user", core.ToHandler(api.GetAll, "up"))
}
