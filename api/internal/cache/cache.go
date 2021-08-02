package cache

import "time"

type RefreshToken interface {
	Save(userID uint, token string, ttl time.Duration) error
	Get(userID uint) (string, error)
	Delete(userID uint) error
}

type Blacklist interface {
	Save(userID uint, accessToken string, ttl time.Duration) error
	Get(userID uint) error
	Delete(userID uint) error
}

type RatingList interface {
	Save(url string, data string, ttl time.Duration) error
	Get(url string) (string, error)
}

type Cache struct {
	RefreshToken
	Blacklist
	RatingList
}
