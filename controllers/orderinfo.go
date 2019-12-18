package controllers

import (
	"../models"
)
type OrderInfoController struct {
	BaseController
}

func (self *OrderInfoController) OrderInfoList() {
	self.Data["pageTitle"] = "下单结果查询"
	self.display()
}

func (self *OrderInfoController) GetOrderInfolist() {

	page, pe := self.GetInt("page")
	if pe != nil {
		page = 1
	}

	limit, le := self.GetInt("limit")
	if le != nil {
		limit = 10
	}
	self.pageSize = limit
	bookList := models.OrderinfoModel{}.GetList(page, limit)
	count := models.OrderinfoModel{}.GetCount()
	if le != nil {
		limit = 10
	}
	list := make([]map[string]interface{}, 0)

	for _, b := range bookList {
		m := make(map[string]interface{})
		m["Id"] = b.Id
		m["Gocount"] = b.Gocount
		m["Vin"] = b.Vin
		m["Types"] = b.Types
		m["Ordercount"] = b.Ordercount
		m["CreateName"] = b.CreateName
		m["Createtime"] = b.Createtime
		m["Gotype"] = b.Gotype
		list = append(list, m)
	}

	self.ajaxList("成功", MSG_OK, count, list)
}

func (self *OrderInfoController) Ajaxdel() {

	id, _ := self.GetInt64("id")
	var b models.OrderinfoModel
	b.Id = id
	bol := b.Delete()
	if bol {

	} else {
		self.ajaxMsg("删除失败", MSG_ERR)
	}
	var bd models.OrderinfodetailModel
	bd.Oid = id
	bol = bd.Delete()

	if bol {
		self.ajaxMsg("删除成功", MSG_OK)
	} else {
		self.ajaxMsg("删除失败", MSG_ERR)
	}
}

func (self *OrderInfoController) Detail() {
	id, _ := self.GetInt("id")
	self.Data["pageTitle"] = "下单结果详情"
	self.Data["id"] = id
	self.display()

}
func (self *OrderInfoController) GetOrderInfodetaillist() {

	page, pe := self.GetInt("page")
	if pe != nil {
		page = 1
	}
	id, _ := self.GetInt("id")

	limit, le := self.GetInt("limit")
	if le != nil {
		limit = 10
	}
	self.pageSize = limit
	bookList := models.OrderinfodetailModel{}.GetList(page, limit,id)
	count := models.OrderinfodetailModel{}.GetCount()
	if le != nil {
		limit = 10
	}
	list := make([]map[string]interface{}, 0)

	for _, b := range bookList {
		m := make(map[string]interface{})
		m["Id"] = b.Id
		m["Oid"] = b.Oid
		m["Vin"] = b.Vin
		m["Des"] = b.Des
		if b.Status==1 {
			m["Status"] = "成功"
		}else {
			m["Status"] = "失败"
		}
		m["Timelength"] = b.Timelength
		list = append(list, m)
	}

	self.ajaxList("成功", MSG_OK, count, list)
}
