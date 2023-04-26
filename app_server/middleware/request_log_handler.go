package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"simple-video-server/app_server/modules/log/request_log"
	"simple-video-server/pkg/log"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {

	w.body.Write(b)

	return w.ResponseWriter.Write(b)
}

func getPayload(c *gin.Context) string {
	payload := ""

	contentType := c.Request.Header.Get("Content-Type")

	if contentType == "application/json" {
		_bytes, err := io.ReadAll(c.Request.Body)

		// 重新赋值, 不然c.Next()之后的操作会读取不到body的值(参考: https://blog.csdn.net/testapl/article/details/122489448
		c.Request.Body = io.NopCloser(bytes.NewBuffer(_bytes))

		if err == nil {
			payload = string(_bytes)
		}
	}

	return payload
}

var RequestLogHandler = func(c *gin.Context) {

	log := log.GetCtx(c.Request.Context()).With(zap.String("handler", "request_log_handler"))

	payload := getPayload(c)

	log.Info("请求开始:", zap.String("url", c.Request.URL.RawPath))
	// TODO: log请求前信息

	blw := &bodyLogWriter{
		body:           bytes.NewBufferString(""),
		ResponseWriter: c.Writer,
	}

	c.Writer = blw

	c.Next()

	// 获取响应数据
	resData := blw.body.String()

	go func() {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println("记录请求日志失败 ", err)
			}
		}()

		logDao := request_log.GetDao()
		logDao.Add(c, payload, resData)
	}()
}
