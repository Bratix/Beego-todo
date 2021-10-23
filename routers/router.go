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
	beego.Router("/admin", &controllers.AdminController{}, "get:GetUsers")
	beego.Router("/admin/addadmin/:id", &controllers.AdminController{}, "get:AddAdmin")
	beego.Router("/admin/removeadmin/:id", &controllers.AdminController{}, "get:RemoveAdmin")
	beego.Router("/admin/addstaff/:id", &controllers.AdminController{}, "get:AddStaff")
	beego.Router("/admin/removestaff/:id", &controllers.AdminController{}, "get:RemoveStaff")
	beego.Router("/staff", &controllers.StaffController{})

	/* Iserting the filter */
	beego.InsertFilter("/*", beego.BeforeRouter, filters.FilterLoggedIn)
	beego.InsertFilter("/admin", beego.BeforeRouter, filters.FilterAdmin)
	beego.InsertFilter("/staff", beego.BeforeRouter, filters.FilterStaff)
}
