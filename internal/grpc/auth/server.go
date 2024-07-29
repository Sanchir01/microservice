package authgrpc

import (
	"context"
	sandjmav1 "github.com/Sanchir01/protos_files_job/pkg/gen/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

type IAuth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
		appID int,
	) (token string, err error)
	RegisterNewUser(
		ctx context.Context,
		email string,
		password string,
	) (userID string, err error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

type serverAPI struct {
	sandjmav1.UnimplementedAuthServer
	auth IAuth
}

func Register(gRPC *grpc.Server, auth IAuth) {
	sandjmav1.RegisterAuthServer(gRPC, &serverAPI{auth: auth})
}

func (s *serverAPI) Login(ctx context.Context, req *sandjmav1.LoginRequest) (*sandjmav1.LoginResponse, error) {
	if err := validateLogin(req); err != nil {
		return nil, err
	}
	token, err := s.auth.Login(ctx, req.Password, req.Phone, 1)

	if err != nil {
		return nil, err
	}
	return &sandjmav1.LoginResponse{
		UserUuid: token,
	}, nil
}

func (s *serverAPI) Register(ctx context.Context, req *sandjmav1.RegisterRequest) (*sandjmav1.RegisterResponse, error) {
	if err := validateRegister(req); err != nil {
		return nil, err
	}
	userId, err := s.auth.RegisterNewUser(ctx, req.GetPhone(), req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "user is required")
	}
	return &sandjmav1.RegisterResponse{UserUuid: userId}, nil
}

func validateLogin(req *sandjmav1.LoginRequest) error {
	if req.Phone == "" {
		return status.Error(codes.InvalidArgument, "phone is required")
	}
	if req.Password == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}
	return nil
}

func validateRegister(req *sandjmav1.RegisterRequest) error {
	if req.Phone == "" {
		return status.Error(codes.InvalidArgument, "phone is required")
	}
	if req.Password == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}
	return nil
}
