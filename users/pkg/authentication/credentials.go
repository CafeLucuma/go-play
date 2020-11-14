package authentication

import (
	"errors"
)

type LoginCredentials struct {
	Password *string `json:"password"`
	Email    *string `json:"email"`
}

func (lc LoginCredentials) Validate() error {
	if lc.Email == nil {
		return errors.New("missing field 'Email'")
	}

	if lc.Password == nil {
		return errors.New("missing field 'Password'")
	}

	if *lc.Email == "" {
		return errors.New("field 'Email' empty")
	}

	if *lc.Password == "" {
		return errors.New("field 'Password' empty")
	}

	return nil
}
