package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id      uint64
	Name    string    `orm:"null"`
	Profile *Profile  `orm:"null;rel(one);on_delete(set_null)"`
	Posts   []*Post   `orm:"reverse(many)"`
	Created time.Time `orm:"null;auto_now_add;type(datetime);precision(0)"`
	Updated time.Time `orm:"null;auto_now;type(datetime);precision(0)"`
}

func init() {
	orm.RegisterModel(new(User))
}
