package core

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"simple-video-server/models"
	"strconv"
)

// Context 自定义context
// 参考: https://blog.csdn.net/qq_29514815/article/details/117228600
// https://www.jianshu.com/p/4ccdcb169345
type Context struct {
	*gin.Context
	AuthUID    *uint        // 认证用户id
	Auth       *models.User //认证用户
	Authorized bool         //是否已认证
	TraceID    string       // trace id
	Log        *zap.Logger  // log
}

// GetParamId 获取path上的id(只能是): /api/v1/video/:id
func (ctx *Context) GetParamId() uint {

	id, _ := strconv.Atoi(ctx.Param("id"))

	return uint(id)
}

// GetParamUID 获取path上的uid(只能是): /api/v1/user/:uid
func (ctx *Context) GetParamUID() uint {

	id, _ := strconv.Atoi(ctx.Param("uid"))

	return uint(id)
}
