package service

import (
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/logging"
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

func (u *UniversityImpl) Get(userID uint) ([]rdto.University, error) {
	return u.universityRepository.Get(userID)
}

func (u *UniversityImpl) Set(userID uint, universityIDs dto.IDs) error {
	err := u.universityRepository.Clear(userID)
	if err != nil {
		return err
	}

	return u.universityRepository.Set(userID, universityIDs)
}
