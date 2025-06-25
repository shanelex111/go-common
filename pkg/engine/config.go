package engine

import (
	"strings"

	"github.com/spf13/viper"
)

const (
	defaultKey = "engine"

	ModeDev     = "dev"
	ModeTest    = "test"
	ModePreProd = "pre-prod"
	ModeRelease = "release"
)

var (
	cfg *config
)

type config struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
	Name string `mapstructure:"name"`
}

func Init(v *viper.Viper) {
	initConfig(v)
}

func initConfig(v *viper.Viper) {
	cfg = &config{}
	if err := v.Sub(defaultKey).Unmarshal(cfg); err != nil {
		panic(err)
	}

}

func GetMode() string {
	switch strings.ToLower(cfg.Mode) {
	case ModeDev:
		return ModeDev
	case ModeTest:
		return ModeTest
	case ModePreProd:
		return ModePreProd
	case ModeRelease:
		return ModeRelease
	default:
		return ModeDev
	}
}

func GetPort() string {
	return cfg.Port
}

func GetName() string {
	return cfg.Name
}
