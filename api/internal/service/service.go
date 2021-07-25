package service

import (
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/cache"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository"
)

type Authorization interface {
	SignUpUser(userData dto.SigningUp) (int, error)
	GenerateTokens(userCredentials dto.UserCredentials) (*dto.AuthorizationTokens, error)
	RefreshTokens(refreshToken string) (*dto.AuthorizationTokens, error)
	LogoutUser(userID int, accessToken string) error
	IsUserLogout(userID int) bool
}

type Service struct {
	Authorization
}

func New(repository *repository.Repository, cache *cache.Cache) *Service {
	return &Service{
		Authorization: NewAuthorizationImpl(repository.User, cache.RefreshToken, cache.Blacklist),
	}
}
