package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"simple-video-server/pkg/app_ctx"
	"simple-video-server/pkg/app_jwt"
	"simple-video-server/pkg/log"
)

var PreAuthorizeHandler = func(c *gin.Context) {
	//log.Info(c, "[PreAuthorizeHandler] start ")
	//log := log.GetCtx(c.Request.Context())
	//log.Info("[PreAuthorizeHandler] start ")

	token := app_ctx.GetHeaderAuthorize(c)
	traceId, _ := app_ctx.GetTraceId(c)

	fields := []zap.Field{}

	fields = append(fields, zap.String("trace-id", traceId))
	if token != "" {

		claims, err := app_jwt.AppJwt.Parse(token)

		if err != nil {
			app_ctx.SetAuthorizeErr(c, err)
		} else {
			app_ctx.SetUid(c, claims.UID)
			fields = append(fields, zap.Uint("uid", claims.UID))
		}
	}

	//ctx, logger := log.AddCtx(c.Request.Context(), zap.String("trace-id", traceId))
	ctx, logger := log.AddCtx(c.Request.Context(), fields...)
	c.Request = c.Request.WithContext(ctx)

	logger.Info("add trace id success ", zap.String("trace id ", traceId))

	c.Next()
}
