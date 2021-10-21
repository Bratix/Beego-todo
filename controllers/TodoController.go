package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"todoapp/global"
	"todoapp/models"

	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/orm"
)

type TodoController struct {
	beego.Controller
}

/* Displays details about a single todo */
func (t *TodoController) GetTodo() {

	/* Get id from route */
	id, _ := strconv.Atoi(t.Ctx.Input.Param(":id"))
	/* Todo with specified id is created */
	todo := models.Todo{Id: id}
	/* Instance Ormer */
	o := orm.NewOrm()

	/* Return todo with the specified id and the user it belongs to */
	err := o.QueryTable("todo").RelatedSel().Filter("Id", id).One(&todo)
	if err != nil {
		fmt.Println(err)
	}

	/* Specify data to use in template and template name */
	t.Data["Todo"] = todo
	t.TplName = "detailView.tpl"
}

/* Function to add todos to database */
func (t *TodoController) AddTodo() {

	/* Handles post request */
	if t.Ctx.Input.IsPost() {
		/* Parse form and instance orm */
		var todo models.Todo
		t.ParseForm(&todo)
		o := orm.NewOrm()

		/* Get user id from session */
		uid, err := global.ExtractTokenMetadata("AccessToken", t.Ctx.Request)
		if err != nil {
			fmt.Println("User not logged in: ", err)
		}
		var user = models.User{Id: uid.UserId}

		/* Logged in user becomes the owner of created todo */
		todo.User = &user
		/* Insert todo in database and redirect to index page */
		o.Insert(&todo)
		fmt.Println("Succesfully added to database!")
		t.Redirect("/", http.StatusFound)
	}

	/* Handles get request */
	t.Data["Form"] = &models.Todo{}
	t.TplName = "addTodo.tpl"
}

func (t *TodoController) EditTodo() {
	/* Instance orm */
	o := orm.NewOrm()
	/* Handles post request */
	if t.Ctx.Input.IsPost() {
		/* Get id from url */
		id, _ := strconv.Atoi(t.Ctx.Input.Param(":id"))
		var todo = models.Todo{Id: id}
		/* Get the id of the logged in user */
		uid, err := global.ExtractTokenMetadata("AccessToken", t.Ctx.Request)
		if err != nil {
			fmt.Println("User not logged in: ", err)
		}
		/* Querry table to find the selected todo */
		err = o.QueryTable("todo").RelatedSel().Filter("Id", id).One(&todo)

		/* If no error occured and the todo user id matches logged in user id todo gets updated */
		if err == nil && todo.User.Id == uid.UserId {

			t.ParseForm(&todo)
			if _, err := o.Update(&todo); err == nil {
				fmt.Println("Successful edit!")
				t.Redirect("/", http.StatusFound)
			}
		} else {
			/* If todo didn't get updated */
			fmt.Println("Unsuccessful edit!")
			t.Redirect("/", http.StatusFound)
		}
	}

	/* get id of todo */
	id, _ := strconv.Atoi(t.Ctx.Input.Param(":id"))
	var todo = models.Todo{Id: id}
	/* Get the todo with specified id (error should have been handled) */
	_ = o.QueryTable("todo").RelatedSel().Filter("Id", id).One(&todo)

	/* Get id of logged in user */
	userAccessDetail, err := global.ExtractTokenMetadata("AccessToken", t.Ctx.Request)
	if err != nil {
		fmt.Println("User not logged in: ", err)
		t.Redirect("/", http.StatusFound)
	}
	/* Get user id from session */

	/* If the logged in user is the owner of todo he gets the edit form, else he gets redirected to index page */
	if todo.User.Id == userAccessDetail.UserId {
		t.Data["Form"] = &todo
		t.TplName = "editTodo.tpl"
	} else {
		fmt.Println("User isn't the owner of todo!", todo.User.Id, " != ", userAccessDetail.UserId)
		t.Redirect("/", http.StatusFound)
	}
}

/* Deletes todos from database */
func (t *TodoController) DeleteTodo() {
	/* Get id from url */
	id, _ := strconv.Atoi(t.Ctx.Input.Param(":id"))
	/* New orm instance */
	o := orm.NewOrm()
	var todo models.Todo

	/* get logged in user id */
	uid, err := global.ExtractTokenMetadata("AccessToken", t.Ctx.Request)
	if err != nil {
		fmt.Println("User not logged in: ", err)
		t.Redirect("/", http.StatusFound)
	}
	/* get requested todo from database*/
	err = o.QueryTable("todo").RelatedSel().Filter("Id", id).One(&todo)

	/* If no error occured and todo owner is the same as logged in user delete the record, else print out that user isn't the owner */
	if err == nil && todo.User.Id == uid.UserId {
		if num, err := o.Delete(&models.Todo{Id: id}); err == nil {
			fmt.Println(num)
			fmt.Println("Record deleted! todo user id and logged in user id:", todo.User.Id, " ", uid.UserId)
		}
	} else {
		fmt.Println("User isn't the owner of todo!")
	}
	/* Redirect to index page */
	t.Redirect("/", http.StatusFound)
}
