package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func() {
		log.Fatal(http.ListenAndServe(":8081", server.GetHandler(adding, listing)))
	}()

	<-c
	fmt.Println("Canceled by the user!")
}
