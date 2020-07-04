package userapi

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/evseevbl/userapi/internal/pkg/store"
)

func (i *implementation) checkUserPassword(user *store.User, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	switch err {
	case nil:
		return nil
	case bcrypt.ErrMismatchedHashAndPassword:
		return ErrPasswordMismatch
	default:
		return errors.Wrap(err, "bcrypt error")
	}
}

func (i *implementation) generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Wrap(err, "hash from password")
	}
	return string(hash), nil
}
