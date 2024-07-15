package grpcapp

import (
	"google.golang.org/grpc"
	"log/slog"
)

type GrpcApp struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       string
}

func NewServer(log *slog.Logger, port string) *GrpcApp {
	gRPCServer := grpc.NewServer()
	return
}
