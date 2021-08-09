package postgres

import (
	"fmt"

	"github.com/ythosa/rating-list-monitoring-platform-api/pkg/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres migrations driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // for migrations
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres database driver

	"github.com/ythosa/rating-list-monitoring-platform-api/internal/repository"
)

const (
	usersTable             = "users"
	universitiesTable      = "universities"
	directionsTable        = "directions"
	usersUniversitiesTable = "users_universities"
	usersDirectionsTable   = "users_directions"
)

func NewDB(cfg *config.DB) (*sqlx.DB, error) {
	db, err := sqlx.Open(cfg.Driver, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode),
	)
	if err != nil {
		return nil, fmt.Errorf("error occurred while opening db connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error occurred while pinging db: %w", err)
	}

	// Create new migrate connections
	m, err := migrate.New(
		fmt.Sprintf("file://%s", cfg.MigrationsPath),
		fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s",
			cfg.Driver, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate new: %w", err)
	}

	_ = m.Up() // Up migrations

	return db, nil
}

func NewRepository(db *sqlx.DB) *repository.Repository {
	return &repository.Repository{
		User:       NewUserImpl(db),
		University: NewUniversityImpl(db),
		Direction:  NewDirectionImpl(db),
	}
}
