package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type RatingListImpl struct {
	rc *redis.Client
}

func NewRatingListImpl(rc *redis.Client) *RatingListImpl {
	return &RatingListImpl{rc: rc}
}

func (r *RatingListImpl) Save(url string, data string, ttl time.Duration) error {
	return r.rc.Set(redisCtx, r.formatKey(url), data, ttl).Err()
}

func (r *RatingListImpl) Get(url string) (string, error) {
	return r.rc.Get(redisCtx, r.formatKey(url)).Result()
}

func (r *RatingListImpl) formatKey(url string) string {
	return fmt.Sprintf("rl_%s", url)
}
