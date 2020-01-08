package controllers

type ArticleController struct {
	BaseController
}

func (c *ArticleController) Prepaer() {
	//先执行
	c.BaseController.Prepare()
}
