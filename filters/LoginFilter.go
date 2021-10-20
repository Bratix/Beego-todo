package filters

import (
	"fmt"
	"net/http"
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
	id, err := ExtractTokenMetadata(ctx.Request)

	if err != nil || id == 0 {
		fmt.Println("Auth token not found: ", err)
		fmt.Println("Id : ", id)
		ctx.Redirect(http.StatusTemporaryRedirect, "/login")
	}
}
