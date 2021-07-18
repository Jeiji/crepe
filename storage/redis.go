package storage

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var rdb *redis.Client

func Get(key string) string {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	return val
}

func Set(key string, value string) {
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "crepe-redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
