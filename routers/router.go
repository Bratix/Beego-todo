package routers

import (
	"todoapp/controllers"
	"todoapp/filters"

	"github.com/astaxie/beego"
)

func init() {
	/* Static route leads to / */
	beego.Router("/", &controllers.MainController{})
	/* Route with parameter that we can access */
	beego.Router("/todo/:id", &controllers.TodoController{}, "get:GetTodo")
	/* Requesting a specific method to handle a server call */
	beego.Router("/add", &controllers.TodoController{}, "get,post:AddTodo")
	beego.Router("/edit/:id", &controllers.TodoController{}, "get,post:EditTodo")
	beego.Router("/delete/:id", &controllers.TodoController{}, "get:DeleteTodo")
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")

	/* Iserting the filter */
	beego.InsertFilter("/*", beego.BeforeRouter, filters.FilterLoggedIn)
}
