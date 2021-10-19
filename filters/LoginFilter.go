package filters

import (
	"strings"

	"github.com/astaxie/beego/context"
)

/* Filter to check if user is logged in */
var FilterLoggedIn = func(ctx *context.Context) {
	/* if the user is on the /login or /register route nothing happens */
	if strings.HasPrefix(ctx.Input.URL(), "/login") || strings.HasPrefix(ctx.Input.URL(), "/register") {
		return
	}

	/* if the user id can't be accessed from the session or it is 0 the user gets redirected to the login page */
	id, ok := ctx.Input.Session("uid").(int)
	if !ok || id == 0 {
		ctx.Redirect(302, "/login")
	}
}
