package main

import (
	_ "github.com/lib/pq"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/pkg/ratingparser"
)

// @title Rating List Monitoring Platform
// @description Rating List Monitoring Platform API
// @version 1.0

// @host localhost:8000/api

// @securityDefinitions.apiKey AccessTokenHeader
// @in header
// @name Authorization

func main() {
	//app.Run()
	d, _ := ratingparser.SPBGU(
		"https://cabinet.spbu.ru/Lists/1k_EntryLists/list_bb827177-7802-419a-91af-bfbb50849975.html",
		"167-174-212 72")

	println("Budget places:", d.BudgetPlaces)
	println("Score:", d.Score)
	println("Position:", d.Position)
	println("P1 upper:", d.PriorityOneUpper)
}
