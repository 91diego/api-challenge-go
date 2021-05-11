package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	models "github.com/91diego/api-rest-challenge/Models"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	body := json.NewDecoder(r.Body)
	var user models.User
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
		response := models.Response{
			Code:    http.StatusAccepted,
			Status:  "success",
			Message: "Logged in successfully!",
			Data:    result,
		}
		json.NewEncoder(w).Encode(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
	} else {
		response := models.Response{
			Code:    http.StatusUnauthorized,
			Status:  "error",
			Message: "Unauthorized!",
			Data:    "",
		}
		json.NewEncoder(w).Encode(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
	}
}
