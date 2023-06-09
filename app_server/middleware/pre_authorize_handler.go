package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"simple-video-server/app_server/modules/auth"
	"simple-video-server/pkg/app_ctx"
	"simple-video-server/pkg/app_jwt"
	"simple-video-server/pkg/log"
)

var PreAuthorizeHandler = func(c *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("pre authorize handler错误 ", err)
		}
	}()
	token := app_ctx.GetHeaderAuthorize(c)
	traceId, _ := app_ctx.GetTraceId(c)

	userCache := auth.GetUserCache()

	var fields []zap.Field

	if traceId != nil {
		fields = append(fields, zap.String("trace-id", *traceId))
	}

	if token != "" {

		claims, err := app_jwt.AppJwt.Parse(token)

		if err != nil {
			app_ctx.SetAuthorizeErr(c, err)
		} else {
			// TODO: 从缓存获取用户信息
			// TODO: redis使用泛型
			// TODO: 从authCache获取
			app_ctx.SetUid(c, claims.UID)

			user, err := userCache.GetUser(claims.UID)

			if err == nil && user != nil {
				app_ctx.SetAuth(c, user)
			}

			fields = append(fields, zap.Uint("uid", claims.UID))
		}
	}

	//ctx, logger := log.AddCtx(c.Request.Context(), zap.String("trace-id", traceId))
	ctx, logger := log.AddCtx(c.Request.Context(), fields...)
	c.Request = c.Request.WithContext(ctx)

	logger.Info("预授权开始", zap.String("handler", "pre_authorize_handler"))

	c.Next()
}
