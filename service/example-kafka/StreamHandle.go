package main

import (
	"encoding/json"

	"github.com/PT-Jojonomic-Indonesia/microkit/kafka"
	"github.com/PT-Jojonomic-Indonesia/microkit/service/example-kafka/entity"

	kafka_go "github.com/segmentio/kafka-go"
)

func HandleKafkaStream(successConn, retryConn *kafka_go.Conn) func(ey, value []byte) {
	return func(key, value []byte) {
		nasabah := &entity.Nasabah{}
		if err := json.Unmarshal(value, nasabah); err != nil {
			kafka.PublishWithClient(retryConn, key, value, nil)
			return
		}

		// handle some business logic here

		kafka.PublishWithClient(successConn, key, value, nil)
	}
}
