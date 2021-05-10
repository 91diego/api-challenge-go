package main

import (
	"log"
	"net/http"
)

func main() {
	// Router create with friendly routes
	router := NewRouter()
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
