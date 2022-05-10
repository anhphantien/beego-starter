package controllers

import (
	"beego-starter/models"

	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
}

func (c *UserController) URLMapping() {
	c.Mapping("getList", c.getList)
}

// @Summary List of Users
// @Success 200 {object} models.User
// @Accept json
// @router / [get]
func (c *UserController) getList() {
	user := models.User{
		Id:   1,
		Name: "Phan Anh",
	}
	c.Data["json"] = user
	c.ServeJSON()
}
