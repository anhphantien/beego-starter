// @APIVersion 1.0.0
// @Title Beego starter
// @Description API docs
// @Schemes http,https
package routers

import (
	"beego-starter/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	ns :=
		web.NewNamespace("/v1",
			web.NSNamespace("/user",
				web.NSInclude(
					&controllers.UserController{},
				),
			),
		)
	web.AddNamespace(ns)
}
