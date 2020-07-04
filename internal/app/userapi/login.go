package userapi

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/pkg/errors"
)

func (i *implementation) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	// check that all required fields are present
	if err := i.validateLogin(req); err != nil {
		return nil, errors.Wrap(err, "validation")
	}

	user, err := i.storage.GetUserByLogin(ctx, req.Login)
	if err != nil {
		return nil, errors.Wrap(err, "get user")
	}

	if err := i.checkUserPassword(user, req.Password); err != nil {
		return nil, errors.Wrap(err, "password check")
	}

	// token mechanism is yet to be implemented
	return &LoginResponse{UserID: user.ID, Token: fmt.Sprintf("#%d", rand.Uint64())}, nil
}

func (i *implementation) validateLogin(req *LoginRequest) error {
	switch {
	case req == nil:
		return ErrNilRequest
	case req.Password == "":
		return errors.Wrap(ErrEmptyField, "password")
	case req.Login == "":
		return errors.Wrap(ErrEmptyField, "login")
	default:
		return nil
	}
}
