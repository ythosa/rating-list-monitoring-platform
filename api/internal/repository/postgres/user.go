package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
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

func (r *User) Create(user rdto.UserCreating) (uint8, error) {
	var id uint8

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

func (r *User) GetUserByID(id uint8) (*models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", usersTable)
	if err := r.db.Get(&user, query, id); err != nil {
		return nil, repository.ErrRecordNotFound
	}

	return &user, nil
}

func (r *User) UpdatePassword(id uint8, password string) error {
	query := fmt.Sprintf("UPDATE %s ut SET password=$1 WHERE ut.id=$2", usersTable)
	if _, err := r.db.Exec(query, password, id); err != nil {
		return repository.ErrRecordNotFound
	}

	return nil
}

func (r *User) PatchUser(id uint8, data rdto.UserPatching) error {
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

func (r *User) GetUsername(id uint8) (*rdto.Username, error) {
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

func (r *User) GetProfile(id uint8) (*rdto.UserProfile, error) {
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
