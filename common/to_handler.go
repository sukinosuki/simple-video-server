package common

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"simple-video-server/core"
	"simple-video-server/pkg/app_ctx"
	"simple-video-server/pkg/log"
)

//func ok(c *gin.Context, data any) {
//	c.JSON(http.StatusOK, &AppResponse[any]{
//		Code: 0,
//		Msg:  "ok",
//		Data: data,
//	})
//}

//func failBusiness(c *gin.Context, code business_code.BusinessCode) {
//	c.JSON(http.StatusOK, &AppResponse[any]{
//		Code:   code.Code(),
//		Msg:    code.Message(),
//		ErrMsg: code.Error(),
//	})
//}
//
//func failErrCode(c *gin.Context, code err_code.ErrCode) {
//	c.JSON(http.StatusBadRequest, &AppResponse[any]{
//		Code:   code.Code,
//		Msg:    code.Msg,
//		ErrMsg: code.ErrMsg,
//	})
//}
//
//func fail(c *gin.Context, err error) {
//
//	c.JSON(http.StatusInternalServerError, &AppResponse[any]{
//		Code:   business_code.ServerErr.Code(),
//		Msg:    business_code.ServerErr.Message(),
//		ErrMsg: err.Error(),
//	})
//}

// func ToHandler[T any](f func(ginContext *gin.Context) (T, error)) func(c *gin.Context) {
func ToHandler[T any](handler func(coreContext *core.Context) (T, error), moduleName string) func(c *gin.Context) {

	return func(c *gin.Context) {
		coreContext := new(core.Context)
		coreContext.Context = c

		// 注入uid
		uid, _ := app_ctx.GetUid(c)
		coreContext.UID = uid

		// 注入logger
		logger := log.GetCtx(c.Request.Context())
		coreContext.Log = logger.With(zap.String("module", moduleName))

		// 注入trace id
		traceId, _ := app_ctx.GetTraceId(c)
		coreContext.TraceID = traceId

		// 执行handler, 得到返回的响应体和err
		resData, err := handler(coreContext)

		// 交由err middleware 处理
		if err != nil {
			panic(err)
			//// 响应业务message错误
			//if code, ok := err.(business_code.BusinessCode); ok {
			//	failBusiness(c, code)
			//	return
			//}
			//
			//// 响应自定义message错误
			//if errCode, ok := err.(err_code.ErrCode); ok {
			//	failErrCode(c, errCode)
			//	return
			//}
			//
			//fail(c, err)
			//
			//return
		}

		// 响应成功
		c.JSON(http.StatusOK, &AppResponse[any]{
			Code: 0,
			Msg:  "ok",
			Data: resData,
		})
	}
}
