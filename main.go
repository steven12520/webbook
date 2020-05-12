package main

import (
	_ "./routers"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	//sor()
	//reques()
	//return
	//reques()
	//fmt.Println(runtime.NumCPU())
	//fmt.Println(runtime.NumCPU())
	logs.Debug("已启动。。。。。。。。。1")
	beego.Run()
}

func reques() error {
	//ziptest.TestCompress()
	////ziptest.TestCompress()
	////ziptest.TestDeCompress()
	////httpdate.GetBackDetail()
	//httpdate.GetTaskDetail()
	////httpdate.GetPhoneCheckNum()
	////httpdate.FastOnLineList()

	//httpdate.PretrailSubmitPass(1926621, 10269)
	return errors.New("ddddddddd")
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
