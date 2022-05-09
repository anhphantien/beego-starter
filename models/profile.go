package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Profile struct {
	Id      uint64
	Age     int16     `orm:"null;default(0)"`
	Money   float64   `orm:"digits(65);decimals(2)"`
	User    *User     `orm:"reverse(one)"`
	Created time.Time `orm:"null;auto_now_add;type(datetime);precision(0)"`
	Updated time.Time `orm:"null;auto_now;type(datetime);precision(0)"`
}

func init() {
	orm.RegisterModel(new(Profile))
}
