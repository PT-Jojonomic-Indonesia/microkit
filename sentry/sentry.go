package sentry

import (
	"log"
	"time"

	sentry_go "github.com/getsentry/sentry-go"
)

type Config sentry_go.ClientOptions

func Init(cfg Config) {
	err := sentry_go.Init(sentry_go.ClientOptions(cfg))
	if err != nil {
		log.Fatalf("Init Senty: %s", err)
	}
}

func Close() {
	sentry_go.Flush(2 * time.Second)
}

var Capture = sentry_go.CaptureMessage

var CaptureError = sentry_go.CaptureException
