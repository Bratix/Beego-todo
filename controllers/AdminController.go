package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"todoapp/models"

	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/orm"
)

type AdminController struct {
	beego.Controller
}

func (ac *AdminController) GetUsers() {
	/* Display all users */
	o := orm.NewOrm()
	var users []*models.User
	_, err := o.QueryTable("user").All(&users)

	if err == nil {
		ac.Data["Users"] = users
	}

	ac.TplName = "admin.tpl"
}

func (ac *AdminController) AddAdmin() {
	/* Selected users IsAdmin gets set to true if user isn't an admin already */
	id, _ := strconv.Atoi(ac.Ctx.Input.Param(":id"))

	o := orm.NewOrm()
	var user = models.User{Id: id}

	err := o.Read(&user)
	if err != nil {
		fmt.Println("Error querrying database: ", err)
	}

	if !user.IsAdmin {

		user.IsAdmin = true
		_, err = o.Update(&user)
		if err != nil {
			fmt.Println("Error saving data to database: ", err)
		}
	}

	ac.Redirect("/admin", http.StatusFound)
}

func (ac *AdminController) RemoveAdmin() {
	/* Selected users IsAdmin gets set to false if user is staff member */
	id, _ := strconv.Atoi(ac.Ctx.Input.Param(":id"))

	o := orm.NewOrm()
	var user = models.User{Id: id}

	err := o.Read(&user)
	if err != nil {
		fmt.Println("Error querrying database: ", err)
	}

	if user.IsAdmin {

		user.IsAdmin = false
		_, err = o.Update(&user)
		if err != nil {
			fmt.Println("Error saving data to database: ", err)
		}
	}

	ac.Redirect("/admin", http.StatusFound)
}

func (ac *AdminController) AddStaff() {
	/* Selected users IsStaff gets set to true if user isn't an staff already */
	id, _ := strconv.Atoi(ac.Ctx.Input.Param(":id"))

	o := orm.NewOrm()
	var user = models.User{Id: id}

	err := o.Read(&user)
	if err != nil {
		fmt.Println("Error querrying database: ", err)
	}

	if !user.IsStaff {

		user.IsStaff = true
		_, err = o.Update(&user)
		if err != nil {
			fmt.Println("Error saving data to database: ", err)
		}
	}

	ac.Redirect("/admin", http.StatusFound)
}

func (ac *AdminController) RemoveStaff() {
	/* Selected users IsStaff gets set to false if user is a staff member */
	id, _ := strconv.Atoi(ac.Ctx.Input.Param(":id"))

	o := orm.NewOrm()
	var user = models.User{Id: id}

	err := o.Read(&user)
	if err != nil {
		fmt.Println("Error querrying database: ", err)
	}

	if user.IsStaff {

		user.IsStaff = false
		_, err = o.Update(&user)
		if err != nil {
			fmt.Println("Error saving data to database: ", err)
		}
	}

	ac.Redirect("/admin", http.StatusFound)
}
