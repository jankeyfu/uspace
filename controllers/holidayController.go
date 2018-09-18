package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"uspace/models"

	"github.com/astaxie/beego"
)

type HolidayController struct {
	beego.Controller
}

func (this *HolidayController) Get() {
	this.TplName = "index.html"
}

/*
	Id          int
	Holiday     time.Time
	Description string
	UserId      int
	Type        int
	CreatedAt   time.Time `orm:"-"`
	UpdatedAt   time.Time `orm:"-"
*/
func (this *HolidayController) AddHoliday() {
	holiday := models.Holiday{}
	if err := this.ParseForm(&holiday); err != nil {
		//handle error
		fmt.Print(err)
	}
	fmt.Printf("%+v\n", holiday)
	// this.Ctx.WriteString(holiday)
	models.AddHoliday(&holiday)
	this.Ctx.WriteString("添加成功")
}

func (this *HolidayController) GetAllHoliday() {
	userId := this.GetString("user_id")
	id, _ := strconv.Atoi(userId)
	holidays, _ := models.GetHolidays(id)
	// fmt.Fprintf(this.Ctx, "%+v", holidays)
	res, _ := json.Marshal(holidays)
	this.Ctx.WriteString(string(res))
}
