package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"simple-video-server/core"
	"simple-video-server/pkg/app_ctx"
	"simple-video-server/pkg/log"
)

const traceIdKey = "trace-id"

var TraceHandler = func(c *gin.Context) {
	log.Logger.Info("请求limit开始")
	//err := core.Limiter.Wait(context.Background())
	//allow := core.Limiter.Allow()
	rateLimiter := core.NewRateLimiter()
	allow := rateLimiter.SlidingWindowTryAcquire(c.Request.URL.RawQuery)

	if allow {
		log.Logger.Info("请求limit结束 ")
	} else {
		panic(errors.New("限流, 请稍后重试"))
	}

	//if err != nil {
	//	fmt.Println("limiter wait 错误 ", err.Error())
	//	panic(err)
	//}

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
