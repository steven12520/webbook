package controllers

import (
	"../models"
	"fmt"
	"time"
)

type BorrowBookController struct {
	BaseController
}

func (self *BorrowBookController) BorrowBookList() {

	self.Data["pageTitle"] = "借阅管理"
	self.display()
}
func (self *BorrowBookController) Add() {
	self.Data["pageTitle"] = "新建借阅"
	self.display()
}

func (self *BorrowBookController) GetBorrowBookList() {

	page, pe := self.GetInt("page")

	if pe != nil {
		page = 1
	}

	limit, le := self.GetInt("limit")
	if le != nil {
		limit = 10
	}
	self.pageSize = limit
	fmt.Println()
	list := models.BorrowBook{}.GetBorrowBookList(page, limit)
	count := models.BorrowBook{}.GetBorrowBookCount()

	maplist := make([]map[string]interface{}, 0)
	for _, v := range list {
		m := make(map[string]interface{})
		m["Id"] = v.Id
		m["Loginname"] = v.Loginname
		m["Author"] = v.Author
		m["BookName"] = v.BookName
		m["Type"] = v.Type
		m["AlsoTime"] = v.AlsoTime
		m["ToAlsoTime"] = v.ToAlsoTime
		m["CreateTime"] = v.CreateTime
		m["UserId"] = v.UserId
		m["BookId"] = v.BookId
		maplist = append(maplist, m)
	}
	self.ajaxList("成功", MSG_OK, count, maplist)
}

func (self *BorrowBookController) Edit() {

	id, _ := self.GetInt("id")

	var b models.BorrowBook
	b.Id = id
	fmt.Println(b)
	book, _ := b.GetBorrowBookById()

	row := make(map[string]interface{})
	row["Id"] = book.Id
	row["Loginname"] = book.Loginname
	row["Author"] = book.Author
	row["BookName"] = book.BookName
	row["Type"] = book.Type
	row["ToAlsoTime"] = book.ToAlsoTime
	row["CreateTime"] = book.CreateTime
	row["UserId"] = book.UserId
	row["BookId"] = book.BookId
	self.Data["Book"] = row

	self.display()
}

func (self BorrowBookController) AjaxSave() {

	var m models.BorrowBook
	m.Id, _ = self.GetInt("Id")
	m.Loginname = self.GetString("Loginname")
	m.Author = self.GetString("Author")
	m.BookName = self.GetString("BookName")
	m.Type, _ = self.GetInt("Type")
	m.AlsoTime = time.Now().Format("2006-01-02 15:04:05")
	m.ToAlsoTime = self.GetString("ToAlsoTime")
	m.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	m.UserId, _ = self.GetInt("UserId")
	m.BookId, _ = self.GetInt("BookId")
	bol := false
	fmt.Println(m)
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
func (self BorrowBookController) Ajaxdel() {
	var m models.BorrowBook
	m.Id, _ = self.GetInt("id")

	bol := m.Delete()
	if bol {
		self.ajaxMsg("", MSG_OK)
	} else {
		self.ajaxMsg("删除失败", MSG_ERR)
	}
}
