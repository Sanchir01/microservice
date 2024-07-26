package postgres

import "github.com/jmoiron/sqlx"

type StorePostgres struct {
	db *sqlx.DB
}

func NewStorePostgres(db *sqlx.DB) *StorePostgres {
	return &StorePostgres{db: db}
}
