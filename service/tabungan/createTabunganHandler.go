package main

import (
	"context"
	"encoding/json"
	"net/http"

	"bitbucket.org/jojocoders/microkit/response"
	"bitbucket.org/jojocoders/microkit/tracer"

	"go.opentelemetry.io/otel/attribute"
)

func createTabunganHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	traceID := r.Header.Get("trace-id")
	spanID := r.Header.Get("span-id")

	ctx2, span := tracer.Start(ctx, "createTabungan()", traceID, spanID)
	defer span.End()

	dataResponse := CreateTabungan(ctx2)

	result, _ := json.Marshal(dataResponse)
	span.SetAttributes(attribute.String("result", string(result)))

	resp := map[string]any{
		"data":  dataResponse,
		"error": false,
	}

	response.JSON(w, resp, http.StatusOK)
}

func CreateTabungan(ctx context.Context) (resp any) {
	return
}
