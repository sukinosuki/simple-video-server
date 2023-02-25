package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-video-server/pkg/business_code"
	"simple-video-server/pkg/err_code"
)

func ok(c *gin.Context, data any) {
	c.JSON(http.StatusOK, &AppResponse[any]{
		Code: 0,
		Msg:  "ok",
		Data: data,
	})
}

func failBusiness(c *gin.Context, code business_code.BusinessCode) {
	c.JSON(http.StatusOK, &AppResponse[any]{
		Code:   code.Code(),
		Msg:    code.Message(),
		ErrMsg: code.Error(),
	})
}

func failErrCode(c *gin.Context, code err_code.ErrCode) {
	c.JSON(http.StatusBadRequest, &AppResponse[any]{
		Code:   code.Code,
		Msg:    code.Msg,
		ErrMsg: code.ErrMsg,
	})
}

func fail(c *gin.Context, err error) {

	c.JSON(http.StatusInternalServerError, &AppResponse[any]{
		Code:   business_code.ServerErr.Code(),
		Msg:    business_code.ServerErr.Message(),
		ErrMsg: err.Error(),
	})
}

func ToHandler[T any](f func(ginContext *gin.Context) (T, error)) func(c *gin.Context) {

	return func(c *gin.Context) {
		resData, err := f(c)

		if err != nil {
			// 响应业务message错误
			if code, ok := err.(business_code.BusinessCode); ok {
				failBusiness(c, code)
				return
			}

			// 响应自定义message错误
			if errCode, ok := err.(err_code.ErrCode); ok {
				failErrCode(c, errCode)
				return
			}

			fail(c, err)

			return
		}

		// 响应成功
		ok(c, resData)
	}
}
