package config

import "github.com/spf13/viper"

func Load(path, name string) (err, *viper.Viper) {
	v := viper.New()
	v.SetConfigName(name)
	v.AddConfigPath(path)
	if err := v.ReadInConfig(); err != nil {
		return err, nil
	}
	return nil, v
}
