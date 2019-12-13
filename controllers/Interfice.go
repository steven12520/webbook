package controllers

import (
	"../models"
)

type InterficeController struct {
	BaseController
}

func (self *InterficeController) InterficeList() {

	self.Data["pageTitle"] = "接口列表管理"
	self.display()
}
func (self *InterficeController) Add() {
	self.Data["pageTitle"] = "新建接口"
	self.display()
}

func (self *InterficeController) GetInterficeList() {

	list := models.Interfice{}.GetInterficeList()

	maplist := make([]map[string]interface{}, 0)
	for _, v := range list {
		m := make(map[string]interface{})
		m["Id"] = v.Id
		m["Iname"] = v.Iname
		m["Txt"] = v.Txt
		m["Createtime"] = v.Createtime

		maplist = append(maplist, m)
	}
	self.ajaxList("成功", MSG_OK, 1, maplist)
}

func (self *InterficeController) Edit() {

	id, _ := self.GetInt("id")

	var b models.Interfice
	b.Id = id
	book, _ := b.GetInterficeById()

	row := make(map[string]interface{})
	row["Id"] = book.Id
	row["Iname"] = book.Iname
	row["Txt"] = book.Txt
	self.Data["Book"] = row

	self.display()
}

func (self InterficeController) AjaxSave() {

	var m models.Interfice
	m.Id, _ = self.GetInt("Id")
	m.Iname = self.GetString("Iname")
	m.Txt = self.GetString("Txt")

	bol := false
	if m.Id > 0 { //更新
		bol = m.Update()
	} else {
		bol = m.Add()
	}
	if bol {
		self.ajaxMsg("保存成功", MSG_OK)
	} else {
		self.ajaxMsg("保存失败", MSG_ERR)
	}
}
func (self InterficeController) Ajaxdel() {
	var m models.Interfice
	m.Id, _ = self.GetInt("id")

	bol := m.Delete()
	if bol {
		self.ajaxMsg("", MSG_OK)
	} else {
		self.ajaxMsg("删除失败", MSG_ERR)
	}
}
