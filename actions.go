package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
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
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email":       user.Email,
			"secret_word": "nomada",
		})
		mySecret := []byte(os.Getenv("JWT_TOKEN"))
		result, err := token.SignedString(mySecret)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(result)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
	} else {
		json.NewEncoder(w).Encode("Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func MeController(w http.ResponseWriter, r *http.Request) {
	bearerToken := r.Header.Get("Authorization")
	extractClaims(bearerToken)
}

func GetLInkController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Download file")
}
