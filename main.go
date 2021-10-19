package main

import (
	"todoapp/models"
	_ "todoapp/routers"

	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/orm"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "bratix:amaramar@tcp(localhost:3306)/todo")
	orm.RegisterModel(new(models.Todo), new(models.User))
}

func main() {
	orm.RunSyncdb("default", false, false)
	beego.Run()
}
