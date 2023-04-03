package email_http

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/modules/email/internal"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	//preAuth := v1.Group("", middleware.PreAuthorizeHandler)

	api := internal.GetApi()

	v1.POST("/send-email", core.ToHandler(api.SendEmail, "email"))
}
