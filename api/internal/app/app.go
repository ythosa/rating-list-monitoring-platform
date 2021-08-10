package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"

	"github.com/ythosa/rating-list-monitoring-platform-api/pkg/config"

	"github.com/sirupsen/logrus"

	"github.com/ythosa/rating-list-monitoring-platform-api/internal/delivery/http"
)

type App struct {
	server      *http.Server
	redisClient *redis.Client
	postgresDB  *sqlx.DB
	cfg         *config.Config
}

func New(server *http.Server, redisClient *redis.Client, postgresDB *sqlx.DB, cfg *config.Config) *App {
	return &App{
		server:      server,
		redisClient: redisClient,
		postgresDB:  postgresDB,
		cfg:         cfg,
	}
}

func (a *App) Deploy() {
	go func() {
		if err := a.server.Run(); err != nil {
			logrus.Fatalf("error occurred while running the server: %s", err)
		}
	}()

	logrus.Infof("rating list monitoring platform api is starting on port=%s", a.cfg.Server.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("rating list monitoring platform api is shutting down...")

	if err := a.server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server shutting down: %s", err)
	}

	if err := a.redisClient.Close(); err != nil {
		logrus.Errorf("error occurred on closing cache connection: %s", err)
	}

	if err := a.postgresDB.Close(); err != nil {
		logrus.Errorf("error occurred on closing db connection: %s", err)
	}
}
