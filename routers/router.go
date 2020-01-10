// @APIVersion 1.0.0
// @Title LMZ API
// @Swagger API.
// @Contact 1078649116@qq.com
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/crshi/lmz/controllers"
)

func init() {

	//运行跨域请求
	//在http请求的响应流头部加上如下信息
	//rw.Header().Set("Access-Control-Allow-Origin", "*")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	//自动化文档
	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/api",
				beego.NSInclude(
					&controllers.MainController{},
					&controllers.ArticleController{},
				),
			),
		)
	beego.AddNamespace(ns)
	beego.SetStaticPath("/swagger", "swagger")
	beego.Router("/", &controllers.MainController{})
	beego.Router("/article", &controllers.ArticleController{})
}
