package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/PT-Jojonomic-Indonesia/microkit/tracer"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	os.Setenv("SERVICE_NAME", "observability")
	os.Setenv("SERVICE_VERSION", "1.0.0")

	tracer.InitOtel(os.Getenv("JAEGER_URL"), os.Getenv("SERVICE_NAME"), os.Getenv("SERVICE_VERSION"), os.Getenv("ENVIRONMENT"))

	ctx := context.Background()

	// Cleanly shutdown and flush telemetry when the application exits.
	defer func(ctx context.Context) {
		if err := tracer.TracerProvider.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}(ctx)
	_, span := tracer.Start(ctx, "init_test", "", "")
	defer span.End()

	time.Sleep(5 * time.Second)

	// tracer.InitOtel(os.Getenv("JAEGER_URL"), os.Getenv("SERVICE_NAME"), os.Getenv("SERVICE_VERSION"), os.Getenv("ENVIRONMENT"))
}
