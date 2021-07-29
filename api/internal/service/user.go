package service

import (
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/logging"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository"
)

type UserImpl struct {
	userRepository repository.User
	logger         *logging.Logger
}

func NewUserImpl(userRepository repository.User) *UserImpl {
	return &UserImpl{
		userRepository: userRepository,
		logger:         logging.NewLogger("user service"),
	}
}

func (u *UserImpl) GetUsername(id uint) (*dto.Username, error) {
	username, err := u.userRepository.GetUsername(id)
	if err != nil {
		u.logger.Error(err)

		return nil, err
	}

	return (*dto.Username)(username), nil
}

func (u *UserImpl) GetProfile(id uint) (*dto.UserProfile, error) {
	userProfile, err := u.userRepository.GetProfile(id)
	if err != nil {
		u.logger.Error(err)

		return nil, err
	}

	return (*dto.UserProfile)(userProfile), nil
}
