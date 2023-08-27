package redis

import (
	"context"
	"log"
	"sync"
	"time"

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
		Addr:     "auth-redis:6379",
		Password: "secret", // no password set
		DB:       0,        // use default DB
	})

	_, err := conn.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}

	return &redisConn{
		conn: conn,
	}

}

func (r *redisConn) SetRefreshToken(val string) error {
	err := r.conn.Set(context.Background(), val, "true", time.Hour*24).Err()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
