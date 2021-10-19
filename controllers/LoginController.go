package controllers

import (
	"fmt"
	"todoapp/models"

	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/orm"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	beego.Controller
}

/* This function gets the form for the login template and displays it */
func (lc *LoginController) Get() {
	lc.Data["Form"] = &models.User{}
	lc.TplName = "login.tpl"
}

/* Post method used to log in users */
func (lc *LoginController) Post() {

	/* Instance new Ormer and parse the entered credentials into userform variable */
	o := orm.NewOrm()
	var userform models.User
	lc.ParseForm(&userform)

	/* Query the database for a user that has the entered username as username */
	user := models.User{}
	num, err := o.QueryTable("user").Filter("username", userform.Username).Distinct().All(&user)

	/* error if something goes wrong when running the query or if no results are found*/
	if err != nil {
		println("Error querring database")
		lc.Redirect("/login", 302)
	} else if num == 0 {
		fmt.Println("No such username in database!")
		lc.Redirect("/login", 302)
	}

	/* Starting the session and hangling the error */
	ses, err := beego.GlobalSessions.SessionStart(lc.Ctx.ResponseWriter, lc.Ctx.Request)
	if err != nil {
		fmt.Println(err)
		lc.Redirect("/login", 302)
	}

	/* User gets logged in if the password matches the one in the database */
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userform.Password))
	if err == nil {
		ses.Set("uid", user.Id)
		ses.Set("username", user.Username)
	} else {
		fmt.Println("Wrong password!")
	}

	/* redirect to index page */
	lc.Redirect("/", 302)

}

/* Logout function it sets uid to 0 and username to an epmty string */
func (lc *LoginController) Logout() {
	ses, _ := beego.GlobalSessions.SessionStart(lc.Ctx.ResponseWriter, lc.Ctx.Request)

	ses.Set("uid", 0)
	ses.Set("username", "")
	lc.Redirect("/", 302)
}
