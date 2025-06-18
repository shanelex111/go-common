package server

import "github.com/spf13/viper"

func Init(v *viper.Viper, inits ...func(v *viper.Viper)) {
	for i := range inits {
		go inits[i](v)
	}
}

func Run(runs ...func()) {
	for i := range runs {
		runs[i]()
	}
}

func Load(v *viper.Viper, loads ...func(v *viper.Viper)) {
	for i := range loads {
		loads[i](v)
	}
}
