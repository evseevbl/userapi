package userapi

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/evseevbl/userapi/internal/pkg/store"
)

func TestImplementation_Register(t *testing.T) {
	a := assert.New(t)

	var someUserID int64 = 123

	testCases := []struct {
		name    string
		req     *RegisterRequest
		expResp *RegisterResponse
		expErr  error
	}{
		{
			name: "all ok",
			req: &RegisterRequest{
				Login:    "test",
				Email:    "test@test",
				Phone:    "999",
				Password: "qwerty",
			},
			expResp: &RegisterResponse{UserID: someUserID},
			expErr:  nil,
		},
		{
			name: "empty login",
			req: &RegisterRequest{
				Login:    "",
				Email:    "test@test",
				Phone:    "999",
				Password: "qwerty",
			},
			expResp: nil,
			expErr:  ErrEmptyField,
		},
		// add testCases for duplicate login
	}

	s := NewStorageMock(t)
	impl := NewImplementation(s, WithMinPasswordLength(1))

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			if tc.expResp != nil {
				s.SaveUserMock.Set(func(ctx context.Context, user *store.User) (i1 int64, err error) {
					a.True(user.PasswordHash != "")
					return someUserID, nil
				})
			}

			ret, err := impl.Register(ctx, tc.req)
			a.Equal(tc.expErr, errors.Cause(err))
			a.Equal(tc.expResp, ret)
		})
	}
}
