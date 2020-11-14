package rest

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/CafeLucuma/go-play/users/pkg/adding"
	"github.com/CafeLucuma/go-play/users/pkg/authentication"
	response "github.com/CafeLucuma/go-play/utils/http-response"
	"github.com/CafeLucuma/go-play/utils/logging"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

type server struct {
	router *httprouter.Router
}

func NewServer() server {
	s := server{
		router: httprouter.New(),
	}

	return s
}

func (s *server) GetHandler(addS adding.Service, authS authentication.Service) http.Handler {

	s.router.POST("/auth/user", handleAuthenticateUser(authS))
	s.router.POST("/user", handleAddUser(addS))

	return s.router
}

func handleAddUser(aS adding.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		logging.Info.Printf("Adding new user")

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		var newUser adding.User
		if err := decoder.Decode(&newUser); err != nil {
			logging.Error.Printf("Error decoding user: %s", err)
			r := response.New(http.StatusBadRequest, response.NewGenericBody(2, "bad json"), nil)
			r.RespondJSON(w)
			return
		}

		if err := newUser.Validate(); err != nil {
			logging.Error.Printf("Invalid user credentials: %s", err)
			r := response.New(http.StatusBadRequest, response.NewGenericBody(1, err.Error()), nil)
			r.RespondJSON(w)
			return
		}

		if err := aS.AddUser(newUser); err != nil {
			logging.Error.Printf("Error adding user: %s", err)
			r := response.New(http.StatusInternalServerError, response.NewGenericBody(1, "internal error"), nil)
			r.RespondJSON(w)
			return
		}

		resp := response.New(http.StatusCreated, response.NewGenericBody(http.StatusCreated, "user added"), nil)
		resp.RespondJSON(w)
	}
}

func handleAuthenticateUser(aS authentication.Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		logging.Info.Printf("Authenticating user")

		var credentials authentication.LoginCredentials
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		if err := decoder.Decode(&credentials); err != nil {
			logging.Error.Printf("Error decoding credentials: %s", err.Error())

			r := response.New(http.StatusInternalServerError, response.NewGenericBody(1, "internal error"), nil)
			r.RespondJSON(w)
			return
		}

		if err := credentials.Validate(); err != nil {
			logging.Error.Printf("error validating credentials: %s", err)
			r := response.New(http.StatusBadRequest, response.NewGenericBody(3, err.Error()), nil)
			r.RespondJSON(w)
			return
		}

		token, err := aS.AuthenticateUser(*credentials.Email, *credentials.Password)
		if err != nil {
			logging.Error.Printf("Error authenticating user: %s", err)

			switch err {
			case sql.ErrNoRows:
				e := response.New(http.StatusUnauthorized, response.NewGenericBody(5, "user not found"), nil)
				e.RespondJSON(w)
			case bcrypt.ErrMismatchedHashAndPassword:
				e := response.New(http.StatusUnauthorized, response.NewGenericBody(5, "incorrect password"), nil)
				e.RespondJSON(w)
			default:
				e := response.New(http.StatusUnauthorized, response.NewGenericBody(5, "unauthorized user"), nil)
				e.RespondJSON(w)
			}
			return
		}

		resp := response.New(http.StatusOK, token, nil)
		resp.RespondJSON(w)
	}
}
