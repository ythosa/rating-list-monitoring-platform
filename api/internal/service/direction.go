package service

import (
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/logging"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository"
)

type DirectionImpl struct {
	directionRepository repository.Direction
	logger              *logging.Logger
}

func NewDirectionImpl(directionRepository repository.Direction) *DirectionImpl {
	return &DirectionImpl{
		directionRepository: directionRepository,
		logger:              logging.NewLogger("directions service"),
	}
}

func (u *DirectionImpl) GetAll() (map[string][]dto.Direction, error) {
	directions, err := u.directionRepository.GetAll()
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

func (u *DirectionImpl) Get(userID uint) (map[string][]dto.Direction, error) {
	directions, err := u.directionRepository.Get(userID)
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

func (u *DirectionImpl) Set(userID uint, directionIDs dto.IDs) error {
	err := u.directionRepository.Clear(userID)
	if err != nil {
		return err
	}

	return u.directionRepository.Set(userID, directionIDs)
}
