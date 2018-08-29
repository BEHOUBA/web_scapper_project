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

	req := make(chan int)
	goroutineTracker := 0

	go func() {
		jumiaData, err := models.JumiaSearch(1, "", q)
		if err != nil {
			log.Println(err)
			// ctx.Output.SetStatus(http.StatusInternalServerError)
			return
		}
		allData = append(allData, jumiaData...)
		goroutineTracker++
		req <- goroutineTracker
	}()

	go func() {
		afData, err := models.AfrimarketSearch(0, "", q)
		if err != nil {
			log.Println(err)
			return
		}
		allData = append(allData, afData...)
		goroutineTracker++
		req <- goroutineTracker
	}()

	for i := 0; i < 2; i++ {
		if 2 == <-req {
			jsonBs, err := json.Marshal(allData)
			if err != nil {
				log.Println(err)
				// ctx.Output.SetStatus(http.StatusInternalServerError)
				return
			}
			ctx.Output.JSON(string(jsonBs), false, false)
			log.Println(<-req)
		}
	}
}
