package service

import (
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/cache"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository"
)

type Authorization interface {
	SignUpUser(userData dto.SigningUp) (uint8, error)
	GenerateTokens(userCredentials dto.UserCredentials) (*dto.AuthorizationTokens, error)
	RefreshTokens(refreshToken string) (*dto.AuthorizationTokens, error)
	LogoutUser(userID uint8, accessToken string) error
	IsUserLogout(userID uint8) bool
}

type User interface {
	GetUsername(id uint8) (*dto.Username, error)
	GetProfile(id uint8) (*dto.UserProfile, error)
	SetUniversities(id uint8, universityIDs dto.IDs) error
}

type Service struct {
	Authorization
	User
}

func New(repository *repository.Repository, cache *cache.Cache) *Service {
	return &Service{
		Authorization: NewAuthorizationImpl(repository.User, cache.RefreshToken, cache.Blacklist),
		User:          NewUserImpl(repository.User),
	}
}
