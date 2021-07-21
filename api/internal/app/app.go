package app

import (
	"github.com/sirupsen/logrus"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/cache/redis"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/config"
)

func Run() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})

	cfg := config.Get()

	redisCache := redis.NewCache(cfg.Cache)
}
