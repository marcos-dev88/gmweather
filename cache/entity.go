package cache

import (
	"context"

	redis "github.com/go-redis/redis/v8"
)

type cache struct {
	Client *redis.Client
	Ctx    context.Context
}
