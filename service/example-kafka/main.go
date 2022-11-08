package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"bitbucket.org/jojocoders/microkit/kafka"
	"bitbucket.org/jojocoders/microkit/server"
	"bitbucket.org/jojocoders/microkit/tracer"

	kafka_go "github.com/segmentio/kafka-go"

	"bitbucket.org/jojocoders/microkit/validator"

	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	godotenv.Load()

	router := getRoutes()

	url := os.Getenv("JAEGER_ENDPOINT")
	tracer.InitOtel(url, "Example Kafka", "v1.0.0", "development")

	validator.Init()

	// setup kafka publiseh reader
	kafkaUrl := os.Getenv("KAFKA_ENDPOINT")
	createTopic := os.Getenv("KAFKA_CREATE_TOPIC")
	partitionStr := os.Getenv("KAFKA_PARTITION")
	partition, err := strconv.Atoi(partitionStr)
	if err != nil {
		log.Panicf("partition must valid integer : %s", err)
	}

	if err := kafka.DialLeader(kafkaUrl, createTopic, partition); err != nil {
		log.Panic(err)
	}
	defer kafka.Close() // important to close

	// multiple kafka conn
	publishSuccessTopic := os.Getenv("KAFKA_SUCCESS_TOPIC")
	successKafkaConn, err := kafka.NewDial(kafkaUrl, publishSuccessTopic, partition)
	if err != nil {
		log.Panicf("can`t connect to success publisher : %s", err)
	}
	defer kafka.CloseWithClient(successKafkaConn)

	publishRetryTopic := os.Getenv("KAFKA_RETRY_TOPIC")
	retryKafkaConn, err := kafka.NewDial(kafkaUrl, publishRetryTopic, partition)
	if err != nil {
		log.Panicf("can`t connect to retry publisher : %s", err)
	}
	defer kafka.CloseWithClient(retryKafkaConn)

	// Handle stream connection
	config := kafka_go.ReaderConfig{
		Brokers:                []string{kafkaUrl},
		GroupID:                "create-nasabah",
		Topic:                  createTopic,
		AllowAutoTopicCreation: true,
	}

	go kafka.HandleReadStream(context.Background(), &config, HandleKafkaStream(successKafkaConn, retryKafkaConn))

	server.Serve("8001", router)
}
