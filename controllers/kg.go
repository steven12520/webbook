package controllers

import (
	"../models"
	"fmt"
	"../http"
)
type KGController struct {
	BaseController
}

func (self *KGController) KG() {

	self.Data["pageTitle"] = "快估补照片"
	self.display()
}

func (self *KGController) GetRes()  {

	Taskid,_:=self.GetInt("Taskid")

	var p models.TaskCarBasicModel
	p.Id=Taskid
	Reconsideration:= p.GetRes()
	list:= p.GetList()
	if len(list)>0 {
		if list[0].Status!=6 {
			Reconsideration=0
		}
	}else {
		Reconsideration=0
	}

	self.ajaxList("成功", MSG_OK, 0, Reconsideration)
}


func (self *KGController) Commit()  {

	Taskid,_:=self.GetInt("Taskid")
	if Taskid==0 {
		self.ajaxList("成功", MSG_ERR, 0, 0)
		return
	}
	userid:=0
	fmt.Print(Taskid)
	var p models.TaskCarBasicModel
	p.Id=Taskid
	list:= p.GetList()
	if len(list)>0 {
		userid=list[0].CreateUserId
	}else {
		self.ajaxList("成功", MSG_ERR, 0, 0)
		return
	}
	d:= httpdate.SendPostKG(Taskid,userid)
	if d>0 {
		self.ajaxList("成功", MSG_OK, 0, d)
		return
	}else {
		self.ajaxList("成功", MSG_ERR, 0, d)
		return
	}

}

