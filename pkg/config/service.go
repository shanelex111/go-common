package config

import "github.com/spf13/viper"

func Load(path, name string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(name)
	v.AddConfigPath(path)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}
