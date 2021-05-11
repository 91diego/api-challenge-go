package routes

import (
	"net/http"

	handlers "github.com/91diego/api-rest-challenge/Handlers"
	middleware "github.com/91diego/api-rest-challenge/Middleware"

	"github.com/gorilla/mux"
)

// Routes definitiion
type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		if route.Name == "Login" {
			router.Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(route.HandleFunc)
		} else {
			router.Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(middleware.JWTmiddleware(route.HandleFunc))
		}
	}

	return router
}

var routes = Routes{
	Route{
		"Login",
		"POST",
		"/login",
		handlers.LoginHandler,
	},
	Route{
		"me",
		"GET",
		"/me",
		handlers.MeHandler,
	},
	Route{
		"get-links",
		"POST",
		"/get-links",
		handlers.GetLInkHandler,
	},
}
