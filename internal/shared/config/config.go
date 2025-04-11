package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Load() (GlobalConfig, error) {
	loadDefaults()

	// load from local env
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName("config")

	// load from environment variables
	viper.AutomaticEnv()

	// read config from sources (config file and environment variables)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("Error reading config file:", err)
			return GlobalConfig{}, err
		}
	}

	// unmarshal config
	c := GlobalConfig{}
	if err := viper.Unmarshal(&c); err != nil {
		return GlobalConfig{}, err
	}

	return c, nil
}
