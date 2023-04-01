package core

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"simple-video-server/common"
	"simple-video-server/pkg/app_ctx"
	"simple-video-server/pkg/log"
)

func ToHandler[T any](handler func(coreContext *Context) (T, error), moduleName string) func(c *gin.Context) {

	return func(c *gin.Context) {
		coreContext := new(Context)
		coreContext.Context = c

		// 注入uid
		uid, _ := app_ctx.GetUid(c)
		coreContext.AuthUID = uid
		// 注入是否已登录
		coreContext.Authorized = uid != nil
		if uid != nil {
			// 注入auth
			auth, ok := app_ctx.GetAuth(c)
			if ok {
				coreContext.Auth = auth
			}
		}

		// 注入logger
		logger := log.GetCtx(c.Request.Context())
		//coreContext.Log = logger.With(zap.String("module", moduleName), zap.Uint("uid", *uid))
		coreContext.Log = logger.With(zap.String("module", moduleName))

		// 注入trace id
		traceId, _ := app_ctx.GetTraceId(c)
		if traceId != nil {
			coreContext.TraceID = *traceId
		}

		// 执行handler, 得到返回的响应体和err
		resData, err := handler(coreContext)

		// TODO: 抛出err交由err_middleware 处理
		if err != nil {
			panic(err)
		}

		// 响应成功
		c.JSON(http.StatusOK, &common.AppResponse[any]{
			Code: 0,
			Msg:  "ok",
			Data: resData,
		})
	}
}
