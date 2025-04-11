package config

type GlobalConfig struct {
	Debug       bool   `mapstructure:"DEBUG"`
	LogLevel    string `mapstructure:"LOG_LEVEL"`
	Port        int    `mapstructure:"PORT"`
	ServiceName string `mapstructure:"SERVICE_NAME"`
}
