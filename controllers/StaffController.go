package controllers

import (
	"fmt"
	"todoapp/models"

	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/orm"
)

type StaffController struct {
	beego.Controller
}

func (sc *StaffController) Get() {
	/* Creating an instance of the Ormer */
	o := orm.NewOrm()
	var todos []*models.Todo

	/* Query the database to find all todos */
	_, err := o.QueryTable("todo").All(&todos)
	if err != nil {
		fmt.Println("Error querrying the database for todos: ", err)
	}

	/* Adding data to use in the template */
	sc.Data["Todos"] = todos
	/* Template name */
	sc.TplName = "staff.tpl"
}
