package main

import (
	"log"
	"net/http"

	"github.com/CafeLucuma/go-play/domain-hex/pkg/adding"
	"github.com/CafeLucuma/go-play/domain-hex/pkg/http/rest"
	"github.com/CafeLucuma/go-play/domain-hex/pkg/listing"
	"github.com/CafeLucuma/go-play/domain-hex/pkg/storage/json"
)

func main() {
	server := rest.NewServer()

	storage, _ := json.NewStorage()

	adding := adding.NewService(storage)
	listing := listing.NewService(storage)

	log.Fatal(http.ListenAndServe(":8080", server.GetHandler(adding, listing)))
}
