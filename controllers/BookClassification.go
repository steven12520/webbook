package controllers

import (
	"../models"
	"fmt"
	"time"
)

type BookClassificationController struct {
	BaseController
}

func (self *BookClassificationController) BookClassificationList() {

	self.Data["pageTitle"] = "类型管理"
	self.display()
}
func (self *BookClassificationController) Add() {

	self.Data["pageTitle"] = "新建类型"

	self.display()
}

func (self *BookClassificationController) GetBookClassificationList() {

	page, pe := self.GetInt("page")

	if pe != nil {
		page = 1
	}

	limit, le := self.GetInt("limit")
	if le != nil {
		limit = 10
	}
	self.pageSize = limit
	list := models.BookClassification{}.GetBookClassificationList(page, limit)
	count := models.BookClassification{}.GetBookClassificationCount()

	maplist := make([]map[string]interface{}, 0)
	for _, v := range list {
		m := make(map[string]interface{})
		m["Id"] = v.Id
		m["Type"] = v.Type
		m["CreateTime"] = v.CreateTime
		m["Name"] = v.Name
		maplist = append(maplist, m)
	}
	self.ajaxList("成功", MSG_OK, count, maplist)
}

func (self *BookClassificationController) Edit() {

	id, _ := self.GetInt("id")

	var b models.BookClassification
	b.Id = id
	fmt.Println(b)
	book, _ := b.GetBookClassificationById()

	row := make(map[string]interface{})
	row["Id"] = book.Id
	row["Name"] = book.Name
	row["CreateTime"] = book.CreateTime
	row["Type"] = book.Type
	self.Data["Book"] = row

	self.display()
}

func (self *BookClassificationController) AjaxSave() {

	id, _ := self.GetInt("id")
	var b models.BookClassification
	b.Id = id
	b.Name = self.GetString("Name")
	b.Type, _ = self.GetInt("Type")
	b.CreateTime = time.Now().Format("2006-01-02 15:04:05")

	bol := false
	if b.Id > 0 {
		bol = b.Update()
	} else {
		bol = b.Save()
	}
	if bol {
		self.ajaxMsg("", MSG_OK)
	} else {
		self.ajaxMsg("保存失败", MSG_ERR)
	}
}
func (self *BookClassificationController) Ajaxdel() {

	id, _ := self.GetInt("id")
	var b models.BookClassification
	b.Id = id
	bol := b.Delete()
	if bol {
		self.ajaxMsg("删除成功", MSG_OK)
	} else {
		self.ajaxMsg("删除失败", MSG_ERR)
	}
}
