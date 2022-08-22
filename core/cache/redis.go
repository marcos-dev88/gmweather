package cache

import (
	"context"
	"errors"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type Cache interface {
	Set(key string, data interface{}, ttl time.Duration) error
	Get(key string) ([]byte, error)
}

func NewClient(addr, pass string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})
}

func New(client *redis.Client, ctx context.Context) Cache {
	return &cache{Client: client, Ctx: ctx}
}

func (c cache) Set(key string, data interface{}, ttl time.Duration) error {
	client := c.Client
	defer client.Close()
	return client.Set(c.Ctx, key, data, ttl).Err()
}

func (c cache) Get(key string) ([]byte, error) {
	client := c.Client
	defer client.Close()
	b, err := client.Get(c.Ctx, key).Bytes()

	if err == redis.Nil {
		return nil, errors.New("no_cache_data")
	}

	return b, nil
}
