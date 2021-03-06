package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/Index", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")
	beego.Router("/login", &controllers.LoginController{}, "*:LoginIn")
	beego.Router("/login_out", &controllers.LoginController{}, "*:LoginOut")

	beego.AutoRouter(&controllers.UserInfoController{})
	beego.AutoRouter(&controllers.BookInfoController{})
	beego.AutoRouter(&controllers.BorrowBookController{})
	beego.AutoRouter(&controllers.BookClassificationController{})
	beego.AutoRouter(&controllers.RpcTestController{})
	beego.AutoRouter(&controllers.DeleteOrderController{})
	beego.AutoRouter(&controllers.InterficeController{})
	beego.AutoRouter(&controllers.OrderInfoController{})
	beego.AutoRouter(&controllers.YstestController{})
	beego.AutoRouter(&controllers.KGController{})
	beego.AutoRouter(&controllers.H5Controller{})

}
