package controllers

import (
	"todoapp/models"

	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/orm"
	"golang.org/x/crypto/bcrypt"
)

type RegisterController struct {
	beego.Controller
}

func (rc *RegisterController) Get() {
	rc.Data["Form"] = &models.User{}
	rc.TplName = "register.tpl"
}

func (rc *RegisterController) Post() {
	var user models.User
	rc.ParseForm(&user)
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		println("Error hashing password!")
		rc.Redirect("/register", 302)
	}

	user.Password = string(hash)

	o := orm.NewOrm()
	_, err = o.Insert(&user)

	if err != nil {
		println(err)
		rc.Redirect("/register", 302)
	}

	ses, err := beego.GlobalSessions.SessionStart(rc.Ctx.ResponseWriter, rc.Ctx.Request)

	if err != nil {
		println("Error logging in user!")
		rc.Redirect("/login", 302)
	}

	ses.Set("uid", user.Id)
	ses.Set("username", user.Username)
	rc.Redirect("/", 302)

}
