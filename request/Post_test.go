package request

import (
	"context"
	"sync"
	"testing"

	"github.com/PT-Jojonomic-Indonesia/microkit/tracer"
	"github.com/sony/gobreaker"
	"github.com/stretchr/testify/assert"
)

func TestPostWithCircuitBreaker(t *testing.T) {
	tracer.InitOtel("http://localhost:14268/api/traces", "test service", "v1.0.0", "testing")

	InitCircuitBreacker(&gobreaker.Settings{
		Name: "HTTP GET",
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 3 && failureRatio > 0.6
		},
	}, []string{"jsonplaceholder.typicode.com"})

	cbValue, ok := mapCp.Load("jsonplaceholder.typicode.com")
	assert.Equal(t, true, ok)
	cb := cbValue.(*gobreaker.CircuitBreaker)

	var wg sync.WaitGroup
	var m sync.Mutex
	var openStatus int

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(iterate int) {
			defer wg.Done()

			testUrl := "htt://jsonplaceholder.typicode.com/posts"
			if iterate > 8 {
				testUrl = "https://jsonplaceholder.typicode.com/posts"
			}

			req := map[string]interface{}{
				"title":  "foo",
				"body":   "bar",
				"userId": 2,
			}
			res := make(map[string]interface{})
			errReq := Post(context.Background(), testUrl, req, &res)
			if cb.State().String() == "open" && errReq == gobreaker.ErrOpenState {
				m.Lock()
				openStatus += 1
				m.Unlock()
			}
		}(i)
	}
	wg.Wait()
	mapCp.Delete("jsonplaceholder.typicode.com")

	assert.GreaterOrEqual(t, openStatus, 1)
}
