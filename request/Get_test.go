package request

import (
	"context"
	"sync"
	"testing"

	"github.com/PT-Jojonomic-Indonesia/microkit/tracer"
	"github.com/sony/gobreaker"
	"github.com/stretchr/testify/assert"
)

func TestGetWithCircuitBreaker(t *testing.T) {
	tracer.InitOtel("http://localhost:14268/api/traces", "test service", "v1.0.0", "testing")
	InitCircuitBreacker(nil, []string{"jsonplaceholder.typicode.com"})

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

			testUrl := "htt://jsonplaceholder.typicode.com/todos/1"
			if iterate == 9 {
				testUrl = "https://jsonplaceholder.typicode.com/todos/1"
			}
			res := make(map[string]interface{})
			errReq := Get(context.Background(), testUrl, &res)
			t.Log(errReq)
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
