package controllers

import (
	"todoapp/models"

	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	o := orm.NewOrm()
	var todos []*models.Todo

	num, err := o.QueryTable("todo").All(&todos)
	if err != nil {
		println(num)
	}

	ses, _ := beego.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	username := ses.Get("username")

	c.Data["Todos"] = todos
	c.Data["Username"] = username
	c.TplName = "index.tpl"
}
