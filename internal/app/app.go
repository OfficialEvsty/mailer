package app

import (
	"log/slog"
	grpcapp "mailer/internal/app/grpc"
	"mailer/internal/config"
	"mailer/internal/libs"
	"mailer/internal/services/mail"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(logger *slog.Logger, cfg *config.MailerConfig) *App {
	mailer := libs.New(cfg)
	mailService := mail.New(logger, mailer)
	grpcServer := grpcapp.New(logger, mailService, cfg.GRPC.Port)
	return &App{
		GRPCSrv: grpcServer,
	}
}
