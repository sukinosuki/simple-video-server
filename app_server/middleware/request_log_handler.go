package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
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

var RequestLogHandler = func(c *gin.Context) {
	log := log.GetCtx(c.Request.Context())

	data := ""
	contentType := c.Request.Header.Get("Content-Type")
	if contentType == "application/json" {
		_bytes, err := io.ReadAll(c.Request.Body)

		// 重新赋值, 不然c.Next()之后的操作会读取不到body的值(参考: https://blog.csdn.net/testapl/article/details/122489448
		c.Request.Body = io.NopCloser(bytes.NewBuffer(_bytes))

		if err == nil {
			data = string(_bytes)
		}
	}

	//log.Profile(c, "%s | %s | data: %s", c.Request.Method, c.Request.RequestURI, data)
	log.Info("请求开始",
		zap.String("method", c.Request.Method),
		zap.String("uri", c.Request.RequestURI),
		zap.String("content-type", contentType),
		zap.String("data", data))

	blw := &bodyLogWriter{
		body:           bytes.NewBufferString(""),
		ResponseWriter: c.Writer,
	}

	c.Writer = blw

	c.Next()

	statusCode := c.Writer.Status()

	resData := blw.body.String()

	log.Info("请求结束",
		zap.Int("status_code", statusCode),
		zap.String("res", resData),
	)
}
