package controllers

import (
	"github.com/astaxie/beego"
	"go-cms/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "111.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	var param = models.GetPage()

	c.Data["Website"] = param.Website
	c.Data["Email"] = param.Email
	c.TplName = "index.tpl"
}
