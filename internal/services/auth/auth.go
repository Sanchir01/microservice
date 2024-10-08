package auth

import (
	"context"
	"errors"
	"fmt"
	errorsUser "github.com/Sanchir01/microservice/internal/database/errors"
	"github.com/Sanchir01/microservice/internal/domain/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

type Auth struct {
	log         *slog.Logger
	usrSaver    UserSaver
	usrProvider UserProvider
}

type UserSaver interface {
	SaveUser(
		ctx context.Context, phone string, passHash []byte, email string,
	) (id uuid.UUID, err error)
}

type UserProvider interface {
	GetUserByPhone(ctx context.Context, phone string) (models.User, error)
	IsAdmin(ctx context.Context, userId uuid.UUID) (bool, error)
}

//type AppProvider interface {
//	App(ctx context.Context, appId string) (models.App, error)
//}

func New(
	log *slog.Logger,
	userSaver UserSaver,
	userProvider UserProvider,
) *Auth {
	return &Auth{
		log:         log,
		usrSaver:    userSaver,
		usrProvider: userProvider,
	}
}

func (a *Auth) Login(ctx context.Context, phone string, password string) (string, error) {
	const op = "auth.Login"
	log := a.log.With(
		slog.String("op", op),
		slog.String("phone", phone),
		slog.String("password", password),
	)

	log.Info("attempting to login")

	user, err := a.usrProvider.GetUserByPhone(ctx, phone)

	if err != nil {
		if errors.Is(err, errorsUser.ErrAppNotFound) {
			a.log.With("user not found")
			return "", fmt.Errorf("%s: %w", op, err)
		}
	}
	return user.Email, nil
}

func (a *Auth) RegisterNewUser(ctx context.Context, phone string, email string, password string) (uuid.UUID, error) {
	const op = "auth.RegisterNewUser"
	log := a.log.With(
		slog.String("op", op),
		slog.String("phone", email),
	)
	log.Info("register new user")

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to hash password")
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}
	id, err := a.usrSaver.SaveUser(ctx, phone, passwordHash, email)
	log.Info("user registered")
	return id, nil
}

func (a *Auth) IsAdmin(ctx context.Context, userId uuid.UUID) (bool, error) {
	return true, nil
}
