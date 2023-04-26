package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"simple-video-server/app_server/modules/log/operation_log"
	"simple-video-server/models"
	"strconv"
	"time"
)

// Context 自定义context
// 参考: https://blog.csdn.net/qq_29514815/article/details/117228600
// https://www.jianshu.com/p/4ccdcb169345
type Context struct {
	*gin.Context
	AuthUID    *uint        // 认证用户id
	Auth       *models.User // 认证用户
	Authorized bool         // 是否已认证
	TraceID    string       // trace id
	Log        *zap.Logger  // log
	Module     string
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

func (ctx *Context) PanicIfErr(err error, handler string, msg string) {

	if err == nil {
		return
	}

	// TODO: 在协程里记录日志
	go func() {
		defer func() {
			_err := recover()
			if _err != nil {

			}
			fmt.Println("_err ", _err)
		}()
		ctx.Log.Error(msg, zap.String("handler", handler), zap.Error(err))

		// TODO: clickhouse记录日志
		dao := operation_log.GetDao()
		var uid uint = 0
		// TODO: ctx.AuthID为nil时, *ctx.AuthID拿到的值是什么
		if ctx.AuthUID != nil {
			uid = *ctx.AuthUID
		}

		operationLog := &models.OperationLog{
			UID:        uint64(uid),
			CreateTime: time.Now(),
			Module:     ctx.Module,
			Handler:    handler,
			TraceID:    ctx.TraceID,
			Err:        err.Error(),
			Msg:        msg,
			Level:      "ERROR",
		}

		dao.Add(operationLog)
	}()

	panic(err)
}

func (ctx *Context) PanicIf(ok bool, err error, handler string, msg string) {
	if !ok {
		return
	}

	go func() {
		defer func() {
			_err := recover()
			if _err != nil {

			}
			fmt.Println("_err ", _err)
		}()

		ctx.Log.Error(msg, zap.String("handler", handler), zap.Error(err))

		// TODO: clickhouse记录日志
		dao := operation_log.GetDao()
		var uid uint = 0
		// TODO: ctx.AuthID为nil时, *ctx.AuthID拿到的值是什么
		if ctx.AuthUID != nil {
			uid = *ctx.AuthUID
		}

		operationLog := &models.OperationLog{
			UID:        uint64(uid),
			CreateTime: time.Now(),
			Module:     ctx.Module,
			Handler:    handler,
			TraceID:    ctx.TraceID,
			Err:        err.Error(),
			Msg:        msg,
			Level:      "ERROR",
		}

		dao.Add(operationLog)
	}()

	panic(err)
}

func (ctx *Context) Panic(err error, handler string, msg string) {
	if err == nil {
		return
	}

	go func() {
		defer func() {
			_err := recover()
			if _err != nil {

			}
			fmt.Println("_err ", _err)
		}()
		ctx.Log.Error(msg, zap.String("handler", handler), zap.Error(err))

		// TODO: clickhouse记录日志
		dao := operation_log.GetDao()
		var uid uint = 0
		// TODO: ctx.AuthID为nil时, *ctx.AuthID拿到的值是什么
		if ctx.AuthUID != nil {
			uid = *ctx.AuthUID
		}

		operationLog := &models.OperationLog{
			UID:        uint64(uid),
			CreateTime: time.Now(),
			Module:     ctx.Module,
			Handler:    handler,
			TraceID:    ctx.TraceID,
			Err:        err.Error(),
			Msg:        msg,
			Level:      "ERROR",
		}

		dao.Add(operationLog)
	}()

	panic(err)
}

func (ctx *Context) Info(msg string, handlerName string, fields ...zap.Field) {
	go func() {
		defer func() {
			_err := recover()
			if _err != nil {

			}
			fmt.Println("_err ", _err)
		}()
		log := ctx.Log.With(zap.String("handler", handlerName))

		log.Info(msg, fields...)
		// TODO: clickhouse记录日志
		dao := operation_log.GetDao()
		var uid uint = 0
		// TODO: ctx.AuthID为nil时, *ctx.AuthID拿到的值是什么
		if ctx.AuthUID != nil {
			uid = *ctx.AuthUID
		}

		operationLog := &models.OperationLog{
			UID:        uint64(uid),
			CreateTime: time.Now(),
			Module:     ctx.Module,
			Handler:    handlerName,
			TraceID:    ctx.TraceID,
			Msg:        msg,
			Level:      "INFO",
		}

		dao.Add(operationLog)
	}()
}

func (ctx *Context) InfoSync(msg string, fields ...zap.Field) {

	ctx.Log.Info(msg, fields...)
	// TODO: clickhouse记录log
}
