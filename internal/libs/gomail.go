package libs

import (
	"context"
	"fmt"
	"github.com/OfficialEvsty/mailer/domain/models"
	"github.com/OfficialEvsty/mailer/internal/config"
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
func (m *Mailer) Send(ctx context.Context, password string, mail *models.Mail) error {

	auth := smtp.PlainAuth("", m.From, password, m.SmtpHost)
	msg := "From: " + m.From + "\n" +
		"To: " + strings.Join(mail.To, ",") + "\n" +
		"Subject: " + mail.Subject + "\n\n" +
		mail.Body

	// отправка письма
	err := smtp.SendMail(m.SmtpHost+":"+m.SmtpPort, auth, m.From, mail.To, []byte(msg))
	if err != nil {
		return fmt.Errorf("ошибка при отправке письма: %w", err)
	}

	return nil
}
