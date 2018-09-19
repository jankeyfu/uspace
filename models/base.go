package models

import (
	"log"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	maxIdle = 30
	maxConn = 30
)

var (
	O orm.Ormer
	L *log.Logger
)

func Init() {
	orm.Debug = true
	error := orm.RegisterDriver("mysql", orm.DRMySQL)
	if error != nil {
		log.Fatal(error.Error())
	}
	error = orm.RegisterDataBase("default", "mysql", "root:9999@/uspace?charset=utf8", maxIdle, maxConn)
	if error != nil {
		log.Fatal(error.Error())
	}
	orm.RegisterModel(new(Comment))
	L = logs.GetLogger("ORM")
}
