package userapi

import (
	"context"

	"github.com/pkg/errors"

	"github.com/evseevbl/userapi/internal/pkg/store"
)

func (i *implementation) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	// check that all required fields are present
	if err := i.validateRegister(req); err != nil {
		return nil, errors.Wrap(err, "validation")
	}

	hash, err := i.generatePasswordHash(req.Password)
	if err != nil {
		return nil, errors.Wrap(err, "generate hash")
	}

	id, err := i.storage.SaveUser(ctx, &store.User{
		Login:        req.Login,
		Email:        req.Email,
		PasswordHash: string(hash),
	})
	if err != nil {
		return nil, errors.Wrap(err, "save user")
	}

	return &RegisterResponse{UserID: id}, nil
}

func (i *implementation) validateRegister(req *RegisterRequest) error {
	switch {
	case req == nil:
		return ErrNilRequest
	case req.Login == "":
		return errors.Wrap(ErrEmptyField, "login")
	case req.Password == "":
		return errors.Wrap(ErrEmptyField, "password")
	case req.Email == "":
		return errors.Wrap(ErrEmptyField, "email")
	case req.Phone == "":
		return errors.Wrap(ErrEmptyField, "phone")
	default:
		return nil
	}
}
