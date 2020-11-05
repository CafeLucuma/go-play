package domain

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNotFound(t *testing.T) {

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

	userID := uint64(1)

	user, err := GetUser(userID)

	assert.Nil(t, err, "Expected error nil user id %v", userID)
	assert.NotNil(t, user, "Expected user not nil for user id %v", userID)
	assert.Equal(t, userID, user.ID)
}
