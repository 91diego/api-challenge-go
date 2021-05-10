package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Router create with friendly routes
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/login", Login)
	router.HandleFunc("/me", Me)
	router.HandleFunc("/get-links", GetLInk)
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login")
}

func Me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Claims info")
}

func GetLInk(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Download file")
}
