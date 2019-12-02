package controllers

import (
	"../models"
	"time"
)

type UserInfoController struct {
	BaseController
}

func (self *UserInfoController) Userinfolist() {

	self.Data["pageTitle"] = "用户管理"

	self.display()
}

//保存用户信息
func (self *UserInfoController) AjaxSave() {

	id, _ := self.GetInt("b_id")

	var user models.UserInfo
	user.B_id = id
	user.B_createTime = time.Now().Format("2006-01-02 15:04:05")
	user.B_idCar = self.GetString("b_idCar")
	user.B_address = self.GetString("b_address")
	user.B_telephone = self.GetString("b_telephone")
	user.B_type, _ = self.GetInt("b_type")
	user.B_role, _ = self.GetInt("b_role")
	user.B_pwd = self.GetString("pwd")
	user.B_loginname = self.GetString("loginname")

	u, _ := models.UserInfoGetByNameNoId(user.B_loginname, id)

	if u.B_id > 0 {
		self.ajaxMsg("登录名已存在！", MSG_ERR)
		return
	}
	var bol bool
	if id == 0 { //新建
		bol = models.UserInfo_Save(user)
	} else { //更新
		bol = models.UserInfo_Edite(user)
	}

	if bol {
		self.ajaxMsg("", MSG_OK)
	} else {
		self.ajaxMsg("保存失败", MSG_ERR)
	}
}

//删除用户信息
func (self *UserInfoController) AjaxDel() {
	id, _ := self.GetInt("id")

	u, err := models.AdminGetById(id)
	if err != nil && u.B_id == 0 {
		self.ajaxMsg("删除失败，用户不存在！", MSG_ERR)
	}
	bol := models.UserInfo_Del(id)
	if bol {
		self.ajaxMsg("", MSG_OK)
	} else {
		self.ajaxMsg("操作失败！", MSG_ERR)
	}

}

func (self *UserInfoController) Add() {

	self.Data["pageTitle"] = "新增用户"
	self.display()
}
func (self *UserInfoController) GetUserInfoList() {

	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 20
	}
	self.pageSize = limit

	result := models.UserInfoList_Get(page, self.pageSize)
	count := models.UserInfoCount_Get()
	list := make([]map[string]interface{}, 0)

	for _, v := range result {
		row := make(map[string]interface{})
		row["B_loginname"] = v.B_loginname
		row["B_id"] = v.B_id
		row["B_pwd"] = v.B_pwd
		row["B_type"] = v.B_type
		row["B_createTime"] = v.B_createTime
		row["B_role"] = v.B_role
		row["B_telephone"] = v.B_telephone
		row["B_address"] = v.B_address
		row["B_idCar"] = v.B_idCar
		list = append(list, row)
	}
	self.ajaxList("成功", MSG_OK, count, list)

}

func (self *UserInfoController) Edit() {
	self.Data["pageTitle"] = "编辑用户"

	id, _ := self.GetInt("id", 0)
	u, _ := models.AdminGetById(id)
	row := make(map[string]interface{})
	row["B_id"] = u.B_id
	row["B_idCar"] = u.B_idCar
	row["B_address"] = u.B_address
	row["B_telephone"] = u.B_telephone
	row["B_type"] = u.B_type
	row["B_role"] = u.B_role
	row["pwd"] = u.B_pwd
	row["loginname"] = u.B_loginname
	self.Data["User"] = row
	self.display()
}
