package userapi

import (
	"errors"
	"fmt"
)

func (i *implementation) Register(req *RegisterRequest) (*RegisterResponse, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	fmt.Printf("user %s with email %s\n", req.Login, req.Email)
	return nil, errors.New("not implemented")
}
