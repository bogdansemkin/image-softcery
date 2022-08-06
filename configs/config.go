package configs

import "github.com/spf13/viper"

type Configs struct {
}

func (c *Configs) InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
