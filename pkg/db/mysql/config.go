package mysql

import (
	"sync"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	defaultKey = "mysql"
)

var (
	DB        *gorm.DB
	initOnce  sync.Once
	closeOnce sync.Once
)

type config struct {
	DSN string `mapstructure:"dsn"`
}

func initConfig(v *viper.Viper) *config {
	cfg := &config{}
	if err := v.Sub(defaultKey).Unmarshal(cfg); err != nil {
		panic(err)
	}
	return cfg
}

func (cfg *config) initClient() {
	db, err := gorm.Open(mysql.Open(cfg.DSN))
	if err != nil {
		panic(err)
	}

	DB = db

}

func Init(v *viper.Viper) {
	initOnce.Do(func() {
		cfg := initConfig(v)
		cfg.initClient()
	})
}

func Close() {
	closeOnce.Do(func() {
		if DB != nil {
			sqlDB, _ := DB.DB()
			if sqlDB != nil {
				sqlDB.Close()
			}
		}
	})
}
