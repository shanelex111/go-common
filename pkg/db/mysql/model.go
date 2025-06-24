package mysql

import (
	"gorm.io/gorm/logger"
)

type GormLogger struct {
	LogLevel logger.LogLevel
}

type logEntry struct {
	StartAt int64  `json:"start_at"`
	Latency int64  `json:"latency"`
	Rows    int64  `json:"rows"`
	SQL     string `json:"sql"`
	EndAt   int64  `json:"end_at"`
	Msg     string `json:"msg"`
}
