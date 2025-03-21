package grpc

import (
	"fmt"
	mailrpc "github.com/OfficialEvsty/mailer/internal/grpc/mail"
	"github.com/OfficialEvsty/mailer/internal/services/mail"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

// New creates new gRPC server app
func New(
	log *slog.Logger,
	mailService *mail.Mailer,
	port int,
) *App {
	gRPCServer := grpc.NewServer()

	mailrpc.Register(gRPCServer, mailService)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

// MustRun runs gRPC server and panic if any occurs
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Run grpc server
func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(slog.String("op", op),
		slog.Int("port", a.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("starting gRPC server", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

// Stop grpc server
func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).Info("stopping gRPC server")
	a.gRPCServer.GracefulStop()
}
