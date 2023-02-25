package middleware

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/pkg/app_ctx"
	"simple-video-server/pkg/app_jwt"
	"simple-video-server/pkg/log"
)

var PreAuthorizeHandler = func(c *gin.Context) {
	log.Info(c, "[PreAuthorizeHandler] start ")

	token := app_ctx.GetHeaderAuthorize(c)

	if token != "" {

		claims, err := app_jwt.AppJwt.Parse(token)

		if err != nil {
			app_ctx.SetAuthorizeErr(c, err)
		} else {
			app_ctx.SetUid(c, claims.UID)
		}
	}

	c.Next()
}
