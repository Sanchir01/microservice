package app

import (
	grpcapp "github.com/Sanchir01/microservice/internal/app/grpc"
	"github.com/Sanchir01/microservice/internal/config"
	"github.com/Sanchir01/microservice/internal/database/postgres"
	"github.com/Sanchir01/microservice/internal/services/auth"
	"github.com/Sanchir01/microservice/pkg/db/connect"
	"log/slog"
)

type App struct {
	GrpcSrv *grpcapp.GrpcApp
}

func NewAppSrv(log *slog.Logger, grpcPort int, cfg *config.Config) *App {
	db := connect.PostgresMain(cfg, log)
	storage := postgres.NewStorePostgres(db)
	authService := auth.New(log, storage, storage)

	grpcApp := grpcapp.NewServer(log, grpcPort, authService)

	return &App{
		GrpcSrv: grpcApp,
	}
}
