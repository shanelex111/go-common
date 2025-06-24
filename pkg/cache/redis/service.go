package redis

import (
	"context"

	"github.com/sirupsen/logrus"
)

func NewRedisLogger() *RedisLogger {
	return &RedisLogger{}
}

func (l *RedisLogger) Printf(ctx context.Context, format string, args ...any) {
	if logrus.GetLevel() >= logrus.InfoLevel {
		logrus.WithContext(ctx).Infof(format, args...)
	}
}
