package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/logging"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/models"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository/rdto"
	"strings"
)

type User struct {
	db     *sqlx.DB
	logger *logging.Logger
}

func NewUser(db *sqlx.DB) *User {
	return &User{
		db:     db,
		logger: logging.NewLogger("user repository"),
	}
}

func (r *User) Create(user rdto.UserCreating) (uint, error) {
	var id uint

	query := fmt.Sprintf(
		`INSERT INTO %s (username, password, first_name, middle_name, last_name, snils) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		usersTable,
	)
	row := r.db.QueryRow(query, user.Username, user.Password, user.FirstName, user.MiddleName, user.LastName, user.Snils)
	if err := row.Scan(&id); err != nil {
		r.logger.Error(err)

		return 0, repository.ErrUserAlreadyExists
	}

	return id, nil
}

func (r *User) GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1", usersTable)
	if err := r.db.Get(&user, query, username); err != nil {
		return nil, repository.ErrRecordNotFound
	}

	return &user, nil
}

func (r *User) GetUserByID(id uint) (*models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", usersTable)
	if err := r.db.Get(&user, query, id); err != nil {
		return nil, repository.ErrRecordNotFound
	}

	return &user, nil
}

func (r *User) UpdatePassword(id uint, password string) error {
	query := fmt.Sprintf("UPDATE %s ut SET password=$1 WHERE ut.id=$2", usersTable)
	if _, err := r.db.Exec(query, password, id); err != nil {
		return repository.ErrRecordNotFound
	}

	return nil
}

func (r *User) PatchUser(id uint, data rdto.UserPatching) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if data.FirstName != nil {
		setValues = append(setValues, fmt.Sprintf("first_name=$%d", argID))
		args = append(args, *data.FirstName)
		argID++
	}

	if data.MiddleName != nil {
		setValues = append(setValues, fmt.Sprintf("middle_name=$%d", argID))
		args = append(args, *data.MiddleName)
		argID++
	}

	if data.LastName != nil {
		setValues = append(setValues, fmt.Sprintf("last_name=$%d", argID))
		args = append(args, *data.LastName)
		argID++
	}

	if data.Snils != nil {
		setValues = append(setValues, fmt.Sprintf("snils=$%d", argID))
		args = append(args, *data.Snils)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s ut SET %s WHERE ut.id=$%d", usersTable, setQuery, argID)
	args = append(args, id)

	result, err := r.db.Exec(query, args...)
	if err != nil {
		r.logger.Error(err)

		return repository.ErrRecordNotFound
	}

	if n, _ := result.RowsAffected(); n == 0 {
		return repository.ErrRecordNotFound
	}

	return nil
}

func (r *User) GetUsername(id uint) (*rdto.Username, error) {
	var username rdto.Username

	query := fmt.Sprintf(
		"SELECT (username) FROM %s WHERE id=$1",
		usersTable,
	)
	if err := r.db.Get(&username, query, id); err != nil {
		r.logger.Error(err)

		return nil, repository.ErrRecordNotFound
	}

	return &username, nil
}

func (r *User) GetProfile(id uint) (*rdto.UserProfile, error) {
	var userProfile rdto.UserProfile

	query := fmt.Sprintf(
		"SELECT username, first_name, middle_name, last_name, snils FROM %s WHERE id=$1",
		usersTable,
	)
	if err := r.db.Get(&userProfile, query, id); err != nil {
		r.logger.Error(err)

		return nil, repository.ErrRecordNotFound
	}

	return &userProfile, nil
}

func (r *User) SetUniversities(id uint, universityIDs dto.IDs) error {
	tx, err := r.db.Begin()
	if err != nil {
		r.logger.Error(err)

		return err
	}

	query := fmt.Sprintf("INSERT INTO %s (user_id, university_id) VALUES ($1, $2)", usersUniversitiesTable)
	for _, universityID := range universityIDs.IDs {
		if _, err := tx.Exec(query, id, universityID); err != nil {
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

func (r *User) GetUniversities(id uint) ([]rdto.University, error) {
	var universities []rdto.University

	query := fmt.Sprintf(
		"SELECT un.id, un.name FROM %s un INNER JOIN %s uu on un.id = uu.university_id WHERE uu.user_id = $1",
		universitiesTable, usersUniversitiesTable,
	)
	err := r.db.Select(&universities, query, id)

	return universities, err
}

func (r *User) ClearUniversities(id uint) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", usersUniversitiesTable)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *User) SetDirections(id uint, directionIDs dto.IDs) error {
	tx, err := r.db.Begin()
	if err != nil {
		r.logger.Error(err)

		return err
	}

	query := fmt.Sprintf("INSERT INTO %s (user_id, direction_id) VALUES ($1, $2)", usersDirectionsTable)
	for _, directionID := range directionIDs.IDs {
		if _, err := tx.Exec(query, id, directionID); err != nil {
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

func (r *User) GetDirections(id uint) ([]rdto.Direction, error) {
	var directions []rdto.Direction

	query := fmt.Sprintf(
		`SELECT d.id as direction_id, d.name as direction_name, 
					un.id as university_id, un.name as university_name FROM %s d 
			INNER JOIN %s ud on d.id = ud.direction_id
			INNER JOIN %s un on d.university_id = un.id
			WHERE ud.user_id = $1`,
		directionsTable, usersDirectionsTable, universitiesTable,
	)
	err := r.db.Select(&directions, query, id)

	return directions, err
}

func (r *User) ClearDirections(id uint) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", usersDirectionsTable)
	_, err := r.db.Exec(query, id)

	return err
}
