package main

import (
	_ "beego-starter/models"
	_ "beego-starter/routers"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbDriver, _ := web.AppConfig.String("db_driver")
	dbUser, _ := web.AppConfig.String("db_user")
	dbPwd, _ := web.AppConfig.String("db_pwd")
	dbName, _ := web.AppConfig.String("db_name")
	dbCharset, _ := web.AppConfig.String("db_charset")
	dbCollation, _ := web.AppConfig.String("db_collation")

	dataSource := dbUser + ":" + dbPwd + "@/" + dbName + "?charset=" + dbCharset + "&collation=" + dbCollation

	orm.RegisterDriver(dbDriver, orm.DRMySQL)
	orm.RegisterDataBase("default", dbDriver, dataSource)

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	if web.BConfig.RunMode == "dev" {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	web.Run()
}
