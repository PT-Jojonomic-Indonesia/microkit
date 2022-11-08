package sentry

import (
	"fmt"
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

func Capture(data interface{}) {
	sentry_go.CaptureMessage(fmt.Sprintf("%+v", data))
}

func CaptureError(err error) {
	sentry_go.CaptureException(err)
}
