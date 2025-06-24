package redis

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func NewRedisHook() redis.Hook {
	return &RedisHook{}
}

func (h *RedisHook) DialHook(next redis.DialHook) redis.DialHook {
	return next
}

func (h *RedisHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		start := time.Now()
		err := next(ctx, cmd)
		startAt := start.UnixMilli()
		endAt := time.Now().UnixMilli()

		logEntry := &logEntry{
			StartAt: startAt,
			Elapsed: endAt - startAt,
			Cmds:    []string{cmd.String()},
			EndAt:   endAt,
		}
		if err != nil {
			logEntry.Msg = err.Error()
		}

		entry := logrus.WithField("context", ctx).WithFields(logrus.Fields{
			"redis": logEntry,
		})

		if err != nil && !errors.Is(err, redis.Nil) {
			entry.Error()
		} else {
			entry.Info()
		}
		return err
	}
}

func (h *RedisHook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		start := time.Now()
		startAt := start.UnixMilli()

		err := next(ctx, cmds)

		result := make([]string, len(cmds))
		for i, cmd := range cmds {
			result[i] = cmd.String()
		}

		endAt := time.Now().UnixMilli()
		entry := logrus.WithFields(logrus.Fields{
			"redis": &logEntry{
				StartAt: startAt,
				Elapsed: endAt - startAt,
				Cmds:    result,
				EndAt:   endAt,
			},
		})

		if err != nil && !errors.Is(err, redis.Nil) {
			entry.Error()
		} else {
			entry.Info()
		}

		return err
	}
}
