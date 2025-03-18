package mail

import (
	"context"
	models2 "github.com/OfficialEvsty/mailer/domain/models"
	mail2 "github.com/OfficialEvsty/mailer/internal/services/mail"
	mailv1 "github.com/OfficialEvsty/protos/gen/go/mailer"
	"google.golang.org/grpc"
)

type mailServer struct {
	mailv1.UnimplementedMailServiceServer
	mailer *mail2.Mailer
}

func Register(gRPC *grpc.Server, mailer *mail2.Mailer) {
	mailv1.RegisterMailServiceServer(gRPC, &mailServer{mailer: mailer})
}

// SendMail message handler for sending mail to recipient
// todo доделать
func (s *mailServer) SendMail(ctx context.Context, request *mailv1.SendMailRequest) (*mailv1.SendMailResponse, error) {
	recipients := []string{request.GetEmailTo()}
	mailObject := models2.Mail{
		To:      recipients,
		Subject: request.GetSubject(),
		Body:    request.GetText(),
		Html:    request.GetHtml(),
	}
	err := s.mailer.SendMail(ctx, request.GetAuthPassword(), mailObject)
	if err != nil {
		return nil, err
	}
	return &mailv1.SendMailResponse{}, nil
}
