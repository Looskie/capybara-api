package utils

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func SetRedisDB() {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))

	if err != nil {
		panic(err)
	}
	rdb = redis.NewClient(opt)
}

func GetRedisDB() *redis.Client {
	return rdb
}

func RedisGet(key string) (string, error) {
	value, err := rdb.Get(ctx, key).Result()

	if err != nil {
		return "", err
	}

	return value, nil
}

func RedisSet(key, value string, days int) (bool, error) {
	err := rdb.Set(ctx, key, value, time.Duration(days)*time.Hour*24).Err()

	if err != nil {
		return false, err
	}

	return true, nil
}

func RedisDel(key string) (bool, error) {
	err := rdb.Del(ctx, key).Err()

	if err != nil {
		return false, err
	}

	return true, nil
}
