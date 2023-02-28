package common

import (
	"github.com/gin-gonic/gin"
)

const traceIdKey = "trace-id"

const uidKey = "uid"

type BaseApiLog struct {
	Ctx    *gin.Context
	Name   string
	Module string
}

//
//func (l *BaseApiLog) Info(msg string, field ...zap.Field) {
//	traceId, _ := app_ctx.GetTraceId(l.Ctx)
//	uid, _ := app_ctx.GetUid(l.Ctx)
//
//	logger := log.Logger.With(
//		zap.String("module", l.Module),
//		zap.String("name", l.Name),
//		zap.String(traceIdKey, traceId),
//		zap.Uint(uidKey, uid),
//	)
//
//	logger.Info(msg, field...)
//}

type BaseApi struct {
	Ctx    *gin.Context
	Module string
	//Log    *BaseApiLog
}

//func (b *BaseApi) Profile(msg string, field ...zap.Field)  {
//
//}

func (b *BaseApi) Build(ctx *gin.Context, name string) {

	b.Ctx = ctx
	//b.Log.Name = name
	//b.Log.Module = b.Module
	//b.Log.Ctx = ctx
	//return BaseApiLog{
	//	Module: b.Module,
	//	Name:   name,
	//}
}
