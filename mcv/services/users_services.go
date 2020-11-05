package services

import (
	"github.com/CafeLucuma/go-play/mcv/domain"
	"github.com/CafeLucuma/go-play/mcv/utils"
)

func GetUser(userId uint64) (*domain.User, *utils.ApiError) {
	user, apiError := domain.GetUser(userId)

	if apiError != nil {
		return nil, apiError
	}

	return user, nil
}
