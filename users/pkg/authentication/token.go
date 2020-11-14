package authentication

import (
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Token string `json:"token"`
}

func generateUserToken(c UserClaims, secret []byte) (*Token, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	signedToken, err := token.SignedString(secret)
	if err != nil {
		return nil, err
	}

	return &Token{
		Token: signedToken,
	}, nil
}
