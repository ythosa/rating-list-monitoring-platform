package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type RefreshToken struct {
	rc *redis.Client
}

func NewRefreshToken(rc *redis.Client) *RefreshToken {
	return &RefreshToken{rc}
}

func (r *RefreshToken) Save(userID int, token string, ttl time.Duration) error {
	return r.rc.Set(redisCtx, r.formatKey(userID), token, ttl).Err()
}

func (r *RefreshToken) Get(userID int) (string, error) {
	return r.rc.Get(redisCtx, r.formatKey(userID)).Result()
}

func (r *RefreshToken) Delete(userID int) error {
	return r.rc.Del(redisCtx, r.formatKey(userID)).Err()
}

func (r *RefreshToken) formatKey(userId int) string {
	return fmt.Sprintf("rt_%d", userId)
}
