package video

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	// 用户查询自己的视频列表 或者 用户查看其他用户的视频列表
	v1.GET("/user/:uid/video")
}
