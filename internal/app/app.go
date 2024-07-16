package app

import (
	grpcapp "github.com/Sanchir01/microservice/internal/app/grpc"
	"log/slog"
)

type App struct {
	GrpcSrv *grpcapp.GrpcApp
}

func NewAppSrv(log *slog.Logger, grpcPort int, storagePath string) *App {
	grpcApp := grpcapp.NewServer(log, grpcPort)
	return &App{
		GrpcSrv: grpcApp,
	}
}
