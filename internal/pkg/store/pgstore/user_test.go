// +build integration

package pgstore

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"github.com/pkg/errors"
	"github.com/pressly/goose"

	"github.com/evseevbl/userapi/internal/pkg/store"
)

const migrationDir = "./migrations"

func TestS_SaveAndGetUser(t *testing.T) {
	// a simple test for the basic scenario (save, then get)

	dsn := os.Getenv("TESTING_DSN")
	if dsn == "" {
		t.Fatal("TESTING_DSN must be set for integration tests")
	}

	a := assert.New(t)

	// connect to database and make storage
	db, err := sqlx.Connect("postgres", dsn)
	a.NoError(err, "database conn")
	defer db.Close()
	s := New(db)

	testCases := []struct {
		name         string
		prepareUsers []*store.User

		errCheck func(error, ...interface{}) bool
		user     *store.User
	}{
		{
			name:         "just one user",
			prepareUsers: nil,
			errCheck:     a.NoError,
			user: &store.User{
				Login:        "test",
				Email:        "test@test",
				PasswordHash: "asdfasdf",
				Phone:        "999",
			},
		},
		{
			name: "login already exists",
			prepareUsers: []*store.User{
				{
					Login:        "test",
					Email:        "test1@test",
					PasswordHash: "111",
					Phone:        "1",
				},
			},
			errCheck: a.Error, // expect an error
			user: &store.User{
				Login:        "test",
				Email:        "test2@test",
				PasswordHash: "222",
				Phone:        "2",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			// prepare db and data if any
			a.NoError(setup(db.DB, migrationDir), "setup")
			for _, u := range tc.prepareUsers {
				a.NoError(addUser(db.DB, u), "add user")
			}

			err := func() error {
				var id int64
				id, err = s.SaveUser(ctx, tc.user)
				if err != nil {
					return errors.Wrap(err, "cannot save")
				}
				tc.user.ID = id // for equality comparison

				userFromDB, err := s.GetUserByLogin(ctx, tc.user.Login)
				if err != nil {
					return errors.Wrap(err, "get by login")
				}

				a.Equal(tc.user, userFromDB)
				return nil
			}()

			tc.errCheck(err)
		})
	}
}

func addUser(db *sql.DB, u *store.User) error {
	q := `insert into "user" (login, email, password_hash, phone) values ($1, $2, $3, $3)`
	if _, err := db.Exec(q, u.Login, u.Email, u.PasswordHash); err != nil {
		return errors.Wrap(err, "exec")
	}
	return nil
}

func setup(db *sql.DB, dir string) error {
	// down migrations so all tables get deleted
	if err := goose.DownTo(db, dir, 0); err != nil {
		return errors.Wrap(err, "down migrations")
	}
	// recreate schema
	if err := goose.Up(db, dir); err != nil {
		return errors.Wrap(err, "down migrations")
	}
	return nil
}
