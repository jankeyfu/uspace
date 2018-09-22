package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["tmp"] = "template  \"header.html\""
	c.Data["tplName"] = "header.html"
	c.TplName = "index.html"
}
