package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type RefreshTokenImpl struct {
	rc *redis.Client
}

func NewRefreshTokenImpl(rc *redis.Client) *RefreshTokenImpl {
	return &RefreshTokenImpl{rc}
}

func (r *RefreshTokenImpl) Save(userID int, token string, ttl time.Duration) error {
	return r.rc.Set(redisCtx, r.formatKey(userID), token, ttl).Err()
}

func (r *RefreshTokenImpl) Get(userID int) (string, error) {
	return r.rc.Get(redisCtx, r.formatKey(userID)).Result()
}

func (r *RefreshTokenImpl) Delete(userID int) error {
	return r.rc.Del(redisCtx, r.formatKey(userID)).Err()
}

func (r *RefreshTokenImpl) formatKey(userID int) string {
	return fmt.Sprintf("rt_%d", userID)
}
