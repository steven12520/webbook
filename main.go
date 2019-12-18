package main

import (
	_ "./routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {

	logs.Debug("已启动。。。。。。。。。")
	beego.Run()

}
