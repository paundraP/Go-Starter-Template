package mailing

import (
	emailconf "Go-Starter-Template/internal/config/emailConf"

	"gopkg.in/gomail.v2"
)

func SendMail(toEmail string, subject string, body string) error {
	emailConfig, err := emailconf.NewEmailConfig()
	if err != nil {
		return err
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", emailConfig.AuthEmail)
	mailer.SetHeader("To", toEmail)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		emailConfig.Host,
		emailConfig.Port,
		emailConfig.AuthEmail,
		emailConfig.AuthPassword,
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}
