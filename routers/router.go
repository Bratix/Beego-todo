package routers

import (
	"todoapp/controllers"
	"todoapp/filters"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/todo/:id", &controllers.TodoController{})
	beego.Router("/add", &controllers.TodoController{}, "get,post:AddTodo")
	beego.Router("/edit/:id", &controllers.TodoController{}, "get,post:EditTodo")
	beego.Router("/delete/:id", &controllers.TodoController{}, "get:DeleteTodo")
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})

	beego.InsertFilter("/*", beego.BeforeRouter, filters.FilterLoggedIn)
}
