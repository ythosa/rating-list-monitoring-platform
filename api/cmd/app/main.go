package main

import (
	_ "github.com/lib/pq"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/app"
)

// @title Rating List Monitoring Platform
// @description Rating List Monitoring Platform API
// @version 1.0
// @basePath /api/

// @host localhost:8000/api

// @securityDefinitions.apiKey AccessTokenHeader
// @in header
// @name Authorization

func main() {
	app.Run()
}
