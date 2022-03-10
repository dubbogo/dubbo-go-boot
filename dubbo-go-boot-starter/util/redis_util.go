package util

import (
	"github.com/go-redis/redis/v8"
)

func NewRedisDb(addr, password string, db, maxRetries int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:       addr,
		Password:   password,
		DB:         db,
		MaxRetries: maxRetries,
	})
	return client
}
