package middleware

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/pkg/app_ctx"
	"simple-video-server/pkg/business_code"
)

//const UidKey = "uid"
//
//const AuthErrKey = "auth_err"

var AuthorizeHandler = func(c *gin.Context) {
	_, exists := app_ctx.GetUid(c)

	if !exists {
		//authErr, exists := c.Get(AuthErrKey)
		authErr, exists := app_ctx.GetAuthorizeErr(c)

		if exists {
			// business code错误
			if err, ok := authErr.(business_code.BusinessCode); ok {
				panic(err)

			}

			panic(authErr)
		}

		// TODO:
		//panic(err_code.UnAuthorizedErr(business_code.BusinessCode.Message(), business_code.BusinessCode.Error()))
		panic(business_code.Unauthorized)
		//panic(err_code.UnAuthorizedErr())
	}

	c.Next()
}
