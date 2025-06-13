package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

const (
	defaultKey = "redis"
)

var (
	cfg *config
	RDB *redis.ClusterClient
	Ctx = context.Background()
)

type config struct {
	Addrs    []string `mapstructure:"addrs"`
	Username string   `mapstructure:"username"`
	Password string   `mapstructure:"password"`
	PoolSize int      `mapstructure:"pool_size"`
}

func Init(v *viper.Viper) {
	initConfig(v)
	initClient()
}

func initConfig(v *viper.Viper) {
	cfg = &config{}
	if err := v.Sub(defaultKey).Unmarshal(cfg); err != nil {
		panic(err)
	}

}

func initClient() {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cfg.Addrs,
		Username: cfg.Username,
		Password: cfg.Password,
		PoolSize: cfg.PoolSize,
	})
	if err := rdb.Ping(Ctx).Err(); err != nil {
		panic(err)
	}

	RDB = rdb
}
