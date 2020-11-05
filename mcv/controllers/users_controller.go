package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CafeLucuma/go-play/mcv/services"
	"github.com/CafeLucuma/go-play/mcv/utils"
)

func GetUser(resp http.ResponseWriter, request *http.Request) {
	userID, err := strconv.ParseUint(request.URL.Query().Get("user_id"), 10, 64)
	if err != nil {

		apiError := &utils.ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    "user_id must be an integer",
		}

		errJSON, _ := json.Marshal(apiError)
		resp.WriteHeader(apiError.StatusCode)
		resp.Write(errJSON)
		return
	}

	user, apiError := services.GetUser(userID)
	if apiError != nil {

		errJSON, _ := json.Marshal(apiError)

		resp.WriteHeader(apiError.StatusCode)
		resp.Write(errJSON)
		return
	}

	userJSON, _ := json.Marshal(user)
	resp.Write(userJSON)
}
