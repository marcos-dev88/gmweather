package cache

import (
	"context"

	redis "github.com/go-redis/redis/v8"
)

type Client redis.Client

type cache struct {
	Client *Client
	Ctx    context.Context
}
