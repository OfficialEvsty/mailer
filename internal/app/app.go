package app

import (
	"github.com/OfficialEvsty/mailer/internal/app/grpc"
	"github.com/OfficialEvsty/mailer/internal/config"
	"github.com/OfficialEvsty/mailer/internal/libs"
	mailservice "github.com/OfficialEvsty/mailer/internal/services/mail"
	"log/slog"
)

type App struct {
	GRPCSrv *grpc.App
}

func New(logger *slog.Logger, cfg *config.MailerConfig) *App {
	mailer := libs.New(cfg)
	mailService := mailservice.New(logger, mailer)
	grpcServer := grpc.New(logger, mailService, cfg.GRPC.Port)
	return &App{
		GRPCSrv: grpcServer,
	}
}
