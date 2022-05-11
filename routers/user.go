package routers

import (
	"beego-starter/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("v1/user", &controllers.ProfileController{}, "get:GetList")
	web.Router("v1/user/:id", &controllers.ProfileController{}, "get:GetById")
}
