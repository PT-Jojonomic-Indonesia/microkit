package main

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/jojocoders/microkit/response"
	"bitbucket.org/jojocoders/microkit/tracer"

	"go.opentelemetry.io/otel/attribute"
)

func openHandler(w http.ResponseWriter, r *http.Request) {
	var input OpenInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		response.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	ctx2, span := tracer.StartHTTP(r, "Open()")

	dataResponse := Open(ctx2, input)

	payload, _ := json.Marshal(input)
	result, _ := json.Marshal(dataResponse)
	span.SetAttributes(attribute.String("payload", string(payload)))
	span.SetAttributes(attribute.String("result", string(result)))

	defer span.End()

	resp := map[string]any{
		"data":  dataResponse,
		"error": false,
	}

	response.JSON(w, resp, http.StatusOK)
}
