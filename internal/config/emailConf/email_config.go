package emailconf

import "github.com/spf13/viper"

type Config struct {
	Host         string `mapstructure:"SMTP_HOST"`
	Port         int    `mapstructure:"SMTP_PORT"`
	SenderName   string `mapstructure:"SMTP_SENDER_NAME"`
	AuthEmail    string `mapstructure:"SMTP_AUTH_EMAIL"`
	AuthPassword string `mapstructure:"SMTP_AUTH_PASSWORD"`
}

func NewEmailConfig() (*Config, error) {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.AutomaticEnv()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
