package mailing

import (
	"Go-Starter-Template/internal/utils"
	"strconv"

	"gopkg.in/gomail.v2"
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

func SendMail(toEmail string, subject string, body string) error {
	emailConfig := LoadMailConfig()

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", emailConfig.SMTPEmail)
	mailer.SetHeader("To", toEmail)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)
	port, err := strconv.Atoi(emailConfig.SMTPPort)
	if err != nil {
		return err
	}
	dialer := gomail.NewDialer(
		emailConfig.SMTPHost,
		port,
		emailConfig.SMTPEmail,
		emailConfig.SMTPPassword,
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}
