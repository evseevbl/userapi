package pgstore

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"github.com/evseevbl/userapi/internal/pkg/store"
)

func (s *storage) SaveUser(ctx context.Context, user *store.User) (int64, error) {
	q := `
		insert into "user" (login, password_hash, email, phone)
		values ($1, $2, $3, $4)
		returning id
		`
	var id int64
	err := s.db.GetContext(ctx, &id, q, user.Login, user.PasswordHash, user.Email, user.Phone)
	if err != nil {
		return 0, errors.Wrap(err, "cannot save user")
	}
	return id, nil
}

func (s *storage) GetUserByLogin(ctx context.Context, login string) (*store.User, error) {
	q := `
		select 
			id, 
		       login,
		    password_hash,
		    email, 
		    phone
		from "user" 
		where login = $1`
	ret := make([]*store.User, 0, 1)
	err := s.db.SelectContext(ctx, &ret, q, login)
	switch err {
	case nil:
		// all ok
	case sql.ErrNoRows:
		// not found
		return nil, ErrNotFound
	default:
		// something else
		return nil, errors.Wrap(err, "cannot get user from db")
	}

	if len(ret) == 0 {
		return nil, ErrNotFound
	}
	return ret[0], nil
}
