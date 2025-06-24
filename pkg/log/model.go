package log

import (
	"context"
	"encoding/json"
	"time"

	"maps"

	"github.com/shanelex111/go-common/pkg/request"
	"github.com/sirupsen/logrus"
)

type logger struct {
	Level      string    `json:"level"`
	Time       time.Time `json:"time"`
	XRequestID string    `json:"X-Request-ID"`
}

type customJSONFormatter struct {
}

func (f *customJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data))
	maps.Copy(data, entry.Data)

	logMeta := &logger{
		Level: entry.Level.String(),
		Time:  entry.Time,
	}

	// 从 context 中提取 request_id
	if ctxVal, ok := entry.Data["context"]; ok {
		if ctx, ok := ctxVal.(context.Context); ok {
			if rid, ok := ctx.Value(request.CtxRequestIDKey).(string); ok {
				// 加入 request_id 字段
				logMeta.XRequestID = rid
			}
		}
		// 删除 context 避免冗余输出
		delete(data, "context")
	}
	data["log"] = logMeta

	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return append(serialized, '\n'), nil
}
