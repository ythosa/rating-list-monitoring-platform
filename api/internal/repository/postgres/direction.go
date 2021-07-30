package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/logging"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository/rdto"
)

type Direction struct {
	db     *sqlx.DB
	logger *logging.Logger
}

func NewDirection(db *sqlx.DB) *Direction {
	return &Direction{
		db:     db,
		logger: logging.NewLogger("direction repository"),
	}
}

func (r *Direction) GetAll() ([]rdto.Direction, error) {
	var directions []rdto.Direction

	query := fmt.Sprintf(
		`SELECT d.id as direction_id, d.name as direction_name, 
					un.id as university_id, un.name as university_name FROM %s d 
			INNER JOIN %s un on d.university_id = un.id`,
		directionsTable, universitiesTable,
	)
	err := r.db.Select(&directions, query)

	return directions, err
}

func (r *Direction) GetByID(id uint) (rdto.Direction, error) {
	var direction rdto.Direction

	query := fmt.Sprintf(
		`SELECT d.id as direction_id, d.name as direction_name, 
					un.id as university_id, un.name as university_name FROM %s d 
			INNER JOIN %s un on d.university_id = un.id
			WHERE d.id = $1`,
		directionsTable, universitiesTable,
	)
	err := r.db.Get(&direction, query, id)

	return direction, err
}

func (r *Direction) Get(userID uint) ([]rdto.Direction, error) {
	var directions []rdto.Direction

	query := fmt.Sprintf(
		`SELECT d.id as direction_id, d.name as direction_name, d.url as direction_url,
					un.id as university_id, un.name as university_name FROM %s d 
			INNER JOIN %s ud on d.id = ud.direction_id
			INNER JOIN %s un on d.university_id = un.id
			WHERE ud.user_id = $1`,
		directionsTable, usersDirectionsTable, universitiesTable,
	)
	err := r.db.Select(&directions, query, userID)

	return directions, err
}

func (r *Direction) Set(userID uint, directionIDs dto.IDs) error {
	tx, err := r.db.Begin()
	if err != nil {
		r.logger.Error(err)

		return err
	}

	query := fmt.Sprintf("INSERT INTO %s (user_id, direction_id) VALUES ($1, $2)", usersDirectionsTable)
	for _, directionID := range directionIDs.IDs {
		if _, err := tx.Exec(query, userID, directionID); err != nil {
			r.logger.Error(err)
			tx.Rollback()

			return err
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()

		return err
	}

	return nil
}

func (r *Direction) Clear(userID uint) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", usersDirectionsTable)
	_, err := r.db.Exec(query, userID)

	return err
}
