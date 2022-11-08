package main

import (
	"fmt"
	"net/http"

	"bitbucket.org/jojocoders/microkit/database/postgres"
	"bitbucket.org/jojocoders/microkit/response"
)

var healthHandler = func(w http.ResponseWriter, r *http.Request) {
	if err := postgres.Health(); err != nil {
		response.ErrorJSON(w, fmt.Errorf("postgres : %v", err), http.StatusInternalServerError)
		return
	}
	resData := map[string]interface{}{
		"message": "server is up",
	}
	response.JSON(w, resData, http.StatusOK)
}
