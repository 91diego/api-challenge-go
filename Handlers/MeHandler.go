package handlers

import (
	"net/http"

	middleware "github.com/91diego/api-rest-challenge/Middleware"
)

func MeHandler(w http.ResponseWriter, r *http.Request) {
	bearerToken := r.Header.Get("Authorization")
	middleware.ClaimsMiddleware(bearerToken)
}
