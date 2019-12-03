package controllers

import "../models"
type RpcTestController struct {
	BaseController
}

func (self *RpcTestController) CreateOrders() {
	self.Data["pageTitle"] = "创建订单"
	var p models.PublicUsersModel
	p.ID=2
	bol:= p.GetPublicUsers()
	println(bol)
	self.display()
}


