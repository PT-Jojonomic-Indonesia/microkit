package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"bitbucket.org/jojocoders/microkit/response"
	"bitbucket.org/jojocoders/microkit/server"
	"bitbucket.org/jojocoders/microkit/tracer"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
)

func main() {
	godotenv.Load()
	log.SetFlags(log.LstdFlags | log.Llongfile)
	router := GetRoutes()
	url := os.Getenv("JAEGER_ENDPOINT")

	tracer.InitOtel(url, "Sum Service 2", "v1.0.0", "development")
	ctx := context.TODO()
	_, span := tracer.Start(ctx, "Initiate Sum", "", "")
	span.End()

	router = otelhttp.NewHandler(router, "")
	router = server.NewTraceparentHandler(router)

	server.Serve("8181", router)
}

func GetRoutes() http.Handler {
	route := mux.NewRouter()

	route.PathPrefix("/sum").Handler(http.HandlerFunc(SumHandler))
	return route
}

func SumHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Number1 float64 `json:"number1"`
		Number2 float64 `json:"number2"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	hs, _ := json.Marshal(r.Header)
	log.Println(string(hs))

	vals, _ := json.Marshal(r.URL.Query())
	log.Println(string(vals))

	ctx := r.Context()
	traceID := r.URL.Query().Get("trace_id")
	spanID := r.URL.Query().Get("span_id")

	log.Println("start span ", traceID, spanID)
	_, span := tracer.Start(ctx, "Sum", traceID, spanID)
	span.AddEvent("sum event")

	dataSum := Sum(data.Number1, data.Number2)

	payload, _ := json.Marshal(data)
	result := dataSum
	span.SetAttributes(attribute.String("payload", string(payload)))
	span.SetAttributes(attribute.Float64("result", result))

	defer span.End()

	resp := map[string]any{
		"data":  dataSum,
		"error": false,
	}

	response.JSON(w, resp, http.StatusOK)
}

func Sum(number1 float64, number2 float64) float64 {
	return number1 + number2
}
