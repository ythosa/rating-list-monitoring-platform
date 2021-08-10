package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	"github.com/ythosa/rating-list-monitoring-platform-api/internal/app"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/cache/redis"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/delivery/http"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/repository/postgres"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/service"
	"github.com/ythosa/rating-list-monitoring-platform-api/pkg/config"
)

// @title Rating List Monitoring Platform
// @description Rating List Monitoring Platform API
// @version 1.0

// @host localhost:8000/api

// @securityDefinitions.apiKey AccessTokenHeader
// @in header
// @name AuthTokens

func main() {
	configureLogger()

	if err := buildContainer().Invoke(func(app *app.App) { app.Deploy() }); err != nil {
		panic(err)
	}
}

func configureLogger() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.ForceColors = true
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(customFormatter)
}

// nolint:errcheck
func buildContainer() *dig.Container {
	container := dig.New()

	container.Provide(config.Get)
	container.Provide(func() *config.Cache { return config.Get().Cache })
	container.Provide(func() *config.DB { return config.Get().DB })
	container.Provide(func() *config.Server { return config.Get().Server })

	container.Provide(redis.NewClient)
	container.Provide(redis.NewCache)
	container.Provide(postgres.NewDB)
	container.Provide(postgres.NewRepository)
	container.Provide(service.New)
	container.Provide(validator.New)
	container.Provide(http.NewHandler)
	container.Provide(func(cfg *config.Server, handler *http.Handler) *http.Server {
		return http.NewServer(cfg, handler.InitRoutes())
	})

	container.Provide(app.New)

	return container
}
