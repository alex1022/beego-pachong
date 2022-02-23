package routers

import (
	"github.com/astaxie/beego"
	"web/controllers"
)

func init() {
	beego.Router("/a", &controllers.MainController{})
	beego.Router("/123", &controllers.CrawlMovieController{}, "*:CrawlMovie")
	//rootpath路径随便取
}
