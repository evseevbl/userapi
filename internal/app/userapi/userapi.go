package userapi

import (
	"context"
	"errors"

	"github.com/evseevbl/userapi/internal/pkg/store"
)

func NewImplementation(
	storage storage,
) *implementation {
	return &implementation{
		storage: storage,
	}
}

type implementation struct {
	storage storage
}

type (
	storage interface {
		SaveUser(ctx context.Context, user *store.User) (int64, error)
		GetUserByLogin(ctx context.Context, login string) (*store.User, error)
	}
)

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
		UserID int64 `json:"id"`
		Token string `json:"token"`
	}
)

var (
	ErrNilRequest       = errors.New("request is nil")
	ErrEmptyField       = errors.New("field cannot be empty")
	ErrPasswordMismatch = errors.New("password does not match")
)
