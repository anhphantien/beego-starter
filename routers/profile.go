package routers

import (
	"beego-starter/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("v1/profile", &controllers.ProfileController{}, "get:GetList")
	web.Router("v1/profile/:id", &controllers.ProfileController{}, "get:GetById")
}
