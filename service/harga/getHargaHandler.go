package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PT-Jojonomic-Indonesia/microkit/response"
	"github.com/PT-Jojonomic-Indonesia/microkit/tracer"

	"go.opentelemetry.io/otel/attribute"
)

func getHargaEmasHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	traceID := r.Header.Get("trace-id")
	spanID := r.Header.Get("span-id")
	log.Println(traceID, spanID)

	ctx2, span := tracer.Start(ctx, "getHargaEmas()", traceID, spanID)
	defer span.End()

	dataResponse := GetHargaEmas(ctx2)

	result, _ := json.Marshal(dataResponse)
	span.SetAttributes(attribute.String("result", string(result)))

	resp := map[string]any{
		"data":  dataResponse,
		"error": false,
	}

	response.JSON(w, resp, http.StatusOK)
}
