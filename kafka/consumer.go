package kafka

import (
	"context"
	"log"
	"strings"

	"github.com/segmentio/kafka-go"
)

var maxReadByte int = 10e6 // 10MB

func HandleReadStream(ctx context.Context, config *kafka.ReaderConfig, handler func(ctx context.Context, key, msg []byte)) {
	if config.MaxBytes == 0 {
		config.MaxBytes = maxReadByte
	}

	reader := kafka.NewReader(*config)
	defer reader.Close()

	streamHosts := strings.Join(config.Brokers, ",")
	log.Printf("\nstarting consume data from : %s with topic %s", streamHosts, config.Topic)

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("receive message at topic/partition/offset %v/%v/%v: %s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key))

		go consumerRecoverHandler(handler)(ctx, msg.Key, msg.Value)
	}
}

func consumerRecoverHandler(h func(ctx context.Context, key, msg []byte)) func(ctx context.Context, key, msg []byte) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	return h
}
