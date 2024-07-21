package models

import "github.com/google/uuid"

type USer struct {
	ID       uuid.UUID
	Email    string
	PassHash []byte
}
