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
	UID        *uint // 用户id
	Auth       *models.User
	Authorized bool
	TraceID    string      // trace id
	Log        *zap.Logger // log
}

// GetId 获取path上的id(只能是): /api/v1/user/:id
func (ctx *Context) GetId() uint {

	id, _ := strconv.Atoi(ctx.Param("id"))

	return uint(id)
}
