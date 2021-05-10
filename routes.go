package main

import (
	"net/http"

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
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Login",
		"POST",
		"/login",
		LoginController,
	},
	Route{
		"me",
		"GET",
		"/me",
		MeController,
	},
	Route{
		"get-links",
		"GET",
		"/get-links",
		GetLInkController,
	},
}
