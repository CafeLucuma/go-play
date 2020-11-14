package rest

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/CafeLucuma/go-play/plates/pkg/adding"
	"github.com/CafeLucuma/go-play/plates/pkg/listing"
	response "github.com/CafeLucuma/go-play/utils/http-response"
	"github.com/CafeLucuma/go-play/utils/logging"
	"github.com/dgrijalva/jwt-go"
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

func (s *server) GetHandler(a adding.Service, l listing.Service) http.Handler {

	s.router.POST("/plates", isAuthorized(handleAddPlate(a)))
	s.router.GET("/plates/:id", handleGetPlate(l))
	s.router.GET("/plates", handleGetPlates(l))

	return s.router
}

func isAuthorized(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			resp := response.New(http.StatusBadRequest, response.NewGenericBody(6, "missing authorization header"), nil)
			resp.RespondJSON(w)
			return
		}

		s := strings.Split(authHeader, " ")
		if len(s) != 2 {
			resp := response.New(http.StatusBadRequest, response.NewGenericBody(6, "invalid token format"), nil)
			resp.RespondJSON(w)
			return
		}

		receivedToken := s[1]

		token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte("my-secret"), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Printf("%+v", claims)
			next(w, r, p)
		} else {
			resp := response.New(http.StatusUnauthorized, response.NewGenericBody(6, "invalid token: "+err.Error()), nil)
			resp.RespondJSON(w)
		}
	}
}

func handleAddPlate(aS adding.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		var newPlate adding.Plate
		if err := decoder.Decode(&newPlate); err != nil {
			logging.Error.Printf("Error decoding plate: %s", err.Error())
			resp := response.New(http.StatusBadRequest, response.NewGenericBody(6, "internal error"), nil)
			resp.RespondJSON(w)
			return
		}

		if err := newPlate.Validate(); err != nil {
			resp := response.New(http.StatusBadRequest, response.NewGenericBody(6, err.Error()), nil)
			resp.RespondJSON(w)
			return
		}

		if err := aS.AddPlate(newPlate); err != nil {
			resp := response.New(
				http.StatusInternalServerError,
				response.NewGenericBody(6, "internal error"),
				nil)
			resp.RespondJSON(w)
			return
		}

		resp := response.New(http.StatusCreated, response.NewGenericBody(6, "plate added"), nil)
		resp.RespondJSON(w)
	}
}

func handleGetPlate(aS listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		ID, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			logging.Error.Printf("invalid id in request: %s", err.Error())
			resp := response.New(http.StatusBadRequest, response.NewGenericBody(6, "invalid id"), nil)
			resp.RespondJSON(w)
			return
		}

		plate, err := aS.GetPlate(ID)
		if err != nil {

			if err == sql.ErrNoRows {
				logging.Error.Printf("user id not found: %v", ID)
				resp := response.New(http.StatusBadRequest, response.NewGenericBody(6, "user id not found"), nil)
				resp.RespondJSON(w)
			} else {
				logging.Error.Printf("error getting user: %s", err)
				resp := response.New(http.StatusInternalServerError, response.NewGenericBody(6, "internal error"), nil)
				resp.RespondJSON(w)
			}

			return
		}

		resp := response.New(http.StatusOK, plate, nil)
		resp.RespondJSON(w)
	}
}

func handleGetPlates(aS listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		plates, err := aS.GetPlates()
		if err != nil {
			logging.Error.Printf("error obtaining plates: %s", err)
			resp := response.New(
				http.StatusInternalServerError,
				response.NewGenericBody(6, "internal error"),
				nil)
			resp.RespondJSON(w)
			return
		}

		resp := response.New(http.StatusBadRequest, plates, nil)
		resp.RespondJSON(w)
	}
}
