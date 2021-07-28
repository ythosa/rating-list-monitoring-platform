package service

import (
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/cache"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository/rdto"
)

type Authorization interface {
	SignUpUser(userData dto.SigningUp) (uint, error)
	GenerateTokens(userCredentials dto.UserCredentials) (*dto.AuthorizationTokens, error)
	RefreshTokens(refreshToken string) (*dto.AuthorizationTokens, error)
	LogoutUser(userID uint, accessToken string) error
	IsUserLogout(userID uint) bool
}

type User interface {
	GetUsername(id uint) (*dto.Username, error)
	GetProfile(id uint) (*dto.UserProfile, error)
	SetUniversities(id uint, universityIDs dto.IDs) error
	GetUniversities(id uint) ([]rdto.University, error)
	SetDirections(id uint, directionIDs dto.IDs) error
	GetDirections(id uint) (map[string][]dto.Direction, error)
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
