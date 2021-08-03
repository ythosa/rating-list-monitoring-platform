package service

import (
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/logging"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/models"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository/rdto"
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
	return u.universityRepository.GetAll()
}

func (u *UniversityImpl) GetByID(id uint) (*models.University, error) {
	return u.universityRepository.GetByID(id)
}

func (u *UniversityImpl) GetForUser(userID uint) ([]rdto.University, error) {
	return u.universityRepository.GetForUser(userID)
}

func (u *UniversityImpl) SetForUser(userID uint, universityIDs dto.IDs) error {
	err := u.universityRepository.Clear(userID)
	if err != nil {
		return err
	}

	return u.universityRepository.SetForUser(userID, universityIDs)
}
