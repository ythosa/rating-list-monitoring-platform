package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/models"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository/rdto"
	"strings"
)

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{db}
}

func (r *User) Create(user rdto.UserCreating) (int, error) {
	var id int

	query := fmt.Sprintf(
		`INSERT INTO %s (nickname, password, first_name, middle_name, last_name, snils) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		usersTable,
	)
	row := r.db.QueryRow(query, user.Nickname, user.Password, user.FirstName, user.MiddleName, user.LastName, user.Snils)
	if err := row.Scan(&id); err != nil {
		logrus.Error(err)

		return 0, repository.ErrUserAlreadyExists
	}

	return id, nil
}

func (r *User) GetUserByNickname(nickname string) (*models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE nickname=$1", usersTable)
	if err := r.db.Get(&user, query, nickname); err != nil {
		return nil, repository.ErrRecordNotFound
	}

	return &user, nil
}

func (r *User) GetUserByID(id int) (*models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", usersTable)
	if err := r.db.Get(&user, query, id); err != nil {
		return nil, repository.ErrRecordNotFound
	}

	return &user, nil
}

func (r *User) UpdatePassword(id int, password string) error {
	query := fmt.Sprintf("UPDATE %s ut SET password=$1 WHERE ut.id=$2", usersTable)
	if _, err := r.db.Exec(query, password, id); err != nil {
		return repository.ErrRecordNotFound
	}

	return nil
}

func (r *User) PatchUser(id int, data rdto.UserPatching) error {
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
		logrus.Error(err)

		return repository.ErrRecordNotFound
	}

	if n, _ := result.RowsAffected(); n == 0 {
		return repository.ErrRecordNotFound
	}

	return nil
}
