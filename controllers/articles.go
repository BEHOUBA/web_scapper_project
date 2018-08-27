package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego/context"
	"github.com/behouba/webScrapperApp/models"
)

func ArticlesController(ctx *context.Context) {
	q := ctx.Input.Query("q")
	jsonData, _ := models.JumiaSearch(1, "", q)
	jsonBs, _ := json.Marshal(jsonData)
	ctx.Output.JSON(string(jsonBs), false, false)
}
