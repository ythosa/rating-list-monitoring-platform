package repository

import (
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/models"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository/rdto"
)

type User interface {
	Create(user rdto.UserCreating) (int, error)
	GetUserByNickname(nickname string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
	UpdatePassword(id int, password string) error
	PatchUser(id int, data rdto.UserPatching) error
}

type Repository struct {
	User
}
