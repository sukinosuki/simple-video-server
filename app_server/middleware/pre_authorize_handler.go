package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"simple-video-server/models"
	"simple-video-server/pkg/app_ctx"
	"simple-video-server/pkg/app_jwt"
	"simple-video-server/pkg/global"
	"simple-video-server/pkg/log"
)

var PreAuthorizeHandler = func(c *gin.Context) {
	token := app_ctx.GetHeaderAuthorize(c)
	traceId, _ := app_ctx.GetTraceId(c)

	var fields []zap.Field

	if traceId != nil {
		fields = append(fields, zap.String("trace-id", *traceId))
	}

	if token != "" {

		claims, err := app_jwt.AppJwt.Parse(token)

		if err != nil {
			app_ctx.SetAuthorizeErr(c, err)
		} else {
			//TODO: 从缓存获取用户信息
			app_ctx.SetUid(c, claims.UID)
			key := fmt.Sprintf("user:%d:info", claims.UID)
			result, err := global.RDB.Get(context.Background(), key).Result()

			if err == nil {
				var user models.User
				err := json.Unmarshal([]byte(result), &user)
				if err == nil {
					app_ctx.SetAuth(c, &user)
				} else {
					//	TODO:
				}
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
