package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getRoutes() http.Handler {
	route := mux.NewRouter()

	route.PathPrefix("/create-tabungan").Handler(http.HandlerFunc(createTabunganHandler))
	return route
}
