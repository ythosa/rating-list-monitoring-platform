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
	d, _ := ratingparser.LETI(
		"https://etu.ru/ru/abiturientam/priyom-na-1-y-kurs/podavshie-zayavlenie/ochnaya/byudzhet/programmnaya-inzheneriya",
		"166-912-183 87")

	println("Budget places:", d.BudgetPlaces)
	println("Score:", d.Score)
	println("Position:", d.Position)
	println("P1 upper:", d.PriorityOneUpper)
}
