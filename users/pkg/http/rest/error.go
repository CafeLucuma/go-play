package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UserRoleError struct {
	Code    int    `json:"code,ommitempty"`
	Message string `json:"message"`
}

func NewUserRoleError(code int, message string) UserRoleError {
	return UserRoleError{
		Code:    code,
		Message: message,
	}
}

func (re *UserRoleError) RespondJSON(w http.ResponseWriter, statusCode int) {

	e := new(bytes.Buffer)
	decoder := json.NewEncoder(e)

	log.Printf("Error: %+v", re)

	if err := decoder.Encode(re); err != nil {
		log.Println("Error decoding error: ", err.Error())
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, e)
}
