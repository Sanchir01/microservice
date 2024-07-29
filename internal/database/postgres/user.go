package postgres

import (
	"context"
	"github.com/Sanchir01/microservice/internal/domain/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"log/slog"
)

func (s *StorePostgres) SaveUser(ctx context.Context, phone string, passHash []byte, email string) (id uuid.UUID, err error) {
	const op = "storage.postgres.SaveUser"

	conn, err := s.db.Connx(ctx)
	slog.Warn(op)
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()
	return uuid.New(), err
}

func (s *StorePostgres) GetUserByPhone(ctx context.Context, phone string) (models.User, error) {
	password := "sadadasadad"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Ошибка хэширования пароля: %v", err)
	}
	return models.User{ID: uuid.Nil, Email: "", PassHash: hashedPassword}, nil
}

func (s *StorePostgres) IsAdmin(ctx context.Context, userId uuid.UUID) (bool, error) {
	return true, nil
}
