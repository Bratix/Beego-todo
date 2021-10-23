package controllers

import (
	"fmt"
	"net/http"
	"todoapp/global"
	"todoapp/models"

	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/orm"
	"golang.org/x/crypto/bcrypt"
)

type RegisterController struct {
	beego.Controller
}

/* Gets the register form and displays it */
func (rc *RegisterController) Get() {
	rc.Data["Form"] = &models.User{}
	rc.TplName = "register.tpl"
}

/* Handles form submition for registration */
func (rc *RegisterController) Post() {
	/* Parses the form to user variable and hashes the password with bcrypt library */
	var user models.User
	rc.ParseForm(&user)
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	/* check if hasing produced an error */
	if err != nil {
		println("Error hashing password!")
		rc.Redirect("/register", http.StatusFound)
	}

	/* the hashed password gets stored */
	user.Password = string(hash)

	/* instance Ormer and insert the user into the database */
	o := orm.NewOrm()
	_, err = o.Insert(&user)
	fmt.Println(user)
	/* If an error occured redirect to /register */
	if err != nil {
		println(err)
		rc.Redirect("/register", http.StatusFound)
	}

	/* Create tokens and if no error occured authenticate the uzer */
	token, err := global.CreateToken(user.Id)

	if err != nil {
		fmt.Println("Error creating token: ", err)
		rc.Redirect("/", http.StatusFound)
	} else {
		global.Authenticate(user.Id, token, rc.Ctx.ResponseWriter)
	}

	rc.Redirect("/", http.StatusFound)
}
