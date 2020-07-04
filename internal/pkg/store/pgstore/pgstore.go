// postgres implementation of `store`
package pgstore

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *storage {
	return &storage{db: db}
}

type storage struct {
	db *sqlx.DB
}

var (
	ErrNotFound = errors.New("not found")
)
