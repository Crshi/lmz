package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/crshi/lmz/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/crshi/lmz/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/crshi/lmz/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/crshi/lmz/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/crshi/lmz/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/crshi/lmz/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
