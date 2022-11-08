package request

import (
	"context"
	"fmt"
	"net/http"
	url_parser "net/url"

	"github.com/sony/gobreaker"
)

func Post(ctx context.Context, url string, reqData interface{}, resp any) (err error) {
	p, err := url_parser.Parse(url)
	if err != nil {
		return fmt.Errorf("%s is not valid url", url)
	}

	if cbValue, ok := mapCp.Load(p.Hostname()); ok {
		cb := cbValue.(*gobreaker.CircuitBreaker)
		_, err := cb.Execute(func() (interface{}, error) {
			return nil, request(ctx, http.MethodPost, url, reqData, resp)
		})

		if err != nil {
			return err
		}
	}

	return request(ctx, http.MethodPost, url, reqData, resp)
}
