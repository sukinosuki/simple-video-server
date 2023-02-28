package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"simple-video-server/pkg/app_ctx"
)

const traceIdKey = "trace-id"

var TraceHandler = func(c *gin.Context) {

	uuid, err := uuid.NewUUID()
	traceId := ""

	if err == nil {
		traceId = uuid.String()

		app_ctx.SetTraceId(c, traceId)
	}

	// 响应header添加trace-id
	c.Writer.Header().Set(traceIdKey, traceId)

	c.Next()
}