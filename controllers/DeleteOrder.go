package controllers

type  DeleteOrderController struct {
	BaseController
}

func (self *DeleteOrderController) Delete_order() {
	self.Data["pageTitle"] = "清除订单"
	self.display()
}
func (self *DeleteOrderController)Deletes()  {


	types,_:=self.GetInt("types")
	vin_delete:=self.GetString("vin_delete")
	vin_user,_:=self.GetInt("vin_user")
	starttime:=self.GetString("starttime")
	endtime:=self.GetString("endtime")
	if types==1 {//vin删除

		Deleteorder(vin_delete,vin_user,starttime,endtime)
	}else if types==2 {//用户时间删除

	}else if types==3 {//清空派单状态

	}else {

	}

}

func Deleteorder(vin_delete string ,vin_user int,starttime string ,endtime string) int  {

	return 0
}