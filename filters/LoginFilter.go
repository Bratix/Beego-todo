package filters

import (
	"strings"

	"github.com/astaxie/beego/context"
)

var FilterLoggedIn = func(ctx *context.Context) {
	if strings.HasPrefix(ctx.Input.URL(), "/login") || strings.HasPrefix(ctx.Input.URL(), "/register") {
		return
	}

	_, ok := ctx.Input.Session("uid").(int)
	if !ok {
		ctx.Redirect(302, "/login")
	}
}
