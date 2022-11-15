package observability

import (
	"os"

	"github.com/PT-Jojonomic-Indonesia/microkit/sentry"
	"github.com/PT-Jojonomic-Indonesia/microkit/tracer"
)

// InitDefault is a function to initialize default observability.
// This function will initialize tracer and sentry.
// This function will be called in main function.
// This function need environment variable: JAEGER_URL, SERVICE_NAME, SERVICE_VERSION, ENVIRONMENT, SENTRY_DSN
func InitDefault() {
	tracer.InitOtel(os.Getenv("JAEGER_URL"), os.Getenv("SERVICE_NAME"), os.Getenv("SERVICE_VERSION"), os.Getenv("ENVIRONMENT"))
	cfg := sentry.Config{
		Dsn:         os.Getenv("SENTRY_DSN"),
		Release:     os.Getenv("SERVICE_VERSION"),
		Environment: os.Getenv("ENVIRONMENT"),
	}
	sentry.Init(cfg)
}
