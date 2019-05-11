package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route structure
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is an array of Route object
type Routes []Route

// NewRouter is to build a route
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Root",
		"GET",
		"/",
		RootHandler,
	},
	Route{
		"Status",
		"GET",
		"/status",
		StatusHandler,
	},
}
