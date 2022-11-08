package request

import (
	"sync"

	"github.com/sony/gobreaker"
)

var mapCp sync.Map

func InitCircuitBreacker(setting *gobreaker.Settings, hosts []string) {
	if setting == nil {
		setting = new(gobreaker.Settings)
		setting.Name = "HTTP GET"
	}

	for _, h := range hosts {
		mapCp.Store(h, gobreaker.NewCircuitBreaker(*setting))
	}
}
