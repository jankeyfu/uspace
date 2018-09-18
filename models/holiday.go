package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Holiday))
}

type Holiday struct {
	Id          int       `form:"-"`
	Holiday     time.Time `form:"holiday"`
	Description string    `form:"description"`
	UserId      int       `form:"user_id"`
	Type        int       `form:"type"`
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time `orm:"auto_now;type(datetime)"`
}

func (Holiday) TableName() string {
	return "t_holiday"
}

func AddHoliday(holiday *Holiday) (int64, error) {
	o := orm.NewOrm()
	return o.Insert(holiday)
}

func GetHolidays(userId int) ([]Holiday, error) {
	o := orm.NewOrm()
	var holidays []Holiday
	_, err := o.QueryTable("t_holiday").Filter("user_id", userId).All(&holidays)
	return holidays, err
}
