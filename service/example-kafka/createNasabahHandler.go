package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PT-Jojonomic-Indonesia/microkit/response"
	"github.com/PT-Jojonomic-Indonesia/microkit/service/example-kafka/entity"
	"github.com/PT-Jojonomic-Indonesia/microkit/validator"

	"github.com/PT-Jojonomic-Indonesia/microkit/tracer"
)

var createNasabahHandler = func(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	traceID := r.Header.Get("trace-id")
	spanID := r.Header.Get("span-id")
	log.Println(traceID, spanID)

	ctx2, span := tracer.Start(ctx, "createNasabah()", traceID, spanID)
	defer span.End()

	nasabah := &entity.Nasabah{}
	if err := json.NewDecoder(r.Body).Decode(&nasabah); err != nil {
		response.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := validator.Validate(ctx2, nasabah); err != nil {
		response.ErrorJSON(w, validator.GetErrors(err), http.StatusBadRequest)
		return
	}

	if err := PublishNasabah(nasabah); err != nil {
		response.ErrorJSON(w, fmt.Errorf("error publish nasabah : %s", err), http.StatusInternalServerError)
		return
	}

	resData := map[string]interface{}{
		"message": "success create nasabah",
		"data":    nasabah,
	}
	response.JSON(w, resData, http.StatusOK)
}
