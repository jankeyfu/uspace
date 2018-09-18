package main

import (
	"fmt"
	"uspace/models"
	_ "uspace/routers"

	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Init()
	orm.Debug = true
	models.Init()
	// StaticDir["/static"] = "static"
	// beego.Run()

	o := orm.NewOrm()
	comment := models.Comment{Id: 1}
	o.Read(&comment)
	fmt.Println(comment)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	// http.ListenAndServe(":8000", nil)
}
