package email

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	//preAuth := v1.Group("", middleware.PreAuthorizeHandler)

	v1.POST("/send-email", core.ToHandler(Api.SendEmail, "email"))
}
