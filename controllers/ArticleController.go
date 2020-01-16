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

// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (c *ArticleController) Update() {

}

func (c *ArticleController) Delete() {

}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (c *ArticleController) Get() {

}

// @Title GetAll
// @Description find all artical by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.ArticleQueryInput
// @Failure 403 {object} is empty
// @router /:objectId [get]
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
