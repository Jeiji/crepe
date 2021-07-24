package storage

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var rdb *redis.Client

func Get(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	// if err != nil {
	// 	// panic(err)
	// }
	return val, err
}

func Set(key string, value string) error {
	err := rdb.Set(ctx, key, value, 0).Err()
	// if err != nil {
	// 	// panic(err)
	// }
	return err
}

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "crepe-redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
