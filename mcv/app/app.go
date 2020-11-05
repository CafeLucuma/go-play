package app

import (
	"fmt"
	"net/http"

	"github.com/CafeLucuma/go-play/mcv/controllers"
)

func StartApp() {
	fmt.Println("---- Starting app ----")
	http.HandleFunc("/users", controllers.GetUser)

	http.ListenAndServe(":8080", nil)
}
