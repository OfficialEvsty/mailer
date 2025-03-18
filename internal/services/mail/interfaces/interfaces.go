package interfaces

import (
	"context"
	"github.com/OfficialEvsty/mailer/domain/models"
)

// MailProvider interface for sending mail on user's email
type MailProvider interface {
	Send(ctx context.Context, password string, mail *models.Mail) error
}
