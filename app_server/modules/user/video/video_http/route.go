package video_http

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/app_server/modules/user/video/internal"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	api := internal.GetApi()

	// 用户查询自己的视频列表 或者 用户查看其他用户的视频列表
	v1.GET("/user/:uid/video", core.ToHandler(api.GetAll, "user_video"))
}
