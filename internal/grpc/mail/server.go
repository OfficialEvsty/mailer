package mail

import (
	"context"
	mailv1 "github.com/OfficialEvsty/protos/gen/go/mailer"
	"google.golang.org/grpc"
	"mailer/internal/services/mail"
)

type mailServer struct {
	mailv1.UnimplementedMailServiceServer
	mailer mail.Mailer
}

func Register(gRPC *grpc.Server, mailer *mail.Mailer) {
	mailv1.RegisterMailServiceServer(gRPC, &mailServer{mailer: *mailer})
}

// Send message handler
// todo доделать
func (s *mailServer) Send(ctx context.Context, request *mailv1.SendMailRequest) (*mailv1.SendMailResponse, error) {
	return &mailv1.SendMailResponse{}, nil
}
