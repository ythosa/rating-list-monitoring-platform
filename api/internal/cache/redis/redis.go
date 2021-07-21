package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/cache"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/config"
)

var redisCtx = context.TODO()

const emptyValue = ""

func NewClient(cfg *config.Cache) *redis.Client {
	rc := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return rc
}

func NewCache(rc *redis.Client) *cache.Cache {
	return &cache.Cache{
		RefreshToken: NewRefreshToken(rc),
	}
}
