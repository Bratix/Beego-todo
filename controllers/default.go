package controllers

import (
	"todoapp/models"

	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/orm"
)

type MainController struct {
	beego.Controller
}

/* This function displays all the todos a user has */
func (c *MainController) Get() {
	/* Creating an instance of the Ormer */
	o := orm.NewOrm()
	var todos []*models.Todo
	/* Acquiring the session that is created by login or register controller */
	ses, _ := beego.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	/* Getting the username and id from session  */
	username := ses.Get("username")
	uid := ses.Get("uid").(int)

	/* Query the database to find all the todos that have the logged in user foreign key */
	num, err := o.QueryTable("todo").Filter("User__id", uid).All(&todos)
	if err != nil {
		println(num)
	}

	/* Adding data to use in the template */
	c.Data["Todos"] = todos
	c.Data["Username"] = username
	/* Template name */
	c.TplName = "index.tpl"
}
