package main

import (
	"net/http"
	"os"

	"bitbucket.org/jojocoders/microkit/kafka"
	"bitbucket.org/jojocoders/microkit/response"
)

var healthHandler = func(w http.ResponseWriter, r *http.Request) {
	if err := kafka.Health(os.Getenv("KAFKA_ENDPOINT")); err != nil {
		response.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	resData := map[string]interface{}{
		"message": "server is up",
	}
	response.JSON(w, resData, http.StatusOK)
}
