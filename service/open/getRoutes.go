package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getRoutes() http.Handler {
	route := mux.NewRouter()

	route.PathPrefix("/open").Handler(http.HandlerFunc(openHandler))
	return route
}
