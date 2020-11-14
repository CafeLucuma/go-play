package main

import (
	"log"
	"net/http"

	"github.com/CafeLucuma/go-play/plates/pkg/adding"
	"github.com/CafeLucuma/go-play/plates/pkg/http/rest"
	"github.com/CafeLucuma/go-play/plates/pkg/listing"
	"github.com/CafeLucuma/go-play/plates/pkg/storage/postgres"
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
	listing := listing.NewService(storage)

	log.Fatal(http.ListenAndServe(":8081", server.GetHandler(adding, listing)))
}
