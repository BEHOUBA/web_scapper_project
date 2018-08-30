package controllers

import (
	"encoding/json"
	"log"

	"github.com/astaxie/beego/context"
	"github.com/behouba/webScrapperApp/models"
)

// ArticlesController handler function handle
// api call for search request from client
func ArticlesController(ctx *context.Context) {
	q := ctx.Input.Query("q")
	var allData []models.Product

	dataChan := make(chan []models.Product)

	go func() {
		jumiaData, err := models.JumiaSearch(1, "", q)
		if err != nil {
			log.Println(err)
			// ctx.Output.SetStatus(http.StatusInternalServerError)
			return
		}
		dataChan <- jumiaData
	}()

	go func() {
		afData, err := models.AfrimarketSearch(0, "", q)
		if err != nil {
			log.Println(err)
			return
		}
		dataChan <- afData
	}()

	go func() {
		afData, err := models.YaatooSearch(1, q)
		if err != nil {
			log.Println(err)
			return
		}
		dataChan <- afData
	}()

	for i := 0; i < 3; i++ {

		allData = append(allData, <-dataChan...)
		if 2 == i {
			close(dataChan)
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
