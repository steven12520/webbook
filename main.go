package main

import (
	"./http"
	_ "./routers"
	"github.com/astaxie/beego/logs"
)

func main() {

	//ziptest.TestCompress()
	//ziptest.TestDeCompress()
	//httpdate.GetBackDetail()
	reques()
	logs.Debug("已启动。。。。。。。。。")
	//beego.Run()

}
func reques() {
	//httpdate.GetPhoneCheckNum()
	httpdate.UserHandler()
}
