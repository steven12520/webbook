package main

import (
	"./http"
	_ "./routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	//ziptest.TestCompress()
	//ziptest.TestCompress()
	//ziptest.TestDeCompress()
	//httpdate.GetBackDetail()
	//httpdate.GetTaskDetail()
	logs.Debug("已启动。。。。。。。。。1")
	beego.Run()

}
func reques() {
	//httpdate.GetPhoneCheckNum()
	httpdate.UserHandler()
}
