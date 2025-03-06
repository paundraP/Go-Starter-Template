package mailing

import (
	emailconf "Go-Starter-Template/internal/config/email_config"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(toEmail string, subject string, body string) error {
	emailConfig := emailconf.LoadMailConfig()

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
