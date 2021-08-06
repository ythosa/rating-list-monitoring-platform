package service

import (
	"errors"
	"fmt"
	"sort"
	"sync"

	"github.com/ythosa/rating-list-monitoring-platform-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/logging"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/models"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/repository"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/repository/rdto"
)

type DirectionImpl struct {
	directionRepository repository.Direction
	userRepository      repository.User
	universityService   University
	parsingService      Parsing
	logger              *logging.Logger
}

func NewDirectionImpl(
	directionRepository repository.Direction,
	userRepository repository.User,
	universityService University,
	parsingService Parsing,
) *DirectionImpl {
	return &DirectionImpl{
		directionRepository: directionRepository,
		userRepository:      userRepository,
		universityService:   universityService,
		parsingService:      parsingService,
		logger:              logging.NewLogger("directions service"),
	}
}

func (u *DirectionImpl) GetByID(id uint) (*models.Direction, error) {
	direction, err := u.directionRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("error while getting direction by id by repository: %w", err)
	}

	return direction, nil
}

func (u *DirectionImpl) GetAll() ([]dto.UniversityDirections, error) {
	directions, err := u.directionRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error while getting all directions by repository: %w", err)
	}

	return u.mapDirectionsToUniversityDirections(directions), nil
}

func (u *DirectionImpl) GetForUser(userID uint) ([]dto.UniversityDirections, error) {
	directions, err := u.directionRepository.GetForUser(userID)
	if err != nil {
		return nil, fmt.Errorf("error while getting user directions by repository: %w", err)
	}

	return u.mapDirectionsToUniversityDirections(directions), nil
}

func (u *DirectionImpl) mapDirectionsToUniversityDirections(directions []rdto.Direction) []dto.UniversityDirections {
	universityDirections := make([]dto.UniversityDirections, 0)

	for _, d := range directions {
		isExists := false

		for i, ud := range universityDirections {
			if ud.UniversityID == d.UniversityID {
				universityDirections[i].Directions = append(
					universityDirections[i].Directions,
					dto.Direction{ID: d.DirectionID, Name: d.DirectionName},
				)
				isExists = true

				break
			}
		}

		if !isExists {
			universityDirections = append(universityDirections, dto.UniversityDirections{
				UniversityID:       d.UniversityID,
				UniversityName:     d.UniversityName,
				UniversityFullName: d.UniversityFullName,
				Directions:         []dto.Direction{{ID: d.DirectionID, Name: d.DirectionName}},
			})
		}
	}

	sort.SliceStable(universityDirections, func(i, j int) bool {
		return universityDirections[i].UniversityID < universityDirections[j].UniversityID
	})

	return universityDirections
}

type parsingDirectionResults struct {
	directionsWithRating chan dto.DirectionWithParsingResult
	errors               chan error
	wg                   *sync.WaitGroup
}

func newParsingDirectionResults(directionsCount int) *parsingDirectionResults {
	return &parsingDirectionResults{
		directionsWithRating: make(chan dto.DirectionWithParsingResult, directionsCount),
		errors:               make(chan error, directionsCount),
		wg:                   &sync.WaitGroup{},
	}
}

func (u *DirectionImpl) GetForUserWithRating(userID uint) ([]dto.UniversityDirectionsWithRating, error) {
	directions, err := u.directionRepository.GetForUser(userID)
	if err != nil {
		u.logger.Error(err)

		return nil, fmt.Errorf("error while getting user directions by repository: %w", err)
	}

	userSnils, err := u.userRepository.GetSnils(userID)
	if err != nil {
		u.logger.Error(err)

		return nil, fmt.Errorf("error while getting user snils by repository: %w", err)
	}

	results := newParsingDirectionResults(len(directions))
	for _, d := range directions {
		results.wg.Add(1)

		go u.parseDirectionRating(results, d, userSnils.Snils)
	}

	results.wg.Wait()

	if len(results.errors) != 0 {
		return nil, fmt.Errorf("error while waiting parsing rating: %w", <-results.errors)
	}

	directionsWithRating := make([]dto.DirectionWithParsingResult, len(directions))
	for i := 0; i < len(directions); i++ {
		directionsWithRating[i] = <-results.directionsWithRating
	}

	return u.mapDirectionsToUniversityDirectionsWithRating(directionsWithRating), nil
}

func (u *DirectionImpl) parseDirectionRating(
	results *parsingDirectionResults,
	direction rdto.Direction,
	userSnils string,
) {
	defer results.wg.Done()

	parsingResult, err := u.parsingService.ParseRating(
		direction.UniversityName,
		direction.DirectionURL,
		userSnils,
	)
	if err != nil && errors.Is(err, ErrUserNotFoundInRatingList) {
		parsingResult = &dto.EmptyParsingResult
	} else if err != nil {
		results.errors <- fmt.Errorf("error while parsong rating list: %w", err)

		return
	}

	results.directionsWithRating <- dto.DirectionWithParsingResult{
		Direction:     direction,
		ParsingResult: *parsingResult,
	}
}

func (u *DirectionImpl) mapDirectionsToUniversityDirectionsWithRating(
	directions []dto.DirectionWithParsingResult,
) []dto.UniversityDirectionsWithRating {
	universityDirectionsWithRating := make([]dto.UniversityDirectionsWithRating, 0)

	for _, d := range directions {
		isExists := false

		for i, ud := range universityDirectionsWithRating {
			if ud.UniversityID == d.Direction.UniversityID {
				universityDirectionsWithRating[i].Directions = append(
					universityDirectionsWithRating[i].Directions,
					dto.NewDirectionWithRating(d),
				)
				isExists = true

				break
			}
		}

		if !isExists {
			universityDirectionsWithRating = append(universityDirectionsWithRating, dto.UniversityDirectionsWithRating{
				UniversityID:       d.Direction.UniversityID,
				UniversityName:     d.Direction.UniversityName,
				UniversityFullName: d.Direction.UniversityFullName,
				Directions:         []dto.DirectionWithRating{dto.NewDirectionWithRating(d)},
			})
		}
	}

	sort.SliceStable(universityDirectionsWithRating, func(i, j int) bool {
		return universityDirectionsWithRating[i].UniversityID < universityDirectionsWithRating[j].UniversityID
	})

	return universityDirectionsWithRating
}

func (u *DirectionImpl) SetForUser(userID uint, directionIDs dto.IDs) error {
	err := u.directionRepository.Clear(userID)
	if err != nil {
		return fmt.Errorf("error while clearing user directions by repository: %w", err)
	}

	if err := u.directionRepository.SetForUser(userID, directionIDs); err != nil {
		return fmt.Errorf("error while setting directions for user by repository: %w", err)
	}

	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)

	universityIDs := make([]uint, 0)

	for _, directionID := range directionIDs.IDs {
		wg.Add(1)

		go func(id uint) {
			d, _ := u.directionRepository.GetUniversityID(id)

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

	if err := u.universityService.SetForUser(userID, dto.IDs{IDs: universityIDs}); err != nil {
		return fmt.Errorf("error while updating user universities by repository: %w", err)
	}

	return nil
}
