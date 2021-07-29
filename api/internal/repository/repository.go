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
}

type University interface {
	GetAll() ([]rdto.University, error)
	Get(userID uint) ([]rdto.University, error)
	Set(userID uint, universityIDs dto.IDs) error
	Clear(userID uint) error
}

type Direction interface {
	GetAll() ([]rdto.Direction, error)
	GetByID(id uint) (rdto.Direction, error)
	Get(userID uint) ([]rdto.Direction, error)
	Set(userID uint, directionIDs dto.IDs) error
	Clear(userID uint) error
}

type Repository struct {
	User
	University
	Direction
}
