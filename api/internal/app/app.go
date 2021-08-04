package app

import (
	"context"
	"os"
	"os/signal"
	"sort"
	"strings"
	"syscall"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"

	"github.com/ythosa/rating-list-monitoring-platform-api/internal/cache/redis"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/config"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/delivery/http"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/repository/postgres"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/service"
)

func Run() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		SortingFunc: func(keys []string) {
			sort.Slice(keys, func(i, j int) bool {
				if keys[j] == "prefix" {
					return false
				}
				if keys[i] == "prefix" {
					return true
				}

				return strings.Compare(keys[i], keys[j]) == -1
			})
		},
	})

	cfg := config.Get()

	redisClient := redis.NewClient(cfg.Cache)

	postgresDB, err := postgres.NewDB(cfg.DB)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err)
	}

	cache := redis.NewCache(redisClient)
	repository := postgres.NewRepository(postgresDB)

	services := service.New(repository, cache)
	validate := validator.New()
	handler := http.NewHandler(services, validate)
	server := http.NewServer(cfg.Server, handler.InitRoutes())

	go func() {
		if err := server.Run(); err != nil {
			logrus.Fatalf("error occurred while running the server: %s", err)
		}
	}()

	logrus.Printf("rating list monitoring platform api is starting on port=%s", cfg.Server.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("rating list monitoring platform api is shutting down...")

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server shutting down: %s", err)
	}

	if err := redisClient.Close(); err != nil {
		logrus.Errorf("error occurred on closing cache connection: %s", err)
	}

	if err := postgresDB.Close(); err != nil {
		logrus.Errorf("error occurred on closing db connection: %s", err)
	}
}
