package request

import (
	"context"
	"log"
	"sync"
	"testing"

	"bitbucket.org/jojocoders/microkit/tracer"
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
	var closeStatus int
	var errCount int

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(iterate int) {
			defer wg.Done()

			testUrl := "https://jsonplaceholder.typicode.com/todos/100000"
			if iterate >= 8 {
				testUrl = "https://jsonplaceholder.typicode.com/todos/1"
			}
			res := make(map[string]interface{})
			err := Get(context.Background(), testUrl, &res)
			log.Printf("cp state : %v", cb.State())

			if err != nil {
				m.Lock()
				errCount += 1
				m.Unlock()
			}

			if cb.State() == gobreaker.StateClosed {
				m.Lock()
				closeStatus += 1
				m.Unlock()
			}

			if cb.State() == gobreaker.StateOpen {
				m.Lock()
				openStatus += 1
				m.Unlock()
			}
		}(i)
	}
	wg.Wait()
	mapCp.Delete("jsonplaceholder.typicode.com")

	assert.Equal(t, 3, openStatus)
	assert.Equal(t, 7, closeStatus)
	assert.Equal(t, 8, errCount)
}
