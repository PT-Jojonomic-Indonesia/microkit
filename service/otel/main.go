package main

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/PT-Jojonomic-Indonesia/microkit/sentry"
	"github.com/PT-Jojonomic-Indonesia/microkit/tracer"
	"github.com/joho/godotenv"
	"go.opentelemetry.io/otel/attribute"
)

func main() {
	godotenv.Load()
	os.Setenv("SERVICE_NAME", "observability")
	os.Setenv("SERVICE_VERSION", "1.0.0")

	// observability.InitDefault()
	// tracer.InitOtel(os.Getenv("JAEGER_URL"), os.Getenv("SERVICE_NAME"), os.Getenv("SERVICE_VERSION"), os.Getenv("ENVIRONMENT"))
	// cfg := sentry.Config{
	// 	Dsn:         os.Getenv("SENTRY_DSN"),
	// 	Release:     os.Getenv("SERVICE_VERSION"),
	// 	Environment: os.Getenv("ENVIRONMENT"),
	// }
	// sentry.Init(cfg)
	// defer sentry.Close()

	tracer.InitOtel(os.Getenv("JAEGER_URL"), os.Getenv("SERVICE_NAME"), os.Getenv("SERVICE_VERSION"), os.Getenv("ENVIRONMENT"))

	_, span := tracer.Start(context.Background(), "init_test", "", "")
	defer span.End()

	span.SetAttributes(attribute.String("test", "test"))
	sentry.CaptureError(errors.New("test error"))
	time.Sleep(5 * time.Second)

	// tracer.InitOtel(os.Getenv("JAEGER_URL"), os.Getenv("SERVICE_NAME"), os.Getenv("SERVICE_VERSION"), os.Getenv("ENVIRONMENT"))
}
