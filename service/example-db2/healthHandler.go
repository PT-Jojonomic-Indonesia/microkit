package main

import (
	"net/http"

	"bitbucket.org/jojocoders/microkit/database/db2"
	"bitbucket.org/jojocoders/microkit/response"
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
