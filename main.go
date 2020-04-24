package main

import (
	"./common"
	"./http"
	_ "./routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	//ziptest.TestCompress()
	//ziptest.TestCompress()
	//ziptest.TestDeCompress()
	///httpdate.GetBackDetail()
	//httpdate.GetTaskDetail()
	//httpdate.GetPhoneCheckNum()
	//httpdate.FastOnLineList()
	logs.Debug("已启动。。。。。。。。。1")
	beego.Run()

}
func reques() {
	//httpdate.GetPhoneCheckNum()
	httpdate.UserHandler()
}
func Sort(list []int, left, right int) {
	if right == 0 {
		return
	}
	for index, num := range list {
		if index < right && num > list[index+1] {
			common.SwapGo(list, index, index+1)
		}
	}
	Sort(list, left, right-1)
}
