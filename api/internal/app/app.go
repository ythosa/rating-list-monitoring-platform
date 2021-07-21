package app

import (
	"github.com/sirupsen/logrus"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/cache/redis"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/config"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository/postgres"
)

func Run() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})

	cfg := config.Get()

	redisClient := redis.NewClient(cfg.Cache)
	postgresDB, err := postgres.NewDB(cfg.DB)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err)
	}

	cache := redis.NewCache(redisClient)
	repo := postgres.NewRepository(postgresDB)
}
