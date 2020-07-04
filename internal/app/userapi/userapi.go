package userapi

import (
	"context"

	"github.com/pkg/errors"

	"github.com/evseevbl/userapi/internal/pkg/store"
)

const defaultMinPasswordLength = 8

func NewImplementation(
	storage storage,
	options ...opt,
) *implementation {
	i := &implementation{
		storage:           storage,
		minPasswordLength: defaultMinPasswordLength,
	}

	for _, option := range options {
		option(i)
	}
	return i
}

type implementation struct {
	minPasswordLength int
	storage           storage
}

type opt func(*implementation)

// WithMinPasswordLength changes default password requirement
func WithMinPasswordLength(length int) opt {
	return func(i *implementation) {
		i.minPasswordLength = length
	}
}

type storage interface {
	SaveUser(ctx context.Context, user *store.User) (int64, error)
	GetUserByLogin(ctx context.Context, login string) (*store.User, error)
}

// request and response types
type (
	RegisterRequest struct {
		Login    string `json:"login"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	RegisterResponse struct {
		UserID int64 `json:"user_id"`
	}

	LoginRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		UserID int64  `json:"id"`
		Token  string `json:"token"`
	}
)

var (
	ErrNilRequest       = errors.New("request is nil")
	ErrEmptyField       = errors.New("field cannot be empty")
	ErrPasswordMismatch = errors.New("password does not match")
)
