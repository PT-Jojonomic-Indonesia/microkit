package tracer

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	trc "go.opentelemetry.io/otel/trace"
)

var otelTracer trc.Tracer

func InitOtel(url string, serviceName, version, environment string) {

	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		log.Println(err)
		return
	}

	f, _ := os.Create("trace.txt")
	exp2, _ := newExporter(f)

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithBatcher(exp2),
		trace.WithResource(newResource(serviceName, version, environment)),
	)

	otel.SetTracerProvider(tp)

	otelTracer = otel.Tracer("service")
}

func newResource(serviceName, version, environment string) *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
			semconv.ServiceVersionKey.String(version),
			attribute.String("environment", environment),
		),
	)
	return r
}
func newExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
}

func Start(ctx context.Context, spanName string, traceID, spanID string) (ctxSpan context.Context, span trc.Span) {
	log.Println(traceID)
	if traceID != "" {
		var spanContextConfig trc.SpanContextConfig
		spanContextConfig.TraceID, _ = trc.TraceIDFromHex(traceID)
		spanContextConfig.SpanID, _ = trc.SpanIDFromHex(spanID)
		spanContextConfig.TraceFlags = 01
		spanContextConfig.Remote = true

		spanContext := trc.NewSpanContext(spanContextConfig)

		fmt.Println("IS VALID? ", spanContext.IsValid()) // check if okay

		ctx = trc.ContextWithSpanContext(ctx, spanContext)
		ctxSpan, span = otelTracer.Start(ctx, spanName)
		return

	}
	ctxSpan, span = otelTracer.Start(ctx, spanName)

	return
}
