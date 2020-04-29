package main

import (
	"./http"
	_ "./routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	//ziptest.TestCompress()
	//ziptest.TestCompress()
	//ziptest.TestDeCompress()
	//httpdate.GetBackDetail()
	//httpdate.GetTaskDetail()
	//httpdate.GetPhoneCheckNum()
	//httpdate.FastOnLineList()
	//reques()
	logs.Debug("已启动。。。。。。。。。1")
	beego.Run()

}
func reques() {
	//httpdate.GetPhoneCheckNum()
	for i := 0; i < 100; i++ {
		go func(s int) {
			for i := 0; i < 1000; i++ {
				httpdate.Getmfe()
				fmt.Println(s + i)
			}
		}(i)
	}

}
func Sort() {

}
