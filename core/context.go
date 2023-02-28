package core

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Context 自定义context
// 参考: https://blog.csdn.net/qq_29514815/article/details/117228600
// https://www.jianshu.com/p/4ccdcb169345
type Context struct {
	*gin.Context
	UID     *uint
	TraceID string
	Log     *zap.Logger
}
