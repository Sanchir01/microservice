package auth

import (
	"context"
	sandjmav1 "github.com/Sanchir01/protos_files_job/gen/go/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

type serverAPI struct {
	sandjmav1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	sandjmav1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(ctx context.Context, req *sandjmav1.LoginRequest) (*sandjmav1.LoginResponse, error) {
	if req.Phone == "" {
		return nil, status.Error(codes.InvalidArgument, "phone is required")
	}
	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}
	return &sandjmav1.LoginResponse{
		UserUuid: req.Phone,
	}, nil
}

func (s *serverAPI) Register(ctx context.Context, req *sandjmav1.RegisterRequest) (*sandjmav1.RegisterResponse, error) {
	panic("implement me")
}
