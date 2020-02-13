package controllers

import (
	"../models"
	"strconv"
	"strings"
)

type DeleteOrderController struct {
	BaseController
}

func (self *DeleteOrderController) Delete_order() {
	self.Data["pageTitle"] = "清除订单"
	self.display()
}
func (self *DeleteOrderController) Deletes() {

	types, _ := self.GetInt("types")
	vin_delete := self.GetString("vin_delete")
	vin_user, _ := self.GetInt("vin_user")
	starttime := self.GetString("starttime")
	endtime := self.GetString("endtime")
	result := 0
	if types == 1 { //vin删除
		result = Deleteorder(vin_delete)
	} else if types == 2 { //用户时间删除
		result = DeleteorderUser(vin_user, starttime, endtime)
	} else if types == 3 { //清空派单状态
		result = DeleteAsstask()
	} else {
		result = -1
	}

	if result == 1 {
		self.ajaxMsg("成功", MSG_OK)
	} else if result == 2 {
		self.ajaxMsg("无可删除数据", MSG_ERR)
	} else if result == 0 {
		self.ajaxMsg("删除数据失败", MSG_ERR)
	} else {
		self.ajaxMsg("参数错误", MSG_ERR)
	}

}

//根据vin删除订单
func Deleteorder(vin_delete string) int {

	var model models.TaskCarBasicModel
	model.Vin = vin_delete
	list := model.GetList()
	idlist := make([]string, 0)
	if list != nil && len(list) > 0 {
		for _, v := range list {
			idlist = append(idlist, strconv.Itoa(v.Id))
		}
	}
	if len(idlist) > 0 {
		model.IdList = strings.Join(idlist, ",")
		model.DeleteAssignedTask()
		return model.DeleteOrder()
	} else {
		return 2
	}
}

//根据用户删除订单
func DeleteorderUser(vin_user int, starttime string, endtime string) int {
	var model models.TaskCarBasicModel
	model.CreateUserId = vin_user
	model.StartTime = starttime
	model.EndTime = endtime
	list := model.GetList()
	idlist := make([]string, 0)
	if list != nil && len(list) > 0 {
		for _, v := range list {
			idlist = append(idlist, strconv.Itoa(v.Id))
		}
	}
	if len(idlist) > 0 {
		model.IdList = strings.Join(idlist, ",")
		model.DeleteAssignedTask()
		return model.DeleteOrder()
	} else {
		return 2
	}
}

//删除所有派单信息
func DeleteAsstask() int {

	var model models.TaskCarBasicModel
	return model.DeleteLoginStatusRecords()
}
