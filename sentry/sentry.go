package sentry

import (
	"fmt"
	"log"
	"time"

	sentry_go "github.com/getsentry/sentry-go"
)

func Init(cfg sentry_go.ClientOptions) {
	err := sentry_go.Init(cfg)
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
