package request

import (
	"bytes"
	"io"

	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	XRequestIDKey = "X-Request-ID"
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

		entry := &logEntry{
			Request: &logEntryRequest{
				Method:    c.Request.Method,
				Path:      c.Request.RequestURI,
				Header:    c.Request.Header,
				Body:      reqBody,
				Timestamp: startTime,
			},
			Response: &logEntryResponse{
				Status:    status,
				Latency:   endTime - startTime,
				Header:    c.Writer.Header(),
				Body:      respBody.String(),
				Timestamp: endTime,
			},
		}

		switch {
		case status >= 500:
			logrus.Error(entry)
		case status >= 400:
			logrus.Warn(entry)
		default:
			logrus.Info(entry)
		}

	}
}

func SetUUID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := strings.TrimSpace(c.Request.Header.Get(XRequestIDKey))
		if len(requestID) == 0 {
			requestID = uuid.New().String()
		}
		c.Request.Header.Set(XRequestIDKey, requestID)
		c.Writer.Header().Set(XRequestIDKey, requestID)
		c.Next()
	}
}
