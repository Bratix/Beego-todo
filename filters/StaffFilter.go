package filters

import (
	"fmt"
	"net/http"
	"strings"
	"todoapp/global"
	"todoapp/models"

	"github.com/astaxie/beego/context"
	"github.com/beego/beego/v2/client/orm"
)

var FilterStaff = func(ctx *context.Context) {
	/* if the user is on the /login or /register route nothing happens */
	if strings.HasPrefix(ctx.Input.URL(), "/staff") {
		extractedTokenData, err := global.ExtractTokenMetadata("AccessToken", ctx.Request)
		if err != nil {
			fmt.Println("Error extracting token metadata: ", err)
			ctx.Redirect(http.StatusFound, "/")
		}

		o := orm.NewOrm()
		user := models.User{}
		err = o.QueryTable("user").Filter("id", extractedTokenData.UserId).One(&user)

		if err != nil {
			fmt.Println("No user with selected id in database: ", err)
			ctx.Redirect(http.StatusFound, "/")
		}

		if !user.IsStaff {
			fmt.Println("User isn't a staff member!")
			ctx.Redirect(http.StatusFound, "/")
		}

	} else {
		return
	}

}
