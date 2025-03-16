package mail

import (
	"context"
	"log/slog"
	"mailer/domain/models"
	"mailer/internal/services/mail/interfaces"
)

// Mailer sends messages on user's emails
type Mailer struct {
	log      *slog.Logger
	provider interfaces.MailProvider
}

// New creates instance of mailer
func New(logger *slog.Logger, provider interfaces.MailProvider) *Mailer {
	return &Mailer{log: logger, provider: provider}
}

// SendMail sends a mail on email, specified in mail object
func (m *Mailer) SendMail(ctx context.Context, authPassword string, mailToSend models.Mail) error {
	err := m.provider.Send(ctx, authPassword, &mailToSend)
	if err != nil {
		m.log.Error("Error sending mail", "error", err)
		return err
	}
	return nil
}
