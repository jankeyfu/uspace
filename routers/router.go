package routers

import (
	"uspace/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.AutoRouter(&controllers.UserController{})

	beego.Router("/holiday", &controllers.HolidayController{})
	beego.Router("/holiday/add", &controllers.HolidayController{}, "*:AddHoliday")
	beego.Router("/holiday/get", &controllers.HolidayController{}, "*:GetAllHoliday")

	beego.Router("/comment/add", &controllers.CommentController{}, "*:AddComment")
	beego.Router("/comment/get/:id", &controllers.CommentController{}, "*:GetComment")
	beego.Router("/comment/list", &controllers.CommentController{}, "*:ListComments")

}
