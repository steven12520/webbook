package models

import (
	_ "code.google.com/P/odbc"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB
var Dbsql *sql.DB

func init() {

	hoststr := beego.AppConfig.String("db.host")
	user := beego.AppConfig.String("db.user")
	password := beego.AppConfig.String("db.password")
	port := beego.AppConfig.String("db.port")

	dbname := beego.AppConfig.String("db.name")
	dburl := user + ":" + password + "@tcp(" + hoststr + ":" + port + ")/" + dbname
	//mysql
	database, err := sqlx.Open("mysql", dburl)

	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	//sql server
	hoststrm := beego.AppConfig.String("dbm.host")
	userm := beego.AppConfig.String("dbm.user")
	passwordm := beego.AppConfig.String("dbm.password")
	portm := beego.AppConfig.String("dbm.port")
	dbnamem := beego.AppConfig.String("dbm.name")
	sql_conn_str := fmt.Sprintf("driver={sql server};server=%s;port=%s;uid=%s;pwd=%s;database=%s", hoststrm, portm, userm, passwordm, dbnamem)

	databasesql, err := sql.Open("odbc", sql_conn_str)

	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db = database
	Dbsql = databasesql
}
