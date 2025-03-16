package libs

import (
	"fmt"
	"mailer/domain/models"
	"mailer/internal/config"
	"net/smtp"
	"strings"
)

type Mailer struct {
	From     string
	SmtpHost string
	SmtpPort string
}

// New creates instance of Mailer
func New(cfg *config.MailerConfig) *Mailer {
	return &Mailer{
		From:     cfg.From,
		SmtpHost: cfg.SmtpHost,
		SmtpPort: cfg.SmtpPort,
	}
}

// auth email and sends mails to recipients
func (m *Mailer) Send(from string, password string, mail models.Mail) error {

	auth := smtp.PlainAuth("", from, password, m.SmtpHost)

	msg := "From: " + from + "\n" +
		"To: " + strings.Join(mail.To, ",") + "\n" +
		"Subject: " + mail.Subject + "\n\n" +
		mail.Body

	// отправка письма
	err := smtp.SendMail(m.SmtpHost+":"+m.SmtpPort, auth, from, mail.To, []byte(msg))
	if err != nil {
		return fmt.Errorf("ошибка при отправке письма: %w", err)
	}

	return nil
}
