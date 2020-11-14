package authentication

import (
	"time"

	"github.com/CafeLucuma/go-play/utils/logging"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	AuthenticateUser(email string, pswd string) (*Token, error)
}

type Repository interface {
	GetUser(email string) (*User, error)
}

type service struct {
	uR Repository
}

func NewService(r Repository) *service {
	return &service{uR: r}
}

func (s *service) AuthenticateUser(email, pswd string) (*Token, error) {

	//get user by email
	user, err := s.uR.GetUser(email)
	if err != nil {
		logging.Info.Printf("Cant get user from db: %s", err.Error())
		return nil, err
	}

	//then validate if password entered is the same as the one in DB
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pswd)); err != nil {
		logging.Error.Printf("password doest not match: %s", err.Error())
		return nil, err
	}

	userClaims := NewUserClaim(user.ID, email, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
	})

	token, tErr := generateUserToken(userClaims, []byte("my-secret"))
	if tErr != nil {
		logging.Error.Printf("Error creating user token: %s", tErr.Error())
		return nil, tErr
	}

	return token, nil
}
