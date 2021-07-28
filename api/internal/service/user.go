package service

import (
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/logging"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository/rdto"
)

type UserImpl struct {
	userRepository repository.User

	logger *logging.Logger
}

func NewUserImpl(userRepository repository.User) *UserImpl {
	return &UserImpl{
		userRepository: userRepository,
		logger:         logging.NewLogger("user service"),
	}
}

func (u *UserImpl) GetUsername(id uint8) (*dto.Username, error) {
	username, err := u.userRepository.GetUsername(id)
	if err != nil {
		u.logger.Error(err)

		return nil, err
	}

	return (*dto.Username)(username), nil
}

func (u *UserImpl) GetProfile(id uint8) (*dto.UserProfile, error) {
	userProfile, err := u.userRepository.GetProfile(id)
	if err != nil {
		u.logger.Error(err)

		return nil, err
	}

	return (*dto.UserProfile)(userProfile), nil
}

func (u *UserImpl) SetUniversities(id uint8, universityIDs dto.IDs) error {
	err := u.userRepository.ClearUniversities(id)
	if err != nil {
		return err
	}

	return u.userRepository.SetUniversities(id, universityIDs)
}

func (u *UserImpl) GetUniversities(id uint8) ([]rdto.University, error) {
	universities, err := u.userRepository.GetUniversities(id)

	return universities, err
}
