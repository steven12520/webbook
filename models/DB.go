package models

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {

	hoststr := beego.AppConfig.String("db.host")
	user := beego.AppConfig.String("db.user")
	password := beego.AppConfig.String("db.password")
	port := beego.AppConfig.String("db.port")
	dbname := beego.AppConfig.String("db.name")
	dburl := user + ":" + password + "@tcp(" + hoststr + ":" + port + ")/" + dbname
	database, err := sqlx.Open("mysql", dburl)

	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db = database
}
