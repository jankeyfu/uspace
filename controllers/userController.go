package controllers

import (
	"fmt"
	"uspace/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) AddUser() {
	user := models.User{}
	if err := this.ParseForm(&user); err != nil {
		//handle error
		fmt.Print(err)
	}
	fmt.Printf("%+v\n", user)
	// this.Ctx.WriteString(holiday)
	models.AddUser(&user)
	this.Ctx.WriteString("添加成功")
}
