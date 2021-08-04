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
	if err := r.rc.Set(redisCtx, r.formatKey(url), data, ttl).Err(); err != nil {
		return fmt.Errorf("error while caching rating list: %w", err)
	}

	return nil
}

func (r *RatingListImpl) Get(url string) (string, error) {
	ratingList, err := r.rc.Get(redisCtx, r.formatKey(url)).Result()
	if err != nil {
		return "", fmt.Errorf("error while getting rating list from cache: %w", err)
	}

	return ratingList, nil
}

func (r *RatingListImpl) formatKey(url string) string {
	return fmt.Sprintf("rl_%s", url)
}
