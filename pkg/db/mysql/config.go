package mysql

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	defaultKey = "mysql"
)

var (
	cfg *config
	DB  *gorm.DB
)

type config struct {
	DSN string `mapstructure:"dsn"`
}

func initConfig(v *viper.Viper) {
	cfg = &config{}
	if err := v.Sub(defaultKey).Unmarshal(cfg); err != nil {
		panic(err)
	}

}

func initClient() {
	db, err := gorm.Open(mysql.Open(cfg.DSN))
	if err != nil {
		panic(err)
	}

	DB = db

}

func Init(v *viper.Viper) {
	initConfig(v)
	initClient()
}
