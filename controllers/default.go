package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"go-cms/models"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "111.me"
	//c.Data["Email"] = "astaxie@gmail.com"

	//session 设置和获取
	c.SetSession("cmsusername", "Jack")
	user := c.GetSession("cmsusername")
	fmt.Println(user)

	//记录日志
	//	logs.Informational("user XXX loged in")

	models.UpdatePage()
	var param = models.GetPage()

	c.Data["Website"] = param.Website
	c.Data["Email"] = param.Email
	c.TplName = "index.tpl"
}
