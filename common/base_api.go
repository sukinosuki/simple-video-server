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

type BaseApi struct {
	Ctx    *gin.Context
	Module string
	//Log    *BaseApiLog
}

func (b *BaseApi) Build(ctx *gin.Context, name string) {
	b.Ctx = ctx
}
