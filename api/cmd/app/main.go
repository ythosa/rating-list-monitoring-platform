package main

import (
	_ "github.com/lib/pq"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/app"
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

	//res, err := http.GetForUser("https://cabinet.spbu.ru/Lists/1k_EntryLists/list_94eb54ce-57aa-45e7-9be3-3acd8bf9ca38.html")
	//if err != nil {
	//	println(err)
	//	return
	//}
	//
	//defer res.Body.Close()
	//if res.StatusCode != http.StatusOK {
	//	return
	//}
	//
	//buf := new(strings.Builder)
	//if _, err := io.Copy(buf, res.Body); err != nil {
	//	println(err)
	//
	//	return
	//}
	//
	//ratingReader := res.Body
	//
	//body, err := ioutil.ReadAll(res.Body)
	//println(string(body))
	//
	//service.Spbgu(&ratingReader, "166-912-183 87")
}
