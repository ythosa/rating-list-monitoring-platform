package service

import (
	"fmt"

	"github.com/ythosa/rating-list-monitoring-platform-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/logging"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/models"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/repository"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/repository/rdto"
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

func (u *UniversityImpl) GetAll() ([]rdto.University, error) {
	universities, err := u.universityRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error while getting all universities by repository: %w", err)
	}

	return universities, nil
}

func (u *UniversityImpl) GetByID(id uint) (*models.University, error) {
	university, err := u.universityRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("error while getting university by ID from repository: %w", err)
	}

	return university, nil
}

func (u *UniversityImpl) GetForUser(userID uint) ([]rdto.University, error) {
	universities, err := u.universityRepository.GetForUser(userID)
	if err != nil {
		return nil, fmt.Errorf("error while getting user universities from repository: %w", err)
	}

	return universities, nil
}

func (u *UniversityImpl) SetForUser(userID uint, universityIDs dto.IDs) error {
	if err := u.universityRepository.Clear(userID); err != nil {
		return fmt.Errorf("error while clearing user universities by repository: %w", err)
	}

	if err := u.universityRepository.SetForUser(userID, universityIDs); err != nil {
		return fmt.Errorf("error while setting universities for user from repository: %w", err)
	}

	return nil
}
