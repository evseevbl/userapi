package userapi

import (
	"errors"
)

func NewImplementation() *implementation {
	return &implementation{}
}

type implementation struct {
}

type (
	RegisterRequest struct {
		Login    string `json:"login"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	RegisterResponse struct {
	}

	LoginRequest struct {
	}

	LoginResponse struct {
	}

	CheckRequest struct {
	}

	CheckResponse struct {
	}
)

var (
	ErrNilRequest = errors.New("request is nil")
)
