package geo

import (
	"github.com/oschwald/geoip2-golang"
	"github.com/spf13/viper"
)

const (
	defaultKey = "geo"
)

var (
	cfg *config
	DB  *geoip2.Reader
)

type config struct {
	Path string `mapstructure:"path"`
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
	db, err := geoip2.Open(cfg.Path)
	if err != nil {
		panic(err)
	}

	DB = db
}
