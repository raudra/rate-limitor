package config

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

var redisClient *redis.Client

func InitRedis() {
	ctx := context.TODO()

	log.Info().
		Msg("Setting Redis ")

	redisClient = redis.NewClient(&redis.Options{
		Addr:     "host.docker.internal:6300", // host:port of the redis server
		Password: "",                          // no password set
		DB:       0,                           // use default DB
	})

	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatal().
			Err(err).
			Msgf("Error while connecting to Redis")
	}
}

func RedisClient() *redis.Client {
	return redisClient
}
