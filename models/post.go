package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Post struct {
	Id      uint64
	Name    string    `orm:"null"`
	User    *User     `orm:"rel(fk)"`
	Created time.Time `orm:"null;auto_now_add;type(datetime);precision(0)"`
	Updated time.Time `orm:"null;auto_now;type(datetime);precision(0)"`
}

func init() {
	orm.RegisterModel(new(Post))
}
