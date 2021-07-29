package postgres

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/config"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository"
	"gopkg.in/errgo.v2/fmt/errors"
)

const (
	usersTable             = "users"
	universitiesTable      = "universities"
	directionsTable        = "directions"
	usersUniversitiesTable = "users_universities"
	usersDirectionsTable   = "users_directions"
)

func NewDB(cfg *config.DB) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		cfg.Driver,
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode,
		),
	)
	if err != nil {
		return nil, errors.Newf("error occurred while opening db connection: %s", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Newf("error occurred while pinging db: %s", err)
	}

	// Create new migrate connections
	m, err := migrate.New(
		fmt.Sprintf("file://%s", cfg.MigrationsPath),
		fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s",
			cfg.Driver, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode),
	)
	if err != nil {
		return nil, errors.Newf("failed to migrate new: %s", err)
	}

	_ = m.Up() // Up migrations

	return db, nil
}

func NewRepository(db *sqlx.DB) *repository.Repository {
	return &repository.Repository{
		User:       NewUser(db),
		University: NewUniversity(db),
		Direction:  NewDirection(db),
	}
}
