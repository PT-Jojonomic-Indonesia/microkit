package rediskit

import (
	"errors"
	"log"
	"os"

	"github.com/go-redis/redis"
)

var initializeRedis = false

var (
	redisClient *redis.Client
	err         error
)

// InitRedisClient is
func InitRedisClient(redisAddress string, password string) (err error) {
	if !initializeRedis {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     redisAddress,
			Password: password,
			DB:       0,
		})
		_, err = redisClient.Ping().Result()
		if err != nil {
			log.Println("Failed connect to Redis", err)
			return
		}

		initializeRedis = true
	}
	return
}

// GetRedisClient is
func GetRedisClient() (*redis.Client, error) {
	if initializeRedis == false || redisClient == nil {
		err := InitRedisClient(os.Getenv("REDIS_ADDRESS"), os.Getenv("REDIS_PASSWORD"))
		if err != nil {
			return nil, errors.New("Failed to initialize Redis")
		}
	}
	return redisClient, nil
}
