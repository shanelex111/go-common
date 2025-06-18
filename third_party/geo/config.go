package geo

import (
	"github.com/oschwald/geoip2-golang"
	"github.com/spf13/viper"
)

const (
	defaultKey = "geo"
)

var (
	DB *geoip2.Reader
)

type config struct {
	Path string `mapstructure:"path"`
}

func Init(v *viper.Viper) {
	cfg := initConfig(v)
	cfg.initClient()
}

func initConfig(v *viper.Viper) *config {
	cfg := &config{}
	if err := v.Sub(defaultKey).Unmarshal(cfg); err != nil {
		panic(err)
	}
	return cfg
}

func (cfg *config) initClient() {
	db, err := geoip2.Open(cfg.Path)
	if err != nil {
		panic(err)
	}

	DB = db
}
