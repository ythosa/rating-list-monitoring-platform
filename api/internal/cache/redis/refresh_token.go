package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RefreshTokenImpl struct {
	rc *redis.Client
}

func NewRefreshTokenImpl(rc *redis.Client) *RefreshTokenImpl {
	return &RefreshTokenImpl{rc}
}

func (r *RefreshTokenImpl) Save(userID uint, token string, ttl time.Duration) error {
	if err := r.rc.Set(redisCtx, r.formatKey(userID), token, ttl).Err(); err != nil {
		return fmt.Errorf("error while caching refresh token: %w", err)
	}

	return nil
}

func (r *RefreshTokenImpl) Get(userID uint) (string, error) {
	refreshToken, err := r.rc.Get(redisCtx, r.formatKey(userID)).Result()
	if err != nil {
		return "", fmt.Errorf("error while getting refresh token from cache: %w", err)
	}

	return refreshToken, nil
}

func (r *RefreshTokenImpl) Delete(userID uint) error {
	if err := r.rc.Del(redisCtx, r.formatKey(userID)).Err(); err != nil {
		return fmt.Errorf("error while deleting user refresh token from cache: %w", err)
	}

	return nil
}

func (r *RefreshTokenImpl) formatKey(userID uint) string {
	return fmt.Sprintf("rt_%d", userID)
}
