package request

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	url_parser "net/url"

	"github.com/PT-Jojonomic-Indonesia/microkit/tracer"

	"github.com/getsentry/sentry-go"
	"github.com/sony/gobreaker"
	"go.opentelemetry.io/otel/attribute"
)

func Get(ctx context.Context, url string, resp any) (err error) {
	p, err := url_parser.Parse(url)
	if err != nil {
		return fmt.Errorf("%s is not valid url", url)
	}

	if cbValue, ok := mapCp.Load(p.Hostname()); ok {
		cb := cbValue.(*gobreaker.CircuitBreaker)
		_, err := cb.Execute(func() (interface{}, error) {
			return nil, request(ctx, http.MethodGet, url, nil, resp)
		})

		if err != nil {
			return err
		}
	}

	return request(ctx, http.MethodGet, url, nil, resp)
}

func request(ctx context.Context, method string, url string, requestData any, resp any) (err error) {

	_, span := tracer.Start(ctx, "Call "+url, "", "")
	defer span.End()

	payload, err := json.Marshal(requestData)
	if err != nil {
		sentry.CaptureException(err)
		span.RecordError(err)
		return
	}

	span.SetAttributes(attribute.String("payload", string(payload)))
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(payload))

	if err != nil {
		sentry.CaptureException(err)
		span.RecordError(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("trace-id", span.SpanContext().TraceID().String())
	req.Header.Set("span-id", span.SpanContext().SpanID().String())

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		sentry.CaptureException(err)
		span.RecordError(err)
		return
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		sentry.CaptureException(err)
		span.RecordError(err)
		return
	}

	err = json.Unmarshal(responseData, resp)
	if err != nil {
		sentry.CaptureException(err)
		span.RecordError(err)
		return
	}

	result, err := json.Marshal(resp)
	if err != nil {
		sentry.CaptureException(err)
		span.RecordError(err)
		return
	}
	span.SetAttributes(attribute.String("result", string(result)))
	return
}
