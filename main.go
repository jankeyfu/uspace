package main

import (
	"net/http"
	"uspace/models"
	_ "uspace/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:9999@/uspace?charset=utf8")
}

func main() {

	models.Init()
	beego.Run()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.ListenAndServe(":8000", nil)
}
