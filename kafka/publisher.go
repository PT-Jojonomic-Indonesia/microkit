package kafka

import (
	"context"
	"log"
	"sync"
	"time"

	kafka_go "github.com/segmentio/kafka-go"
)

var kafkaConn *kafka_go.Conn
var once sync.Once

// setup kafka to leader connection
func DialLeader(kafkaURL, topic string, partition int) error {
	var initError error

	once.Do(func() {
		conn, err := kafka_go.DialLeader(context.Background(), "tcp", kafkaURL, topic, partition)
		if err != nil {
			log.Fatal(err.Error())
		}

		kafkaConn = conn
	})
	return initError
}

func NewDial(kafkaURL, topic string, partition int) (*kafka_go.Conn, error) {
	return kafka_go.DialLeader(context.Background(), "tcp", kafkaURL, topic, partition)
}

func Close() error {
	return CloseWithClient(kafkaConn)
}

func CloseWithClient(conn *kafka_go.Conn) error {
	if conn != nil {
		return conn.Close()
	}
	return nil
}

var Publish = func(key, value []byte, deadline *time.Time) error {
	return PublishWithClient(kafkaConn, key, value, deadline)
}

var PublishWithClient = func(conn *kafka_go.Conn, key, value []byte, deadline *time.Time) error {
	if deadline != nil {
		conn.SetWriteDeadline(*deadline)
	}

	msg := kafka_go.Message{
		Key:   key,
		Value: value,
	}
	_, err := conn.WriteMessages(msg)
	return err
}
