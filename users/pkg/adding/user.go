package adding

import "errors"

type User struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
}

func (u User) Validate() error {
	if u.Name == "" {
		return errors.New("field 'name' empty")
	}
	if u.LastName == "" {
		return errors.New("field 'last_name' empty")
	}
	if u.Password == "" {
		return errors.New("field 'password' empty")
	}
	if u.Email == "" {
		return errors.New("field 'email' empty")
	}

	return nil
}
