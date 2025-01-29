package cache

import (
	"context"

	"github.com/Lucasanim/shortly/config"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var Instance *redis.Client

func InitializeRedis() {
	Instance = redis.NewClient(&redis.Options{
		Addr:     config.Env.RedisAddr,
		Password: config.Env.RedisPassword,
		DB:       config.Env.RedisDb,
	})
}

type Cache struct{}

func (c *Cache) Set(key string, value string) error {
	return Instance.Set(ctx, key, value, 0).Err()
}

func (c *Cache) Get(key string) (string, error) {
	return Instance.Get(ctx, key).Result()
}
