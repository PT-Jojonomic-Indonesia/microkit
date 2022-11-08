package tracer

import (
	"context"
	"net/http"

	trc "go.opentelemetry.io/otel/trace"
)

func StartHTTP(r *http.Request, spanName string) (ctxSpan context.Context, span trc.Span) {
	ctx := r.Context()
	traceID := r.Header.Get("trace-id")
	spanID := r.Header.Get("span-id")

	ctxSpan, span = Start(ctx, spanName, traceID, spanID)
	defer span.End()

	return
}
