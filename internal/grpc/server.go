package auth

import (
	"context"
	sandjmav1 "github.com/Sanchir01/protos_files_job/gen/go/auth"

	"google.golang.org/grpc"
)

type serverAPI struct {
	sandjmav1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	sandjmav1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(ctx context.Context, req *sandjmav1.LoginRequest) (*sandjmav1.LoginResponse, error) {
	panic("implement me")
}

func (s *serverAPI) Register(ctx context.Context, req *sandjmav1.RegisterRequest) (*sandjmav1.RegisterResponse, error) {
	panic("implement me")
}
