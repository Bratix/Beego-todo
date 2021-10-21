package controllers

import (
	"fmt"
	"todoapp/global"
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

	uid, err := global.ExtractTokenMetadata("AccessToken", c.Ctx.Request)
	if err != nil {
		fmt.Println("User not logged in: ", err)
	}

	/* Query the database to find all the todos that have the logged in user foreign key */
	_, err = o.QueryTable("todo").Filter("User__id", uid.UserId).All(&todos)
	if err != nil {
		fmt.Println("Error querrying the database for todos: ", err)
	}

	/* Adding data to use in the template */
	c.Data["Todos"] = todos
	/* Template name */
	c.TplName = "index.tpl"
}
