package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	ID        int64 `orm:"auto"`
	Username  string
	Password  string
	Phone     string
	Email     string
	City      string
	Birthday  time.Time `orm:"default(1970-01-01 00:00:00)"`
	Sex       string
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
	Actived   uint      `orm:"default(0)"`
}

func (User) TableName() string {
	return "t_users"
}

func (user User) String() string {
	return fmt.Sprintf("ID:%d Username:%s Password:%s Phone:%s Email:%s City:%s", user.ID, user.Username, user.Password, user.Phone, user.Email, user.City)
}
func AddUser(user *User) (bool, error) {
	o := orm.NewOrm()
	fmt.Print(o.Insert(user))
	return true, nil
}

func GetUserById(id int64) *User {
	o := orm.NewOrm()
	user := User{ID: id}
	o.Read(&user)
	return &user
}

func VerifyUser(username, password string) (bool, error) {
	o := orm.NewOrm()
	user := User{Username: username, Password: password}
	err := o.Read(&user)
	if err != nil {
		return false, err
	}
	return true, nil
}
