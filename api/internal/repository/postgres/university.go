package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/logging"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/models"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository/rdto"
)

type UniversityImpl struct {
	db     *sqlx.DB
	logger *logging.Logger
}

func NewUniversityImpl(db *sqlx.DB) *UniversityImpl {
	return &UniversityImpl{
		db:     db,
		logger: logging.NewLogger("university repository"),
	}
}

func (r *UniversityImpl) GetAll() ([]rdto.University, error) {
	var universities []rdto.University
	err := r.db.Select(&universities, fmt.Sprintf("SELECT id, name FROM %s", universitiesTable))

	return universities, err
}

func (r *UniversityImpl) GetForUser(userID uint) ([]rdto.University, error) {
	var universities []rdto.University

	query := fmt.Sprintf(
		"SELECT un.id, un.name FROM %s un INNER JOIN %s uu on un.id = uu.university_id WHERE uu.user_id = $1",
		universitiesTable, usersUniversitiesTable,
	)
	err := r.db.Select(&universities, query, userID)

	return universities, err
}

func (r *UniversityImpl) GetByID(id uint) (*models.University, error) {
	var university models.University

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", universitiesTable)
	if err := r.db.Get(&university, query, id); err != nil {
		return nil, err
	}

	return &university, nil
}

func (r *UniversityImpl) Set(userID uint, universityIDs dto.IDs) error {
	tx, err := r.db.Begin()
	if err != nil {
		r.logger.Error(err)

		return err
	}

	query := fmt.Sprintf("INSERT INTO %s (user_id, university_id) VALUES ($1, $2)", usersUniversitiesTable)
	for _, universityID := range universityIDs.IDs {
		if _, err := tx.Exec(query, userID, universityID); err != nil {
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

func (r *UniversityImpl) Clear(userID uint) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", usersUniversitiesTable)
	_, err := r.db.Exec(query, userID)

	return err
}
