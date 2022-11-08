package server

import (
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type TraceparentHandler struct {
	next  http.Handler
	props propagation.TextMapPropagator
}

func NewTraceparentHandler(next http.Handler) *TraceparentHandler {
	return &TraceparentHandler{
		next:  next,
		props: otel.GetTextMapPropagator(),
	}
}

func (h *TraceparentHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.props.Inject(req.Context(), propagation.HeaderCarrier(w.Header()))
	h.next.ServeHTTP(w, req)
}
