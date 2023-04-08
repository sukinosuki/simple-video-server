package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"net/http"
	"simple-video-server/common"
	"simple-video-server/pkg/business_code"
	"simple-video-server/pkg/log"
	"simple-video-server/pkg/validation"
	"strings"
)

var ErrorHandler = func(c *gin.Context) {
	handlerField := zap.String("handler", "error_handler")
	defer func() {
		log := log.GetCtx(c.Request.Context())

		err := recover()

		if err != nil {

			//log.Error("全局错误 err: ", err)

			//if errors.Is(e, &mysql.MySQLError{}) {
			if e, ok := err.(*mysql.MySQLError); ok {
				log.Warn("mysql错误", zap.String("msg", e.Message), handlerField)

				if strings.Contains(e.Error(), "Duplicate") {
					c.AbortWithStatusJSON(http.StatusInternalServerError, common.AppResponse[any]{
						Code:   500,
						Msg:    "资源重复",
						ErrMsg: e.Error(),
					})

					return
				}
				c.AbortWithStatusJSON(http.StatusInternalServerError, common.AppResponse[any]{
					Code:   500,
					Msg:    "mysql错误",
					ErrMsg: e.Error(),
				})
				return
			}

			// 校验错误
			if errs, ok := err.(validator.ValidationErrors); ok {
				errorsMap := errs.Translate(validation.Trans)

				msg := ""

				for _, v := range errorsMap {
					msg = v
					break
				}
				log.Warn("参数校验错误", zap.String("msg", msg), handlerField)

				c.AbortWithStatusJSON(http.StatusBadRequest, &common.AppResponse[any]{
					Code:   business_code.RequestErr.Code(),
					Msg:    msg,
					ErrMsg: msg,
				})

				return
			}

			// 自定义的校验错误
			if e, ok := err.(*validation.ValidationError); ok {
				c.AbortWithStatusJSON(http.StatusBadRequest, &common.AppResponse[any]{
					Code: business_code.RequestErr.Code(),
					Msg:  e.Msg,
					//ErrMsg: msg,
				})

				return
			}

			//// 自定义错误
			//if errCode, ok := err.(err_code.ErrCode); ok {
			//	c.AbortWithStatusJSON(http.StatusOK, &common.AppResponse[any]{
			//		Code:   errCode.Code,
			//		Msg:    errCode.Msg,
			//		ErrMsg: errCode.ErrMsg,
			//	})
			//
			//	return
			//}

			// 自定义business错误
			if e, ok := err.(business_code.BusinessCode); ok {
				log.Warn("业务错误", zap.String("msg", e.Message()), zap.Int("code", e.Code()), handlerField)

				// TODO:
				c.AbortWithStatusJSON(http.StatusOK, common.AppResponse[any]{
					Code:   e.Code(),
					Msg:    e.Message(),
					ErrMsg: e.Error(),
				})

				return
			}

			// err为string的情况
			if _, ok := err.(string); ok {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code": 500,
					"msg":  err,
				})
				return
			}

			e := err.(error)

			log.Warn("未知错误 ", zap.String("err msg", e.Error()), handlerField)

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  e.Error(),
			})

			return
		}
	}()

	c.Next()
}
