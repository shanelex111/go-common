package mysql

import (
	"gorm.io/gorm/logger"
)

type GormLogger struct {
	LogLevel logger.LogLevel
}

type logEntry struct {
	Elapsed int64  `json:"elapsed"`
	Rows    int64  `json:"rows"`
	SQL     string `json:"sql"`
}
