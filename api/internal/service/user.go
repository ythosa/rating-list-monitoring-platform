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

func (u *UserImpl) SetUniversities(id uint, universityIDs dto.IDs) error {
	err := u.userRepository.ClearUniversities(id)
	if err != nil {
		return err
	}

	return u.userRepository.SetUniversities(id, universityIDs)
}

func (u *UserImpl) GetUniversities(id uint) ([]rdto.University, error) {
	return u.userRepository.GetUniversities(id)
}

func (u *UserImpl) SetDirections(id uint, directionIDs dto.IDs) error {
	err := u.userRepository.ClearUniversities(id)
	if err != nil {
		return err
	}

	return u.userRepository.SetDirections(id, directionIDs)
}

func (u *UserImpl) GetDirections(id uint) (map[string][]dto.Direction, error) {
	directions, err := u.userRepository.GetDirections(id)
	if err != nil {
		return nil, err
	}

	directionsUniversity := make(map[string][]dto.Direction)
	for _, d := range directions {
		direction := dto.Direction{
			ID:   d.DirectionID,
			Name: d.DirectionName,
		}
		directionsUniversity[d.UniversityName] = append(directionsUniversity[d.UniversityName], direction)
	}

	return directionsUniversity, nil
}
