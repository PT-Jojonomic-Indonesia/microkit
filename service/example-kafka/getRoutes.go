package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getRoutes() http.Handler {
	route := mux.NewRouter()

	route.Path("/health").Methods(http.MethodGet).HandlerFunc(healthHandler)
	route.Path("/nasabah").Methods(http.MethodPost).HandlerFunc(createNasabahHandler)
	return route
}
