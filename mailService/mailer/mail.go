package mailer

import (
	"mailService/config"
	"net/smtp"
)

func SendMail(toMail string, subject string, body string, config config.Config) error {
	// Sender Data
	senderMail := config.SENDER_GMAIL
	senderPassword := config.SENDER_PASSWORD

	// Receiver Email
	to := []string{
		toMail,
	}

	// smtp configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// message
	message := []byte(body)

	// auth
	auth := smtp.PlainAuth("", senderMail, senderPassword, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderMail, to, message)
	if err != nil {
		return err
	}

	return nil
}
