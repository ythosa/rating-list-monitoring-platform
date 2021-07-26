package cache

import "time"

type RefreshToken interface {
	Save(userID uint8, token string, ttl time.Duration) error
	Get(userID uint8) (string, error)
	Delete(userID uint8) error
}

type Blacklist interface {
	Save(userID uint8, accessToken string, ttl time.Duration) error
	Get(userID uint8) error
	Delete(userID uint8) error
}

type Cache struct {
	RefreshToken
	Blacklist
}
