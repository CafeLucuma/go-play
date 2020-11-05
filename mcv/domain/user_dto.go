package domain

import (
	"fmt"
	"net/http"

	"github.com/CafeLucuma/go-play/mcv/utils"
)

var (
	users = map[uint64]*User{
		1: {ID: 1,
			FirstName: "Oscar",
			LastName:  "Carrasco",
			Email:     "oscar@gmail.com",
		},
	}
)

func GetUser(userID uint64) (*User, *utils.ApiError) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	apiError := &utils.ApiError{
		StatusCode: http.StatusNotFound,
		Message:    fmt.Sprintf("user id %v not found", userID),
	}

	return nil, apiError
}
