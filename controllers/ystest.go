package controllers

import (
	"../http"
	"github.com/astaxie/beego"
	"strconv"
	"encoding/json"
	"../models"
)
type YstestController struct {
	BaseController
}

func (self *YstestController) Ystest()  {

	self.Data["pageTitle"]="预审测试"
	self.display()
}

func (self *YstestController) Plays()  {


	//defer func() {
	//	v:=recover()
	//	if v!=nil {
	//		logs.Error("Plays error ",v)
	//		self.ajaxMsg("报错了",MSG_ERR)
	//		return
	//	}
	//}()
	//
	Taskid,_:=self.GetInt("Taskid")
	Userid,_:=self.GetInt("Userid")
	m,b:= PretrialPush(Userid)
	if b && m.Status==100 {

	}
	if Taskid==0 || Userid==0 {
		self.ajaxMsg("参数错误",MSG_ERR)
	}

	Data,bol:=GetImgList(Userid,Taskid)
	if bol{
		go YSPassO(Data)
	}

	self.ajaxMsg("成功",MSG_OK)
}

func YSPassO(Data models.ResultDate)  {


}
//获取图片列表
func GetImgList(userid,taskid int) (models.ResultDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"
	m:=make(map[string]string,0)
	m["op"]="GetImgList"
	m["taskId"]=strconv.Itoa(taskid)
	m["userId"]=strconv.Itoa(userid)
	res,b:= httpdate.SendPost(url,m)
	var Data models.ResultDate
	if b {
		err:= json.Unmarshal(res,&Data)
		if err!=nil {
			return Data,false
		}
	}
	if Data.Status==100 {
		return Data,true
	}else {
		return Data,false
	}
}

//开始接单/停止接单
func Working(userid,Working int)(models.ResultDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"
	m:=make(map[string]string,0)
	m["op"]="Working"
	m["Working"]=strconv.Itoa(Working)//1接单，0停止接单
	m["userId"]=strconv.Itoa(userid)
	res,b:= httpdate.SendPost(url,m)
	var Data models.ResultDate
	if b {
		err:= json.Unmarshal(res,&Data)
		if err!=nil {
			return Data,false
		}
	}
	if Data.Status==100 {
		return Data,true
	}else {
		return Data,false
	}
}

//拉取派单数据
func PretrialPush(userid int)(models.PretrialPush,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialPush.ashx?userId=2082"
	m:=make(map[string]string,0)
	//m["userId"]=strconv.Itoa(userid)
	res,b:= httpdate.SendPost(url,m)
	var Data models.PretrialPush
	if b {
		err:= json.Unmarshal(res,&Data)
		if err!=nil {
			return Data,false
		}
	}
	if Data.Status==100 {
		return Data,true
	}else {
		return Data,false
	}
}
//获取历史订单
func GetHistoryReports(vin string)(models.OrderHistoryDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"
	m:=make(map[string]string,0)
	m["vin"]=vin
	m["op"]="GetHistoryReports"
	res,b:= httpdate.SendPost(url,m)
	var Data models.OrderHistoryDate
	if b {
		err:= json.Unmarshal(res,&Data)
		if err!=nil {
			return Data,false
		}
	}
	if Data.Status==100 {
		return Data,true
	}else {
		return Data,false
	}
}
//获取操作历史
func GetOperateRecord (taskid int)(models.OperateLogDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"
	m:=make(map[string]string,0)
	m["op"]="GetOperateRecord"
	m["taskId"]=strconv.Itoa(taskid)

	res,b:= httpdate.SendPost(url,m)
	var Data models.OperateLogDate
	if b {
		err:= json.Unmarshal(res,&Data)
		if err!=nil {
			return Data,false
		}
	}
	if Data.Status==100 {
		return Data,true
	}else {
		return Data,false
	}
}
//获取图片详情
func GetImgDetail(taskid,itemId,userId int )(models.GetImgDetailReplyDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"
	m:=make(map[string]string,0)
	m["op"]="GetImgDetail"
	m["taskId"]=strconv.Itoa(taskid)
	m["itemId"]=strconv.Itoa(itemId)
	m["userId"]=strconv.Itoa(userId)

	res,b:= httpdate.SendPost(url,m)
	var Data models.GetImgDetailReplyDate
	if b {
		err:= json.Unmarshal(res,&Data)
		if err!=nil {
			return Data,false
		}
	}
	if Data.Status==100 {
		return Data,true
	}else {
		return Data,false
	}
}

//获取订单基本信息
func GetOrderInfo(taskid,userId int)()  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"
	m:=make(map[string]string,0)
	m["op"]="GetOrderInfo"
	m["taskId"]=strconv.Itoa(taskid)
	m["userId"]=strconv.Itoa(userId)

	res,b:= httpdate.SendPost(url,m)
	var Data models.GetImgDetailReplyDate
	if b {
		err:= json.Unmarshal(res,&Data)
		if err!=nil {
			return Data,false
		}
	}
	if Data.Status==100 {
		return Data,true
	}else {
		return Data,false
	}
}