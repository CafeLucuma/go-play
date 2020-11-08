package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/CafeLucuma/go-play/domain-hex/pkg/adding"
	"github.com/CafeLucuma/go-play/domain-hex/pkg/listing"
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

	s.router.POST("/plates", handleAddPlate(a))
	s.router.GET("/plates/:id", handleGetPlate(l))

	return s.router
}

func handleAddPlate(aS adding.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		var newPlate adding.Plate
		if err := decoder.Decode(&newPlate); err != nil {
			log.Printf("Error decoding plate: %s", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		aS.AddPlate(newPlate)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode("Added new plate"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func handleGetPlate(aS listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		ID, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		plate, err := aS.GetPlate(ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(plate); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(plate)
	}
}
