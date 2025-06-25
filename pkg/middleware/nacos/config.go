package nacos

import (
	"fmt"
	"os"
	"sync"

	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
)

const (
	defaultKey = "nacos"
)

var (
	initOnce sync.Once
)

type config struct {
	ServerConfigs []serverConfig `mapstructure:"server_configs"`
	ClientConfig  clientConfig   `mapstructure:"client_config"`
}

type serverConfig struct {
	IpAddr string `mapstructure:"ip_addr"`
	Port   uint64 `mapstructure:"port"`
}
type clientConfig struct {
	NamespaceId string `mapstructure:"namespace_id"`
	DataID      string `mapstructure:"data_id"`
	Group       string `mapstructure:"group"`
}

func Init(v *viper.Viper) {
	initOnce.Do(func() {
		cfg := initConfig(v)
		cfg.initClient()
	})

}

func initConfig(v *viper.Viper) *config {
	cfg := &config{}
	if err := v.Sub(defaultKey).Unmarshal(cfg); err != nil {
		panic(err)
	}
	return cfg
}

func (c *config) initClient() {
	sc := []constant.ServerConfig{}

	if len(c.ServerConfigs) != 0 {
		for _, serverConfig := range c.ServerConfigs {
			sc = append(sc, constant.ServerConfig{
				IpAddr: serverConfig.IpAddr,
				Port:   serverConfig.Port,
			})
		}
	}

	cc := constant.NewClientConfig(
		constant.WithNamespaceId(c.ClientConfig.NamespaceId),
	)

	configClient, err := clients.CreateConfigClient(map[string]any{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}

	// 监听配置变化
	if err = configClient.ListenConfig(vo.ConfigParam{
		DataId: c.ClientConfig.DataID,
		Group:  c.ClientConfig.Group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("[Nacos] Config Changed, restarting app...")
			os.Exit(1)
		},
	}); err != nil {
		panic(err)
	}
}
