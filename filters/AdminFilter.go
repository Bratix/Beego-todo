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

var FilterAdmin = func(ctx *context.Context) {
	/* if the user is on the /admin and the logged in user isn't a admin member redirect to index*/
	if strings.HasPrefix(ctx.Input.URL(), "/admin") {
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

		if !user.IsAdmin {
			fmt.Println("User isn't an admin!")
			ctx.Redirect(http.StatusFound, "/")
		}

	} else {
		return
	}

}
