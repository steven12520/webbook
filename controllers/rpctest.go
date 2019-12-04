package controllers

import (
	"../models"
	"strconv"
	"../common"
	"../http"
)
type RpcTestController struct {
	BaseController
}

func (self *RpcTestController) CreateOrders() {
	self.Data["pageTitle"] = "创建订单"
	self.display()
}
//获取用户信息
func (self *RpcTestController) GetPublicUsers()  {

	userid,_:=self.GetInt("userid")

	var p models.PublicUsersModel
	p.ID=userid
	list:= p.GetPublicUsers()

	self.ajaxList("成功", MSG_OK, 0, list)
}

//获取方案
func (self *RpcTestController) GetProgrammeConfiguration()  {

	userid,_:=self.GetInt("userid")

	var p models.ProgrammeConfigurationModel
	list:= p.Get(strconv.Itoa(userid))

	self.ajaxList("成功", MSG_OK, 0, list)
}
//获取特殊照片
func (self *RpcTestController) GetPicSpecialSet()  {

	userid,_:=self.GetInt("userid")

	var p models.PicSpecialSetModel
	list:= p.GetPicSpecialSet(strconv.Itoa(userid))

	self.ajaxList("成功", MSG_OK, 0, list)
}
//获取产品类型
func (self *RpcTestController) GetUserProductType()  {

	userid,_:=self.GetInt("userid")

	var p models.UserProductTypeModel
	list:= p.GetUserProductType(userid)

	self.ajaxList("成功", MSG_OK, 0, list)
}

func (self *RpcTestController) SaveOrder()  {
	userid,_:=self.GetInt("userid")
	configID,_:=self.GetInt("configID")
	procductlist,_:=self.GetInt("procductlist")
	vin:=self.GetString("vin")
	ordercount,_:=self.GetInt("ordercount")
	gocount,_:=self.GetInt("gocount")

	vin=common.GetRandvin(vin)

	httpdate.RequesSaveOrder(userid,configID,procductlist,vin,ordercount,gocount)

	//fmt.Println(userid,configID,procductlist,vin,ordercount,gocount)





	self.ajaxMsg("成功", MSG_OK)
}