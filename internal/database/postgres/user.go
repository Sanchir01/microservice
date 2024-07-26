package postgres

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
)

func (s *StorePostgres) SaveUser(ctx context.Context) (uuid.UUID, error) {
	const op = "storage.postgres.SaveUser"

	conn, err := s.db.Connx(ctx)
	slog.Warn(op)
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()
	return uuid.New(), err
}

func (s *StorePostgres) GetUserByPhone(ctx context.Context, username string) (uuid.UUID, error) {
	return uuid.Nil, nil
}

func (s *StorePostgres) IsAdmin(ctx context.Context, username string) (bool, error) {
	return true, nil
}
