package emailconf

import (
	"Go-Starter-Template/internal/utils"
)

type MailConfig struct {
	AppURL       string
	SMTPHost     string
	SMTPPort     string
	SMTPSender   string
	SMTPEmail    string
	SMTPPassword string
}

func LoadMailConfig() MailConfig {
	return MailConfig{
		AppURL:       utils.GetEnv("APP_URL"),
		SMTPHost:     utils.GetEnv("SMTP_HOST"),
		SMTPPort:     utils.GetEnv("SMTP_PORT"),
		SMTPSender:   utils.GetEnv("SMTP_SENDER_NAME"),
		SMTPEmail:    utils.GetEnv("SMTP_AUTH_EMAIL"),
		SMTPPassword: utils.GetEnv("SMTP_AUTH_PASSWORD"),
	}
}
