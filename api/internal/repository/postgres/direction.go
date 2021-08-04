package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/ythosa/rating-list-monitoring-platform-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/logging"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/models"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/repository/rdto"
)

type DirectionImpl struct {
	db     *sqlx.DB
	logger *logging.Logger
}

func NewDirectionImpl(db *sqlx.DB) *DirectionImpl {
	return &DirectionImpl{
		db:     db,
		logger: logging.NewLogger("direction repository"),
	}
}

func (r *DirectionImpl) GetAll() ([]rdto.Direction, error) {
	var directions []rdto.Direction

	query := fmt.Sprintf(
		`SELECT d.id as direction_id, d.name as direction_name, 
					un.id as university_id, un.name as university_name FROM %s d 
			INNER JOIN %s un on d.university_id = un.id`,
		directionsTable, universitiesTable,
	)
	if err := r.db.Select(&directions, query); err != nil {
		return nil, fmt.Errorf("error while getting all directions: %w", err)
	}

	return directions, nil
}

func (r *DirectionImpl) GetByID(id uint) (*models.Direction, error) {
	var direction models.Direction

	query := fmt.Sprintf(`SELECT * FROM %s d WHERE d.id = $1`, directionsTable)
	if err := r.db.Get(&direction, query, id); err != nil {
		return nil, fmt.Errorf("error while getting direction by id: %w", err)
	}

	return &direction, nil
}

func (r *DirectionImpl) GetUniversityID(id uint) (*rdto.UniversityID, error) {
	var universityID rdto.UniversityID

	query := fmt.Sprintf(
		`SELECT un.id as university_id FROM %s d INNER JOIN %s un on d.university_id = un.id WHERE d.id = $1`,
		directionsTable, universitiesTable,
	)
	if err := r.db.Get(&universityID, query, id); err != nil {
		return nil, fmt.Errorf("error while getting university by id: %w", err)
	}

	return &universityID, nil
}

func (r *DirectionImpl) GetForUser(userID uint) ([]rdto.Direction, error) {
	var directions []rdto.Direction

	query := fmt.Sprintf(
		`SELECT d.id as direction_id, d.name as direction_name, d.url as direction_url,
					un.id as university_id, un.name as university_name FROM %s d 
			INNER JOIN %s ud on d.id = ud.direction_id
			INNER JOIN %s un on d.university_id = un.id
			WHERE ud.user_id = $1`,
		directionsTable, usersDirectionsTable, universitiesTable,
	)
	if err := r.db.Select(&directions, query, userID); err != nil {
		return nil, fmt.Errorf("error while getting user directions: %w", err)
	}

	return directions, nil
}

func (r *DirectionImpl) SetForUser(userID uint, directionIDs dto.IDs) error {
	tx, err := r.db.Begin()
	if err != nil {
		r.logger.Error(err)

		return fmt.Errorf("error while beginning transaction: %w", err)
	}

	query := fmt.Sprintf("INSERT INTO %s (user_id, direction_id) VALUES ($1, $2)", usersDirectionsTable)
	for _, directionID := range directionIDs.IDs {
		if _, err := tx.Exec(query, userID, directionID); err != nil {
			r.logger.Error(err)

			if err := tx.Rollback(); err != nil {
				return fmt.Errorf("error while rollbacking transaction: %w", err)
			}

			return fmt.Errorf("error while adding directions to user: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		r.logger.Error(err)

		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("error while rollbacking transaction: %w", err)
		}

		return fmt.Errorf("error while committing transaction: %w", err)
	}

	return nil
}

func (r *DirectionImpl) Clear(userID uint) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", usersDirectionsTable)
	if _, err := r.db.Exec(query, userID); err != nil {
		return fmt.Errorf("error while deleting user directions: %w", err)
	}

	return nil
}
