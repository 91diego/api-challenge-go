package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	middleware "github.com/91diego/api-rest-challenge/Middleware"
	models "github.com/91diego/api-rest-challenge/Models"
)

func MeHandler(w http.ResponseWriter, r *http.Request) {
	bearerToken := r.Header.Get("Authorization")
	claims := middleware.ClaimsMiddleware(bearerToken)
	fmt.Println(claims)
	claimsResponse := ""
	for k, v := range claims {
		key := fmt.Sprintf("%v", k)
		value := fmt.Sprintf("%v", v)
		claimsResponse = string(key) + ": " + string(value)
	}
	response := models.Response{
		Code:    http.StatusAccepted,
		Status:  "success",
		Message: "User claims!",
		Data:    claimsResponse,
	}
	json.NewEncoder(w).Encode(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}
