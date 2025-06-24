package mysql

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
)

func (gl *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if gl.LogLevel <= logger.Silent {
		return
	}

	sql, rows := fc()

	elapsed := time.Since(begin)
	startAt := begin.UnixMilli()
	endAt := time.Now().UnixMilli()

	entry := logrus.WithFields(logrus.Fields{
		"mysql": &logEntry{
			StartAt: startAt,
			Elapsed: endAt - startAt,
			Rows:    rows,
			SQL:     sql,
			EndAt:   endAt,
		},
	})
	switch {
	case err != nil && gl.LogLevel >= logger.Error:
		entry.Error()
	case elapsed > 200*time.Millisecond && gl.LogLevel >= logger.Warn:
		entry.Warn()
	case gl.LogLevel >= logger.Info:
		entry.Info()
	}

}

func (gl *GormLogger) Error(ctx context.Context, msg string, data ...any) {
	if gl.LogLevel >= logger.Error {
		logrus.WithContext(ctx).Errorf(msg, data...)
	}
}

func (gl *GormLogger) Info(ctx context.Context, msg string, data ...any) {
	if gl.LogLevel >= logger.Info {
		logrus.WithContext(ctx).Infof(msg, data...)
	}
}

func (gl *GormLogger) Warn(ctx context.Context, msg string, data ...any) {
	if gl.LogLevel >= logger.Warn {
		logrus.WithContext(ctx).Warnf(msg, data...)
	}
}

func (gl *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *gl
	newLogger.LogLevel = level
	return &newLogger
}
func GormLogLevel() logger.LogLevel {
	switch logrus.GetLevel() {
	case logrus.DebugLevel, logrus.TraceLevel:
		return logger.Info
	case logrus.InfoLevel:
		return logger.Info
	case logrus.WarnLevel:
		return logger.Warn
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return logger.Error
	default:
		return logger.Silent
	}
}
