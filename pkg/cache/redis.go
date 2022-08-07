package cache

import (
	"context"
	"fmt"
	redisClient "github.com/go-redis/redis/v8"
	"time"
)

const DEFAULT_TTL = 0
const OneDayTtl = time.Hour * 24

var redis *redisClient.Client
var ctx = context.Background()

func Initialize() *redisClient.Client {
	redisC := redisClient.NewClient(&redisClient.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	redis = redisC

	return redisC
}

func Set(key, data string, ttl time.Duration) {
	err := redis.Set(ctx, key, data, ttl).Err()
	if err != nil {
		fmt.Println(err)

		panic(err)
	}
}

func Get(key string) (value string, isExists bool) {
	val, err := redis.Get(ctx, key).Result()

	if err == redisClient.Nil {
		return "", false
	} else if err != nil {
		panic(err)
	} else {
		return val, true
	}
}

func Exists(key string) int64 {
	val, err := redis.Exists(ctx, key).Result()

	if err != nil {
		panic(err)
	} else {
		return val
	}
}

func Keys(key string) []string {
	val, err := redis.Keys(ctx, key).Result()

	if err != nil {
		panic(err)
	} else {
		return val
	}
}
