package request

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func SetLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetLevel(logrus.InfoLevel)

		var (
			startTime = time.Now().UnixMilli()
			reqBody   string
			respBody  = bytes.NewBufferString("")
		)

		// 1. 读取请求体并恢复
		requestBodyBytes, _ := io.ReadAll(c.Request.Body)
		reqBody = string(requestBodyBytes)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBodyBytes))

		// 2. 设置响应体
		c.Writer = &bodyWriter{c.Writer, respBody}

		// 3. 处理请求
		c.Next()

		// 4. json格式打印控制台
		var (
			endTime = time.Now().UnixMilli()
			status  = c.Writer.Status()
		)

		entry := logrus.WithFields(logrus.Fields{
			"request": &logEntryRequest{
				Method: c.Request.Method,
				Path:   c.Request.URL.Path,
				Body:   reqBody,
				Header: c.Request.Header,
			},
			"response": &logEntryResponse{
				Status:  status,
				Latency: endTime - startTime,
				Header:  c.Writer.Header(),
				Body:    respBody.String(),
			},
		})

		switch {
		case status >= 500:
			entry.Error()
		case status >= 400:
			entry.Warn()
		default:
			entry.Info()
		}

	}
}
