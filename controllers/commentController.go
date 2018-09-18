package controllers

import (
	"fmt"
	"uspace/models"

	"github.com/astaxie/beego"
)

type CommentController struct {
	beego.Controller
}

func (c *CommentController) Get() {
	c.TplName = "index.html"
}

var (
	dbo models.CommentDBO
)

func (this *CommentController) AddComment() {
	comment := new(models.Comment)
	if err := this.ParseForm(comment); err != nil {
		fmt.Print(err.Error())
		this.Ctx.WriteString("add comment failed")
		return
	}
	err := dbo.AddComment(comment)
	if err != nil {
		fmt.Println(err.Error())
		this.Ctx.WriteString(err.Error())
		return
	}
	fmt.Printf("%+v\n", comment)
	this.Ctx.WriteString("add comment success")
}
