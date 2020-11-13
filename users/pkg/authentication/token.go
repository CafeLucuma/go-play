package authentication

import (
	"github.com/CafeLucuma/go-play/users/pkg/logging"
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Token string `json:"token"`
}

func generateUserToken(c UserClaims) (*Token, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	mySecret := "my-secret"

	signedToken, err := token.SignedString([]byte(mySecret))
	if err != nil {
		logging.Error.Printf("Cant generate jwt token: %s", err.Error())
		return nil, err
	}

	logging.Info.Println("Generated jwt token")

	return &Token{
		Token: signedToken,
	}, nil
}
