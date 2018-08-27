package routers

import (
	"github.com/astaxie/beego"
	"github.com/behouba/webScrapperApp/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Get("/articles/", controllers.ArticlesController)
}
