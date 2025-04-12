package config

import "github.com/spf13/viper"

func loadDefaults() {
	// global config
	viper.SetDefault("ENV", "development")
	viper.SetDefault("PORT", 9000)
	viper.SetDefault("DB_NAME", "amartha.db")
	viper.SetDefault("AUTO_MIGRATE", true)
	viper.SetDefault("PUBLIC_URL", "http://localhost:9000")
}
