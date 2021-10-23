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
		fmt.Println("Error querrying database: ", err)
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
		extractedTokenData, _ := global.ExtractTokenMetadata("AccessToken", t.Ctx.Request)
		var user = models.User{Id: extractedTokenData.UserId}

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
		extractedTokenData, _ := global.ExtractTokenMetadata("AccessToken", t.Ctx.Request)

		/* Querry table to find the selected todo */
		err := o.QueryTable("todo").RelatedSel().Filter("Id", id).One(&todo)

		/* Querry table to find logged in user */
		user := models.User{Id: extractedTokenData.UserId}
		errUser := o.Read(&user)

		/* Check if there are no errors and if logged in user is the owner of todo, an admin member or staff */
		if err == nil && todo.User.Id == extractedTokenData.UserId || errUser == nil && user.IsAdmin || errUser == nil && user.IsStaff {

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
	extractedTokenData, _ := global.ExtractTokenMetadata("AccessToken", t.Ctx.Request)

	user := models.User{Id: extractedTokenData.UserId}
	errUser := o.Read(&user)

	/* Check if there are no errors and if logged in user is the owner of todo, an admin member or staff */
	if todo.User.Id == extractedTokenData.UserId || errUser == nil && user.IsAdmin || errUser == nil && user.IsStaff {
		t.Data["Form"] = &todo
		t.TplName = "editTodo.tpl"
	} else {
		fmt.Println("User can't access todo!")
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
	extractedTokenData, _ := global.ExtractTokenMetadata("AccessToken", t.Ctx.Request)
	/* get requested todo from database*/
	err := o.QueryTable("todo").RelatedSel().Filter("Id", id).One(&todo)

	user := models.User{Id: extractedTokenData.UserId}
	errUser := o.Read(&user)

	/* Check if there are no errors and if logged in user is the owner of todo, an admin member or staff */
	if err == nil && todo.User.Id == extractedTokenData.UserId || errUser == nil && user.IsAdmin {
		if num, err := o.Delete(&models.Todo{Id: id}); err == nil {
			fmt.Println(num)
			fmt.Println("Record deleted!")
		}
	} else {
		fmt.Println("User can't access todo!")
	}
	/* Redirect to index page */
	t.Redirect("/", http.StatusFound)
}
