package authgrpc

import (
	"context"
	sandjmav1 "github.com/Sanchir01/protos_files_job/pkg/gen/golang/auth"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"

	"google.golang.org/grpc"
)

type Auth interface {
	Login(
		ctx context.Context,
		phone string,
		password string,
	) (string, error)
	RegisterNewUser(ctx context.Context, phone string, email string, password string) (uuid.UUID, error)
	IsAdmin(ctx context.Context, userID uuid.UUID) (bool, error)
}

type serverAuthAPI struct {
	sandjmav1.UnimplementedAuthServer
	auth Auth
}

func Register(gRPC *grpc.Server, auth Auth) {
	sandjmav1.RegisterAuthServer(gRPC, &serverAuthAPI{auth: auth})
}

func (s *serverAuthAPI) Login(ctx context.Context, req *sandjmav1.LoginRequest) (*sandjmav1.LoginResponse, error) {
	if err := validateLogin(req); err != nil {
		return nil, err
	}
	token, err := s.auth.Login(ctx, req.Password, req.Phone)
	slog.Info("token", slog.String("token", token))
	if err != nil {
		return nil, err
	}
	return &sandjmav1.LoginResponse{
		UserUuid: req.Phone,
		TokenTTL: req.Password,
	}, nil
}

func (s *serverAuthAPI) Register(ctx context.Context, req *sandjmav1.RegisterRequest) (*sandjmav1.RegisterResponse, error) {
	if err := validateRegister(req); err != nil {
		return nil, err
	}
	userId, err := s.auth.RegisterNewUser(ctx, req.GetPhone(), req.GetPassword(), req.Email)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "user is required")
	}
	return &sandjmav1.RegisterResponse{UserUuid: userId.String()}, nil
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
