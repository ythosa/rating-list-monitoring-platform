package service

import (
	"fmt"

	"github.com/ythosa/rating-list-monitoring-platform-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/repository"
	"github.com/ythosa/rating-list-monitoring-platform-api/pkg/logging"
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

func (s *UserImpl) GetUsername(id uint) (*dto.Username, error) {
	username, err := s.userRepository.GetUsername(id)
	if err != nil {
		s.logger.Error(err)

		return nil, fmt.Errorf("error while getting user username: %w", err)
	}

	return (*dto.Username)(username), nil
}

func (s *UserImpl) GetProfile(id uint) (*dto.UserProfile, error) {
	userProfile, err := s.userRepository.GetProfile(id)
	if err != nil {
		s.logger.Error(err)

		return nil, fmt.Errorf("error while getting user profile: %w", err)
	}

	return (*dto.UserProfile)(userProfile), nil
}
