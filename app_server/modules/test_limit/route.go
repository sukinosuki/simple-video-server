package test_limit

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	v1.GET("/test-limit", core.ToHandler(func(c *core.Context) (bool, error) {

		return true, nil
	}, "test limit"))
}
