package test

import (
	"fmt"
	"github.com/astaxie/beego"
	"testing"
)

func TestFrom(t *testing.T) {

	defer func() {
		i := recover()
		if i != nil {
			fmt.Println("recover error")
			fmt.Println(i)
		}
	}()

	url := beego.AppConfig.String("app.url") + "/app/TaskSaveSimple.ashx"
	fmt.Println(url)

}
