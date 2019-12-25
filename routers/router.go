package routers

import (
	"github.com/astaxie/beego"
	"github.com/crshi/lmz/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
