package cache

import "time"

type RefreshToken interface {
	Save(userID int, token string, ttl time.Duration) error
	Get(userID int) (string, error)
	Delete(userID int) error
}

type Blacklist interface {
	Save(userID int, accessToken string, ttl time.Duration) error
	Get(userID int) error
	Delete(userID int) error
}

type Cache struct {
	RefreshToken
	Blacklist
}
