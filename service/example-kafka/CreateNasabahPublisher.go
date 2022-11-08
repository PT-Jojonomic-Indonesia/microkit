package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/PT-Jojonomic-Indonesia/microkit/kafka"
	"github.com/PT-Jojonomic-Indonesia/microkit/service/example-kafka/entity"
)

func PublishNasabah(data *entity.Nasabah) error {
	key := fmt.Sprintf("create-nasabah-%v", time.Now().UnixNano())
	value, _ := json.Marshal(data)
	return kafka.Publish([]byte(key), value, nil) // without context deadline
}
