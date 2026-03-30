package config

import (
	"github.com/spf13/viper"
)

func viper_config() {
	viper.SetConfigName("config")
	viper.SetDefault("APP_PORT", ":4043")
}
