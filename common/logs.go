package common

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init() {

	logfile := beego.AppConfig.String("logfile")
	logfile = "{\"filename\":\"" + logfile + "\"}"
	logs.SetLogger(logs.AdapterFile, `{"filename":"test.log"}`)

}
