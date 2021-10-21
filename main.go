package main

import (
	"todoapp/global"
	"todoapp/models"
	_ "todoapp/routers"

	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/orm"

	_ "github.com/go-sql-driver/mysql"
)

/* initialize function */
func init() {
	/* Orm driver */
	orm.RegisterDriver("mysql", orm.DRMySQL)
	/* Database name, driver and connection string */
	orm.RegisterDataBase("default", "mysql", "bratix:amaramar@tcp(localhost:3306)/todo")
	/* Registering models */
	orm.RegisterModel(new(models.Todo), new(models.User))
}

func main() {
	/* Connect to redis */
	global.CreateRedisConnection()
	defer global.Redisclient.Close()
	/* Sync db */
	orm.RunSyncdb("default", false, false)
	/* Run beego server */
	beego.Run()

}
