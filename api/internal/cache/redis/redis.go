package redis

import (
	"context"

	"github.com/go-redis/redis/v8"

	"github.com/ythosa/rating-list-monitoring-platform-api/internal/cache"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/config"
)

var redisCtx = context.TODO()

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
		RefreshToken: NewRefreshTokenImpl(rc),
		Blacklist:    NewBlacklistImpl(rc),
		RatingList:   NewRatingListImpl(rc),
	}
}
