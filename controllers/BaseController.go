package controllers

import (
	"github.com/astaxie/beego"
	"github.com/crshi/lmz/enums"
	"github.com/crshi/lmz/models"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	//curUser
}

func (c *BaseController) Prepare() {
	//赋值
	c.controllerName, c.actionName = c.GetControllerAndAction()
}

func (c *BaseController) jsonResult(code enums.JsonResultCode, msg string, obj interface{}) {
	r := &models.JsonResult{Code: code, Msg: msg, Obj: obj}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

// 重定向
func (c *BaseController) redirect(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}

// 重定向 去错误页
func (c *BaseController) pageError(msg string) {
	errorurl := c.URLFor("HomeController.Error") + "/" + msg
	c.Redirect(errorurl, 302)
	c.StopRun()
}
