package service

import (
	"fmt"

	"github.com/ythosa/rating-list-monitoring-platform-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/models"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/repository"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/repository/rdto"
	"github.com/ythosa/rating-list-monitoring-platform-api/pkg/logging"
)

type UniversityImpl struct {
	universityRepository repository.University
	logger               *logging.Logger
}

func NewUniversityImpl(universityRepository repository.University) *UniversityImpl {
	return &UniversityImpl{
		universityRepository: universityRepository,
		logger:               logging.NewLogger("university service"),
	}
}

func (s *UniversityImpl) GetAll() ([]dto.University, error) {
	universities, err := s.universityRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error while getting all universities by repository: %w", err)
	}

	return mapRepositoryUniversitiesResultToDTOs(universities), nil
}

func (s *UniversityImpl) GetByID(id uint) (*models.University, error) {
	university, err := s.universityRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("error while getting university by ID from repository: %w", err)
	}

	return university, nil
}

func (s *UniversityImpl) GetForUser(userID uint) ([]dto.University, error) {
	universities, err := s.universityRepository.GetForUser(userID)
	if err != nil {
		return nil, fmt.Errorf("error while getting user universities from repository: %w", err)
	}

	return mapRepositoryUniversitiesResultToDTOs(universities), nil
}

func mapRepositoryUniversitiesResultToDTOs(universities []rdto.University) []dto.University {
	result := make([]dto.University, 0)
	for _, u := range universities {
		result = append(result, dto.University{
			ID:       u.ID,
			Name:     u.Name,
			FullName: u.FullName,
		})
	}

	return result
}

func (s *UniversityImpl) SetForUser(userID uint, universityIDs dto.IDs) error {
	if err := s.universityRepository.Clear(userID); err != nil {
		return fmt.Errorf("error while clearing user universities by repository: %w", err)
	}

	if err := s.universityRepository.SetForUser(userID, universityIDs); err != nil {
		return fmt.Errorf("error while setting universities for user from repository: %w", err)
	}

	return nil
}
