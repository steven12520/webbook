package main

import (
	"./http"
	_ "./routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	//sor()
	reques()
	//return
	logs.Debug("已启动。。。。。。。。。1")
	beego.Run()
}

func reques() {
	//ziptest.TestCompress()
	////ziptest.TestCompress()
	////ziptest.TestDeCompress()
	////httpdate.GetBackDetail()
	////httpdate.GetTaskDetail()
	////httpdate.GetPhoneCheckNum()
	//httpdate.IniteData(1726)
	httpdate.GetServiceProgram(1726)
}
func sor() {

	list := [...]int{1, 2, 6, 9, 0, 5, 8, 4}
	for i := 0; i < len(list); i++ {
		for j := 1; j < len(list)-i; j++ {
			if list[j-1] > list[j] {
				list[j-1], list[j] = list[j], list[j-1]
			}
		}
	}
	fmt.Println(list)

}
