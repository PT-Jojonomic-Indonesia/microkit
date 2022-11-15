package rediskit

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

var initializeRedis = false

var (
	redisClient *redis.Client
	err         error
)

// InitRedisClient is
func InitRedisClient(redisAddress string, password string, db int) (err error) {
	if !initializeRedis {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     redisAddress,
			Password: password,
			DB:       db,
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
func GetRedisClient(db int) (*redis.Client, error) {
	if !initializeRedis || redisClient == nil {
		dbString := os.Getenv("REDIS_DB")
		db, _ := strconv.ParseInt(dbString, 10, 64)
		err := InitRedisClient(os.Getenv("REDIS_ADDRESS"), os.Getenv("REDIS_PASSWORD"), int(db))
		if err != nil {
			return nil, errors.New("failed to initialize redis: " + err.Error())
		}
	}
	return redisClient, nil
}
