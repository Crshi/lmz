package controllers

import (
	"encoding/json"

	"github.com/crshi/lmz/models"
)

type ArticleController struct {
	BaseController
}

func (c *ArticleController) Prepaer() {
	//先执行
	c.BaseController.Prepare()
}

func (c *ArticleController) Detail() {

}

func (c *ArticleController) Add() {

}

func (c *ArticleController) Edit() {

}

func (c *ArticleController) Delete() {

}

func (c *ArticleController) Get() {

}

// @Title GetAll
// @Description GetAllArticle
// @Param   body        body    models.User true        "body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [get]
func (c *ArticleController) GetAll() {
	//直接反序化获取json格式的requestbody里的值
	var input models.ArticleQueryInput
	json.Unmarshal(c.Ctx.Input.RequestBody, &input)
	//获取数据列表和总数
	data, total := models.GetAllArticlesByPage(&input)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}
