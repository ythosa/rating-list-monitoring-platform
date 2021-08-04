package main

import (
	_ "github.com/lib/pq"

	"github.com/ythosa/rating-list-monitoring-platform-api/internal/app"
)

// @title Rating List Monitoring Platform
// @description Rating List Monitoring Platform API
// @version 1.0

// @host localhost:8000/api

// @securityDefinitions.apiKey AccessTokenHeader
// @in header
// @name Authorization

func main() {
	app.Run()
}
