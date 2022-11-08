package main

import (
	"net/http"

	"github.com/PT-Jojonomic-Indonesia/microkit/database/db2"
	"github.com/PT-Jojonomic-Indonesia/microkit/response"
)

var healthHandler = func(w http.ResponseWriter, r *http.Request) {
	if err := db2.Health(); err != nil {
		response.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	resData := map[string]interface{}{
		"message": "server is up",
	}
	response.JSON(w, resData, http.StatusOK)
}
