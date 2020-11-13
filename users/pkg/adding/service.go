package adding

import (
	"github.com/CafeLucuma/go-play/users/pkg/logging"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	AddUser(User) error
}

type Repository interface {
	AddUser(User) error
}

type service struct {
	uR Repository
}

func NewService(r Repository) *service {
	return &service{uR: r}
}

func (s *service) AddUser(u User) error {

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), 5)
	if err != nil {
		logging.Error.Printf("%s: %s", "Cant generate hashed password for user", err.Error())
		return err
	}

	u.Password = string(hashedPwd)

	if err := s.uR.AddUser(u); err != nil {
		logging.Error.Printf("Error adding user %+v, err: %s", u, err.Error())
		return err
	}

	return nil
}
