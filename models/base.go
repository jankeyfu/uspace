package models

import (
	"github.com/astaxie/beego/orm"
)

const (
	maxIdle = 30
	maxConn = 30
)

func Init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:9999@/uspace?charset=utf8", maxIdle, maxConn)
	orm.RegisterModel(new(Comment))
}
