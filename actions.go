package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	body := json.NewDecoder(r.Body)
	var user User
	err := body.Decode(&user)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	errDotEnv := godotenv.Load()
	if errDotEnv != nil {
		log.Fatal("Error loading .env file")
	}

	userTest := os.Getenv("USER_TEST")
	userPasswordTest := os.Getenv("USER_PASSWORD_TEST")
	if user.Email == userTest && user.Password == userPasswordTest {
		fmt.Println("Logged in!")
	} else {
		fmt.Println("Unauthorized!")
	}
}

func MeController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Claims info")
}

func GetLInkController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Download file")
}
