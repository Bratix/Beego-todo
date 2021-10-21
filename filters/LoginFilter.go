package filters

import (
	"fmt"
	"net/http"
	"strings"
	"todoapp/global"

	"github.com/astaxie/beego/context"
)

/* Filter to check if user is logged in */
var FilterLoggedIn = func(ctx *context.Context) {
	/* if the user is on the /login or /register route nothing happens */
	if strings.HasPrefix(ctx.Input.URL(), "/login") || strings.HasPrefix(ctx.Input.URL(), "/register") {
		return
	}

	/* if the user id can't be accessed from the session or it is 0 the user gets redirected to the login page */

	_, err := global.ExtractTokenMetadata("AccessToken", ctx.Request)
	if err != nil {
		fmt.Println("Auth token not found: ", err)

		errRefresh := global.RefreshToken(ctx.ResponseWriter, ctx.Request)
		if errRefresh == nil {
			fmt.Println("Access token refreshed!")
			ctx.Redirect(http.StatusFound, ctx.Request.RequestURI)
		} else {
			fmt.Println("Error refresing token!", errRefresh)
		}

		ctx.Redirect(http.StatusTemporaryRedirect, "/login")
	}

}
