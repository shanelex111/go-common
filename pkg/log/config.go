package log

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	defaultKey = "log"
)

type config struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

func initConfig(v *viper.Viper) *config {
	cfg := &config{
		Level:      LevelInfo,
		Filename:   "./logs/app.log",
		MaxSize:    100,
		MaxBackups: 10,
		MaxAge:     7,
		Compress:   true,
	}
	if v.IsSet(defaultKey) {
		if err := v.Sub(defaultKey).Unmarshal(cfg); err != nil {
			panic(err)
		}
	}

	return cfg
}
func (cfg *config) initClient() {
	// 1. 输出路径
	logFile := &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	logrus.SetOutput(multiWriter)

	// 2. 日志级别
	lvl, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		lvl = logrus.InfoLevel
	}
	logrus.SetLevel(lvl)

	// 3. 格式化
	logrus.SetFormatter(&customJSONFormatter{})

}
func Init(v *viper.Viper) {
	cfg := initConfig(v)
	cfg.initClient()
}
