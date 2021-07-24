package app

import (
	"github.com/sirupsen/logrus"
	"sort"
	"strings"
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

	//cfg := config.Get()
	//
	//redisClient := redis.NewClient(cfg.Cache)
	//postgresDB, err := postgres.NewDB(cfg.DB)
	//if err != nil {
	//	logrus.Fatalf("failed to initialize db: %s", err)
	//}
	//
	//cache := redis.NewCache(redisClient)
	//repo := postgres.NewRepository(postgresDB)
}
