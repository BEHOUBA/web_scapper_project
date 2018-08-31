package controllers

import (
	"encoding/json"
	"log"
	"sort"

	"github.com/astaxie/beego/context"
	"github.com/behouba/webScrapperApp/models"
)

// ArticlesController handler function handle
// api call for search request from client
func ArticlesController(ctx *context.Context) {
	q := ctx.Input.Query("q")
	var allData []models.Product
	var reqNumber = 6
	dataChan := make(chan []models.Product)

	go func() {
		jumiaData, err := models.AllFromJumia(q)
		if err != nil {
			log.Println(err)
			// ctx.Output.SetStatus(http.StatusInternalServerError)
		}
		dataChan <- jumiaData
	}()

	go func() {
		afData, err := models.AllFromAfrimarket(q)
		if err != nil {
			log.Println(err)
		}
		dataChan <- afData
	}()

	// go func() {
	// 	afData, err := models.AllFromYaatoo(q)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	dataChan <- afData
	// }()

	go func() {
		afData, err := models.BabikenSearch(q)
		if err != nil {
			log.Println(err)
		}
		dataChan <- afData
	}()
	go func() {
		afData, err := models.AllFromSitcom(q)
		if err != nil {
			log.Println(err)
		}
		dataChan <- afData
	}()

	go func() {
		afData, err := models.AllFromAfrikdiscount(q)
		if err != nil {
			log.Println(err)
		}
		dataChan <- afData
	}()

	go func() {
		afData, err := models.AllFromPdastoreci(q)
		if err != nil {
			log.Println(err)
		}
		dataChan <- afData
	}()

	for i := 1; i <= reqNumber; i++ {
		allData = append(allData, <-dataChan...)
		if reqNumber == i {
			close(dataChan)
			sort.Slice(allData, func(i, j int) bool {
				return allData[i].Price < allData[j].Price
			})
			jsonBs, err := json.Marshal(allData)
			if err != nil {
				log.Println(err)
				// ctx.Output.SetStatus(http.StatusInternalServerError)
				return
			}
			ctx.Output.JSON(string(jsonBs), false, false)
		}
	}
}
