package database

import (
	"os"
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func createClient(dbNo int) *redis.Client {
	rdb := redis.NewClinet(&redis.Options{
		Addr : os.Getenv("DB_ADDR"),
		Password : os.Getenv("DB_PASS"),
		DB : dbNo,
	})

	return rdb
}
