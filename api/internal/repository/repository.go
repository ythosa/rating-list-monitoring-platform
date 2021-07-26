package repository

import (
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/models"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository/rdto"
)

type User interface {
	Create(user rdto.UserCreating) (uint8, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByID(id uint8) (*models.User, error)
	UpdatePassword(id uint8, password string) error
	PatchUser(id uint8, data rdto.UserPatching) error
}

type Repository struct {
	User
}
