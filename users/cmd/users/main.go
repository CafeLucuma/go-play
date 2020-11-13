package main

import (
	"log"
	"net/http"

	"github.com/CafeLucuma/go-play/users/pkg/adding"
	"github.com/CafeLucuma/go-play/users/pkg/authentication"
	"github.com/CafeLucuma/go-play/users/pkg/http/rest"
	"github.com/CafeLucuma/go-play/users/pkg/storage/postgres"
	"github.com/joho/godotenv"
)

func main() {
	//loading environment variables
	godotenv.Load()

	server := rest.NewServer()

	storage, err := postgres.NewStorage()
	if err != nil {
		panic(err)
	}
	defer storage.CloseDB()

	adding := adding.NewService(storage)
	authentication := authentication.NewService(storage)

	log.Fatal(http.ListenAndServe(":8080", server.GetHandler(adding, authentication)))
}
