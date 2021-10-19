package controllers

import (
	"fmt"
	"strconv"
	"todoapp/models"

	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/orm"
)

type TodoController struct {
	beego.Controller
}

func (t *TodoController) Get() {

	id, _ := strconv.Atoi(t.Ctx.Input.Param(":id"))
	todo := models.Todo{Id: id}
	o := orm.NewOrm()
	if err := o.Read(&todo); err != nil {
		fmt.Println(err)
	}

	t.Data["Todo"] = todo
	t.TplName = "detailView.tpl"
}

func (t *TodoController) AddTodo() {

	if t.Ctx.Input.IsPost() {
		var todo models.Todo
		t.ParseForm(&todo)
		o := orm.NewOrm()
		o.Insert(&todo)
		t.Redirect("/", 302)
	}

	t.Data["Form"] = &models.Todo{}
	t.TplName = "addTodo.tpl"
}

func (t *TodoController) EditTodo() {
	o := orm.NewOrm()
	if t.Ctx.Input.IsPost() {

		id, _ := strconv.Atoi(t.Ctx.Input.Param(":id"))
		var todo = models.Todo{Id: id}

		if o.Read(&todo) == nil {
			t.ParseForm(&todo)
			if _, err := o.Update(&todo); err == nil {
				fmt.Println("successful edit")
				t.Redirect("/", 302)
			}
		}
	}

	id, _ := strconv.Atoi(t.Ctx.Input.Param(":id"))
	var todo = models.Todo{Id: id}

	if o.Read(&todo) == nil {
		t.Data["Form"] = &todo
	} else {
		t.Redirect("/", 302)
	}

	t.TplName = "editTodo.tpl"
}

func (t *TodoController) DeleteTodo() {
	id, _ := strconv.Atoi(t.Ctx.Input.Param(":id"))

	o := orm.NewOrm()
	if num, err := o.Delete(&models.Todo{Id: id}); err == nil {
		fmt.Println(num)
	}
	t.Redirect("/", 302)
}
