package service

import (
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/cache"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/models"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/repository"
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
}

type Parsing interface {
	ParseRating(universityName string, ratingURL string, userSnils string) (*dto.ParsingResult, error)
}

type University interface {
	GetAll() ([]dto.University, error)
	GetByID(id uint) (*models.University, error)
	GetForUser(userID uint) ([]dto.University, error)
	SetForUser(userID uint, universityIDs dto.IDs) error
}

type Direction interface {
	GetAll() ([]dto.UniversityDirections, error)
	GetByID(id uint) (*models.Direction, error)
	GetForUser(userID uint) ([]dto.UniversityDirections, error)
	GetForUserWithRating(userID uint) ([]dto.UniversityDirectionsWithRating, error)
	SetForUser(userID uint, directionIDs dto.IDs) error
}

type Service struct {
	Authorization
	User
	Parsing
	University
	Direction
}

func New(repository *repository.Repository, cache *cache.Cache) *Service {
	authorizationService := NewAuthorizationImpl(repository.User, cache.RefreshToken, cache.Blacklist)
	userService := NewUserImpl(repository.User)
	parsingService := NewParsingImpl(cache.RatingList)
	universityService := NewUniversityImpl(repository.University)
	directionService := NewDirectionImpl(repository.Direction, repository.User, universityService, parsingService)

	return &Service{
		Authorization: authorizationService,
		User:          userService,
		Parsing:       parsingService,
		University:    universityService,
		Direction:     directionService,
	}
}
