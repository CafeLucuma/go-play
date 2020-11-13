package authentication

import (
	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	ID    int
	Email string
	jwt.StandardClaims
}

func NewUserClaim(ID int, email string, stdC jwt.StandardClaims) UserClaims {
	return UserClaims{
		ID:             ID,
		Email:          email,
		StandardClaims: stdC,
	}
}
