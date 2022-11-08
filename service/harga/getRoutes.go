package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getRoutes() http.Handler {
	route := mux.NewRouter()

	route.PathPrefix("/get-harga-emas").Handler(http.HandlerFunc(getHargaEmasHandler))
	return route
}
