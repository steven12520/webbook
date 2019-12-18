package common

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
)

func init()  {

	logfile:=beego.AppConfig.String("logfile")
	logfile="{\"filename\":\""+logfile+"\"}"
	logs.SetLogger(logs.AdapterFile, `{"filename":"test.log"}`)

}

