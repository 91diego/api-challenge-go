package main

import (
	"log"
	"net/http"

	routing "github.com/91diego/api-rest-challenge/Routes"
)

func main() {
	// Router create with friendly routes
	router := routing.NewRouter()
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
