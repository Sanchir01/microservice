package auth

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
)

type Auth struct {
	log *slog.Logger
}

type UserSaver interface {
	SaveUser(ctx context.Context, phone string, passHash []byte) (id uuid.UUID, err error)
}
