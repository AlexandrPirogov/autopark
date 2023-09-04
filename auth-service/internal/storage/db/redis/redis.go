package redis

import (
	"auth-service/internal/config"
	"context"
	"sync"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/redis/go-redis/v9"
)

// Instance of singleton
var singletonConn *redisConn = nil
var initOnce sync.Once

type redisConn struct {
	conn *redis.Client
}

func GetInstance() *redisConn {
	initOnce.Do(func() {
		singletonConn = new()
	})
	return singletonConn
}

func new() *redisConn {
	conn := redis.NewClient(&redis.Options{
		Addr:     "redis-auth:6379",
		Password: config.RedisPwd(), // no password set
		DB:       0,                 // use default DB
	})

	_, err := conn.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}

	return &redisConn{
		conn: conn,
	}

}

func (r *redisConn) SetRefreshToken(val string) error {
	err := r.conn.Set(context.Background(), val, "true", time.Hour*24).Err()
	if err != nil {
		log.Warn().Msgf("%v", err)
		return err
	}

	return nil
}
