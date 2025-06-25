package config

import (
	"strings"

	"github.com/spf13/viper"
)

func Load(path, name string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(name)
	v.AddConfigPath(path)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}

func Parse(data string) (*viper.Viper, error) {
	v := viper.New()
	if err := v.ReadConfig(strings.NewReader(data)); err != nil {
		return nil, err
	}
	return v, nil
}
