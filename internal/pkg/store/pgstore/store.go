package pgstore

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *s {
	return &s{db: db}
}

type s struct {
	db *sqlx.DB
}

var (
	ErrNotFound = errors.New("not found")
)
