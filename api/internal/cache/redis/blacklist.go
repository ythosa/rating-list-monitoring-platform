package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"gopkg.in/errgo.v2/fmt/errors"
	"time"
)

type BlacklistImpl struct {
	rc *redis.Client
}

func NewBlacklistImpl(rc *redis.Client) *BlacklistImpl {
	return &BlacklistImpl{rc}
}

func (b *BlacklistImpl) Save(userID int, accessToken string, ttl time.Duration) error {
	return b.rc.Set(redisCtx, b.formatKey(userID), accessToken, ttl).Err()
}

func (b *BlacklistImpl) Get(userID int) error {
	if err := b.rc.Get(redisCtx, b.formatKey(userID)).Err(); err != nil {
		return errors.New("there is no user in the blacklist")
	}

	return nil
}

func (b *BlacklistImpl) Delete(userID int) error {
	return b.rc.Del(redisCtx, b.formatKey(userID)).Err()
}

func (b *BlacklistImpl) formatKey(userID int) string {
	return fmt.Sprintf("bl_%d", userID)
}
