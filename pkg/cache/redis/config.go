package redis

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

const (
	defaultKey = "redis"
)

var (
	RDB       *redis.ClusterClient
	Ctx       = context.Background()
	initOnce  sync.Once
	closeOnce sync.Once
)

type config struct {
	Addrs    []string `mapstructure:"addrs"`
	Username string   `mapstructure:"username"`
	Password string   `mapstructure:"password"`
	PoolSize int      `mapstructure:"pool_size"`
}

func Init(v *viper.Viper) {
	initOnce.Do(func() {
		cfg := initConfig(v)
		cfg.initClient()
	})
}

func initConfig(v *viper.Viper) *config {
	cfg := &config{}
	if err := v.Sub(defaultKey).Unmarshal(cfg); err != nil {
		panic(err)
	}
	return cfg

}

func (cfg *config) initClient() {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cfg.Addrs,
		Username: cfg.Username,
		Password: cfg.Password,
		PoolSize: cfg.PoolSize,
	})
	if err := rdb.Ping(Ctx).Err(); err != nil {
		panic(err)
	}
	rdb.AddHook(NewRedisHook())

	RDB = rdb
}

func Close() {
	closeOnce.Do(func() {
		if RDB != nil {
			_ = RDB.Close()
		}
	})
}
