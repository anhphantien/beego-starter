package controllers

import (
	"beego-starter/models"
	"time"

	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
}

func (c *UserController) URLMapping() {
	c.Mapping("GetList", c.GetList)
	c.Mapping("GetById", c.GetById)
}

// @Summary List of Users
// @Success 200 {array} []models.User
// @Accept json
// @router / [get]
func (c *UserController) GetList() {
	users := []models.User{
		{
			Id:      1,
			Name:    "Phan Anh",
			Created: time.Now(),
			Updated: time.Now(),
		},
		{
			Id:      2,
			Name:    "Phan Anh",
			Created: time.Now(),
			Updated: time.Now(),
		},
	}
	c.Data["json"] = users
	c.ServeJSON()
}

// @Summary Get User by ID
// @Success 200 {object} models.User
// @Accept json
// @router /:id [get]
func (c *UserController) GetById() {
	user := models.User{
		Id:      1,
		Name:    "Phan Anh",
		Created: time.Now(),
		Updated: time.Now(),
	}
	c.Data["json"] = user
	c.ServeJSON()
}
