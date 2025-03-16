package interfaces

import (
	"context"
	"mailer/domain/models"
)

// MailProvider interface for sending mail on user's email
type MailProvider interface {
	Send(ctx context.Context, password string, mail *models.Mail) error
}
