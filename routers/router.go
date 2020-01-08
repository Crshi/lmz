package routers

import (
	"github.com/astaxie/beego"
	"github.com/crshi/lmz/controllers"
)

func init() {
	beego.SetStaticPath("/swagger", "swagger")

	beego.Router("/", &controllers.MainController{})
}
