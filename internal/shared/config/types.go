package config

type GlobalConfig struct {
	Debug       bool   `mapstructure:"DEBUG"`
	LogLevel    string `mapstructure:"LOG_LEVEL"`
	Port        int    `mapstructure:"PORT"`
	ServiceName string `mapstructure:"SERVICE_NAME"`
	DBName      string `mapstructure:"DB_NAME"`
	AutoMigrate bool   `mapstructure:"AUTO_MIGRATE"`
	PublicURL   string `mapstructure:"PUBLIC_URL"`

	MailerType     string `mapstructure:"MAILER_TYPE"`
	MailerHostname string `mapstructure:"MAILER_HOSTNAME"`
	MailerUsername string `mapstructure:"MAILER_USERNAME"`
	MailerPassword string `mapstructure:"MAILER_PASSWORD"`
	MailerSender   string `mapstructure:"MAILER_SENDER"`
	MailerPort     int    `mapstructure:"MAILER_PORT"`
	MailerTLS      bool   `mapstructure:"MAILER_TLS"`
}
