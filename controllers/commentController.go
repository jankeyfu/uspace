package controllers

import (
	"fmt"

	"strconv"
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
	err := dbo.AddComment(*comment)
	if err != nil {
		fmt.Println(err.Error())
		this.Ctx.WriteString(fialedRes("400", err.Error()))
		return
	}
	fmt.Printf("%+v\n", comment)
	ret := map[string]string{
		"status": "true",
	}
	this.Ctx.WriteString(successRes(ret))
}

func (this *CommentController) GetComment() {
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err != nil {
		this.Ctx.WriteString("请求参数错误")
	}

	comment, err := dbo.GetCommentById(int64(id))
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	this.Ctx.WriteString(fmt.Sprintf("%v", comment))

}

func (this *CommentController) ListComments() {
	comments, err := dbo.ListComments()
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	this.Ctx.WriteString(fmt.Sprintf("%v", comments))
}
