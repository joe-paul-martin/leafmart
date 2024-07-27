package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServiceName string
	Env         string
}

func SetupConfig() (Config, error) {

	var config Config

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, nil
	}

	viper.Unmarshal(&config)

	return config, nil
}
