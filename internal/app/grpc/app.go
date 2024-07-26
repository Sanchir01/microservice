package grpcapp

import (
	"fmt"
	authgrpc "github.com/Sanchir01/microservice/internal/grpc/auth"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"strconv"
)

type GrpcApp struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewServer(log *slog.Logger, port int, authService authgrpc.IAuth) *GrpcApp {
	gRPCServer := grpc.NewServer()
	authgrpc.Register(gRPCServer, authService)
	return &GrpcApp{
		log: log, gRPCServer: gRPCServer, port: port,
	}
}
func (g *GrpcApp) MustRun() {
	if err := g.Run(); err != nil {
		panic(err)
	}
}

func (g *GrpcApp) Run() error {
	const op = "grpcapp.Run"
	log := g.log.With(
		slog.String("op", op),
		slog.String("port", strconv.Itoa(g.port)),
	)
	log.Info("starting grpc server")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", g.port))

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server is running", slog.String("address", l.Addr().String()))

	if err := g.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (g *GrpcApp) Stop() {
	const op = "grpcapp.Stop"

	g.log.With(slog.String("op", op)).Info("stopping grpc server", slog.Int("port", g.port))

	g.gRPCServer.GracefulStop()

}
