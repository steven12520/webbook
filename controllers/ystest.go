package controllers

import (
	"../http"
	"github.com/astaxie/beego"
	"strconv"
	"encoding/json"
	"../models"
	"github.com/astaxie/beego/logs"
)
type YstestController struct {
	BaseController
}

func (self *YstestController) Ystest()  {

	self.Data["pageTitle"]="预审测试"
	self.display()
}

func (self *YstestController) Plays()  {


	defer func() {
		v:=recover()
		if v!=nil {
			logs.Error("Plays error ",v)
			self.ajaxMsg("报错了",MSG_ERR)
			return
		}
	}()

	Taskid,_:=self.GetInt("Taskid")
	Userid,_:=self.GetInt("Userid")
	if Taskid==0 || Userid==0 {
		self.ajaxMsg("参数错误",MSG_ERR)
	}

	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"

	m:=make(map[string]string,0)
	m["op"]="GetImgList"
	m["taskId"]=strconv.Itoa(Taskid)
	m["userId"]=strconv.Itoa(Userid)
	res,b:= httpdate.SendPost(url,m)
	var Data models.ResultDate
	if b {
		err:= json.Unmarshal(res,&Data)
		if err!=nil {
			self.ajaxMsg("报错了",MSG_ERR)
			return
		}
	}

	go YSPassO(Data)
	self.ajaxMsg("成功",MSG_OK)
}

func YSPassO(Data models.ResultDate)  {


}