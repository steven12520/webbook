package main

import (
	_ "./routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {

	//ziptest.TestCompress()
	//ziptest.TestDeCompress()
	logs.Debug("已启动。。。。。。。。。")
	beego.Run()

}
