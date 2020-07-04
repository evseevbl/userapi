package userapi

import (
	"context"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/evseevbl/userapi/internal/pkg/store"
)

func (i *implementation) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {

	if err := i.validateRegister(req); err != nil {
		return nil, errors.Wrap(err, "validation")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
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
		return errors.Wrap(ErrFieldRequired, "login")
	case req.Email == "":
		return errors.Wrap(ErrFieldRequired, "email")
	case req.Phone == "":
		return errors.Wrap(ErrFieldRequired, "phone")
	default:
		return nil
	}
}
