package rest

import (
	"encoding/json"
	"net/http"

	"github.com/CafeLucuma/go-play/users/pkg/adding"
	"github.com/CafeLucuma/go-play/users/pkg/authentication"
	"github.com/CafeLucuma/go-play/users/pkg/logging"
	"github.com/julienschmidt/httprouter"
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
			logging.Error.Printf("Error decoding user: %s", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := aS.AddUser(newUser); err != nil {
			logging.Error.Printf("Error adding user: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode("Added new user"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func handleAuthenticateUser(aS authentication.Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		logging.Info.Printf("Authenticating user")

		var credentials authentication.LoginCredentials
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		if err := decoder.Decode(&credentials); err != nil {
			logging.Error.Println("Error decoding credentials: %s", err.Error())

			r := NewUserRoleError(1, "internal error")
			r.RespondJSON(w, http.StatusInternalServerError)
			return
		}

		if credentials.Email == nil {
			logging.Info.Println("missing field Email")
			r := NewUserRoleError(3, "missing field 'Email'")
			r.RespondJSON(w, http.StatusBadRequest)
			return
		}

		if credentials.Password == nil {
			logging.Info.Println("missing field Password")
			r := NewUserRoleError(3, "missing field 'Password'")
			r.RespondJSON(w, http.StatusBadRequest)
			return
		}

		if *credentials.Email == "" {
			logging.Info.Println("credentials email empty")
			r := NewUserRoleError(3, "field 'Email' empty")
			r.RespondJSON(w, http.StatusBadRequest)
			return
		}

		if *credentials.Password == "" {
			logging.Info.Println("credentials password empty")
			r := NewUserRoleError(3, "field 'Password' empty")
			r.RespondJSON(w, http.StatusBadRequest)
			return
		}

		token, err := aS.AuthenticateUser(*credentials.Email, *credentials.Password)
		if err != nil {
			logging.Error.Printf("Error authenticating user: %s", err.Error())

			e := NewUserRoleError(5, "unauthorized user")
			e.RespondJSON(w, http.StatusUnauthorized)
			return
		}

		resp := NewResponse(http.StatusOK, token, nil)
		resp.RespondJSON(w)
	}
}
