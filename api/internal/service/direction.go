package service

import (
	"context"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/logging"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/pkg/ratingparser"
	"golang.org/x/sync/errgroup"
	"sync"
)

type DirectionImpl struct {
	directionRepository repository.Direction
	userRepository      repository.User
	universityService   University
	logger              *logging.Logger
}

func NewDirectionImpl(
	directionRepository repository.Direction,
	userRepository repository.User,
	universityService University,
) *DirectionImpl {
	return &DirectionImpl{
		directionRepository: directionRepository,
		userRepository:      userRepository,
		universityService:   universityService,
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

func (u *DirectionImpl) GetWithRating(userID uint) (map[string][]dto.DirectionWithRating, error) {
	directions, err := u.directionRepository.Get(userID)
	if err != nil {
		u.logger.Error(err)

		return nil, err
	}

	userSnils, err := u.userRepository.GetSnils(userID)
	if err != nil {
		u.logger.Error(err)

		return nil, err
	}
	formattedSnils := ratingparser.FormatSnils(userSnils.Snils)

	var mu sync.Mutex

	errs, _ := errgroup.WithContext(context.TODO())
	directionsUniversity := make(map[string][]dto.DirectionWithRating)
	for _, d := range directions {
		direction := d
		errs.Go(func() error {
			parsingResult, err := ratingparser.ParseRating(
				direction.UniversityName,
				direction.DirectionURL,
				formattedSnils,
			)

			switch err {
			case ratingparser.UserNotFoundErr:
				parsingResult = &ratingparser.EmptyResult
			default:
				if err != nil {
					return err
				}
			}

			directionWithRating := dto.DirectionWithRating{
				ID:               direction.DirectionID,
				Name:             direction.DirectionName,
				Position:         parsingResult.Position,
				Score:            parsingResult.Score,
				PriorityOneUpper: parsingResult.PriorityOneUpper,
				BudgetPlaces:     parsingResult.BudgetPlaces,
			}

			mu.Lock()
			directionsUniversity[direction.UniversityName] = append(
				directionsUniversity[direction.UniversityName],
				directionWithRating,
			)
			mu.Unlock()

			return nil
		})
	}

	if err := errs.Wait(); err != nil {
		u.logger.Error(err)

		return nil, err
	}

	return directionsUniversity, nil
}

func (u *DirectionImpl) Set(userID uint, directionIDs dto.IDs) error {
	err := u.directionRepository.Clear(userID)
	if err != nil {
		return err
	}

	if err := u.directionRepository.Set(userID, directionIDs); err != nil {
		return err
	}

	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)

	universityIDs := make([]uint, 0)
	for _, directionID := range directionIDs.IDs {
		wg.Add(1)
		go func(id uint) {
			d, _ := u.directionRepository.GetByID(id)

			mu.Lock()
			stored := false
			for _, uID := range universityIDs {
				if uID == d.UniversityID {
					stored = true
					break
				}
			}
			if !stored {
				universityIDs = append(universityIDs, d.UniversityID)
			}
			mu.Unlock()

			wg.Done()
		}(directionID)
	}
	wg.Wait()

	return u.universityService.Set(userID, dto.IDs{IDs: universityIDs})
}
