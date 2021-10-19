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

func (lc *LoginController) Get() {
	lc.Data["Form"] = &models.User{}
	lc.TplName = "login.tpl"
}

func (lc *LoginController) Post() {

	o := orm.NewOrm()
	var userform models.User
	lc.ParseForm(&userform)

	hash, err := bcrypt.GenerateFromPassword([]byte(userform.Password), bcrypt.DefaultCost)
	if err != nil {
		println("Error hashing entered password!")
		lc.Redirect("/login", 302)
	}
	userform.Password = string(hash)

	user := models.User{}
	_, err = o.QueryTable("user").Filter("username", userform.Username).Distinct().All(&user)
	fmt.Println(user)
	if err != nil {
		println("Error querring database")
		lc.Redirect("/login", 302)
	}

	ses, err := beego.GlobalSessions.SessionStart(lc.Ctx.ResponseWriter, lc.Ctx.Request)
	if err != nil {
		fmt.Println(err)
		lc.Redirect("/login", 302)
	}
	ses.Set("uid", user.Id)
	ses.Set("username", user.Username)
	lc.Redirect("/", 302)

}
