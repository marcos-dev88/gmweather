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
	Close() error
}

func NewClient(addr, pass string, db int) *Client {
	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})
	return (*Client)(c)
}

func New(client *Client, ctx context.Context) Cache {
	return &cache{Client: client, Ctx: ctx}
}

func (c cache) Set(key string, data interface{}, ttl time.Duration) error {
	client := c.Client
	e := client.Set(c.Ctx, key, data, ttl)

	if e.Err() != nil {
		return e.Err()
	}
	return nil
}

func (c cache) Get(key string) ([]byte, error) {
	client := c.Client
	b, err := client.Get(c.Ctx, key).Bytes()

	if err == redis.Nil {
		return nil, errors.New("no_cache_data")
	}

	return b, nil
}

func (c cache) Close() error {
	return c.Client.Close()
}
