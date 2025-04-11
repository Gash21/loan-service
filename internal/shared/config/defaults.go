package config

import "github.com/spf13/viper"

func loadDefaults() {
	// global config
	viper.SetDefault("ENV", "development")
	viper.SetDefault("PORT", 9000)
}
