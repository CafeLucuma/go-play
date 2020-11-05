package services

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/CafeLucuma/go-play/mcv/domain"
	"github.com/CafeLucuma/go-play/mcv/utils"
	"github.com/stretchr/testify/assert"
)

var (
	getUserFunc func(uint64) (*domain.User, *utils.ApiError)
)

func init() {
	domain.UserDao = usersDaoMock{}
}

type usersDaoMock struct{}

func (u usersDaoMock) GetUser(userID uint64) (*domain.User, *utils.ApiError) {
	return getUserFunc(userID)
}

func TestGetUserNotFound(t *testing.T) {

	getUserFunc = func(userID uint64) (*domain.User, *utils.ApiError) {
		return nil, &utils.ApiError{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("user id %v not found", userID),
		}
	}

	userID := uint64(0)
	expectedMsg := fmt.Sprintf("user id %v not found", userID)
	expectedStatusCode := http.StatusNotFound

	user, err := GetUser(userID)

	assert.NotNil(t, err, "Expected error for user id %v", userID)
	assert.Nil(t, user, "Expected user nil for user id %v", userID)
	assert.Equal(t, expectedStatusCode, err.StatusCode, "Expected status code %v but got %v", expectedStatusCode, err.StatusCode)
	assert.Equal(t, expectedMsg, err.Message, "Expected status code %v but got %v", expectedMsg, err.Message)
}

func TestGetUserFound(t *testing.T) {

	getUserFunc = func(userID uint64) (*domain.User, *utils.ApiError) {
		user := domain.User{
			ID:        1,
			FirstName: "Pepe",
			LastName:  "The frog",
			Email:     "pf@gmail.com",
		}

		return &user, nil
	}

	userID := uint64(1)

	user, err := GetUser(userID)

	assert.Nil(t, err, "Expected error nil user id %v", userID)
	assert.NotNil(t, user, "Expected user not nil for user id %v", userID)
	assert.Equal(t, userID, user.ID)
	assert.Equal(t, "Pepe", user.FirstName)
	assert.Equal(t, "The frog", user.LastName)
	assert.Equal(t, "pf@gmail.com", user.Email)
}
