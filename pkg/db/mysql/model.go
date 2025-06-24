package mysql

import (
	"gorm.io/gorm/logger"
)

type GormLogger struct {
	LogLevel logger.LogLevel
}
