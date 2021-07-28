package repository

import (
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/models"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository/rdto"
)

type User interface {
	Create(user rdto.UserCreating) (uint, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdatePassword(id uint, password string) error
	PatchUser(id uint, data rdto.UserPatching) error
	GetUsername(id uint) (*rdto.Username, error)
	GetProfile(id uint) (*rdto.UserProfile, error)
	SetUniversities(id uint, universityIDs dto.IDs) error
	GetUniversities(id uint) ([]rdto.University, error)
	ClearUniversities(id uint) error
	SetDirections(id uint, directionIDs dto.IDs) error
	GetDirections(id uint) ([]rdto.Direction, error)
	ClearDirections(id uint) error
}

type Repository struct {
	User
}
