package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gopkg.in/errgo.v2/fmt/errors"
)

type BlacklistImpl struct {
	rc *redis.Client
}

func NewBlacklistImpl(rc *redis.Client) *BlacklistImpl {
	return &BlacklistImpl{rc}
}

func (b *BlacklistImpl) Save(userID uint, accessToken string, ttl time.Duration) error {
	if err := b.rc.Set(redisCtx, b.formatKey(userID), accessToken, ttl).Err(); err != nil {
		return fmt.Errorf("error while saving user in blacklist cache: %w", err)
	}

	return nil
}

func (b *BlacklistImpl) Get(userID uint) error {
	if err := b.rc.Get(redisCtx, b.formatKey(userID)).Err(); err != nil {
		return errors.New("there is no user in the blacklist")
	}

	return nil
}

func (b *BlacklistImpl) Delete(userID uint) error {
	if err := b.rc.Del(redisCtx, b.formatKey(userID)).Err(); err != nil {
		return fmt.Errorf("error while deleting user from blacklist cache: %w", err)
	}

	return nil
}

func (b *BlacklistImpl) formatKey(userID uint) string {
	return fmt.Sprintf("bl_%d", userID)
}
