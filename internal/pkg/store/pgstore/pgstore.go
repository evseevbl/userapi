// postgres implementation of `store`
package pgstore

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

// New constructor
func New(db *sqlx.DB) *storage {
	return &storage{db: db}
}

type storage struct {
	db *sqlx.DB
}

var (
	// ErrNotFound entity not found
	ErrNotFound = errors.New("not found")
)
