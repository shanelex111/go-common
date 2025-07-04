package request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shanelex111/go-common/pkg/util"
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

		entry := logrus.WithFields(logrus.Fields{
			"request": &logEntryRequest{
				Method:    c.Request.Method,
				Path:      c.Request.RequestURI,
				Header:    c.Request.Header,
				Body:      reqBody,
				Timestamp: startTime,
			},
			"response": &logEntryResponse{
				Status:    status,
				Latency:   endTime - startTime,
				Header:    c.Writer.Header(),
				Body:      respBody.String(),
				Timestamp: endTime,
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

func SetUUID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := strings.TrimSpace(c.Request.Header.Get(XRequestIDKey))
		if len(requestID) == 0 {
			requestID = util.GetUUID()
		}
		c.Request.Header.Set(XRequestIDKey, requestID)
		c.Writer.Header().Set(XRequestIDKey, requestID)

		c.Next()
	}
}

func AuthAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) < 8 || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		accessToken := authHeader[7:]
		c.Set("access_token", accessToken)
		c.Next()
	}
}

const (
	TokenInfoKey = "token_info"
)

func AuthTokenInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenInfoHeader := c.GetHeader("Token-Info")
		if len(tokenInfoHeader) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var tokenInfo TokenInfo
		if err := json.Unmarshal([]byte(tokenInfoHeader), &tokenInfo); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set(TokenInfoKey, &tokenInfo)
		c.Next()
	}
}

func GetTokenInfo(c *gin.Context) *TokenInfo {
	tokenInfo, exists := c.Get(TokenInfoKey)
	if exists {
		return tokenInfo.(*TokenInfo)
	}
	return nil
}
