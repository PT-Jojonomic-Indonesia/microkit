package kafka

import (
	"errors"
	"log"

	kafka_go "github.com/segmentio/kafka-go"
)

var Health = func(kafkaUrl string) error {
	conn, err := kafka_go.Dial("tcp", kafkaUrl)
	if err != nil {
		log.Printf("kafka : %s", err)
		return errors.New("kafka is not available")
	}
	defer conn.Close()
	return nil
}
