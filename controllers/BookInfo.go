package controllers

import (
	"../models"
	"fmt"
	"time"
)

type BookInfoController struct {
	BaseController
}

func (self *BookInfoController) BookInfoList() {
	self.Data["pageTitle"] = "图书管理"
	self.display()
}

func (self *BookInfoController) Add() {
	self.Data["pageTitle"] = "新建图书"
	self.display()
}

func (self *BookInfoController) Getbookinfolist() {

	page, pe := self.GetInt("page")
	if pe != nil {
		page = 1
	}

	limit, le := self.GetInt("limit")
	if le != nil {
		limit = 10
	}
	self.pageSize = limit
	bookList := models.BookInfo{}.GetBookList(page, limit)
	count := models.BookInfo{}.GetBookCount()
	if le != nil {
		limit = 10
	}
	list := make([]map[string]interface{}, 0)

	for _, b := range bookList {
		m := make(map[string]interface{})
		m["Id"] = b.Id
		m["CreateTime"] = b.CreateTime
		m["Money"] = b.Money
		m["ClassificationId"] = b.ClassificationId
		m["Author"] = b.Author
		m["PublishedTime"] = b.PublishedTime
		m["BookName"] = b.BookName
		list = append(list, m)
	}

	self.ajaxList("成功", MSG_OK, count, list)

}

func (self *BookInfoController) AjaxSave() {

	id, _ := self.GetInt("id")
	var b models.BookInfo
	b.Id = id
	b.BookName = self.GetString("bookName")
	b.PublishedTime = self.GetString("publishedTime")
	b.Author = self.GetString("author")
	b.ClassificationId, _ = self.GetInt("classificationId")
	b.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	b.Money, _ = self.GetFloat("money")
	bol := false
	if b.Id > 0 {
		book, _ := b.GetBookByNameNoId()

		if book.Id > 0 {
			self.ajaxMsg("保存失败,图书名重复!", MSG_ERR)
			return
		}

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

func (self *BookInfoController) Edit() {
	id, _ := self.GetInt("id")

	var b models.BookInfo
	b.Id = id
	fmt.Println(b)
	book, _ := b.GetbookById()

	row := make(map[string]interface{})
	row["Id"] = book.Id
	row["Money"] = book.Money
	row["ClassificationId"] = book.ClassificationId
	row["Author"] = book.Author
	row["PublishedTime"] = book.PublishedTime
	row["BookName"] = book.BookName
	self.Data["Book"] = row
	self.display()
}

func (self *BookInfoController) Ajaxdel() {

	id, _ := self.GetInt("id")
	var b models.BookInfo
	b.Id = id
	bol := b.Delete()
	if bol {
		self.ajaxMsg("删除成功", MSG_OK)
	} else {
		self.ajaxMsg("删除失败", MSG_ERR)
	}
}
