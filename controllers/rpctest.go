package controllers

import (
	"../models"
	"strconv"
	"../common"
	"../http"

	"github.com/chenhg5/collection"
	"github.com/astaxie/beego/logs"
)
type RpcTestController struct {
	BaseController
}

func (self *RpcTestController) CreateOrders() {
	self.Data["pageTitle"] = "创建订单"
	self.display()
}

//获取用户信息
func (self *RpcTestController) GetPublicUsers()  {

	userid,_:=self.GetInt("userid")

	var p models.PublicUsersModel
	p.ID=userid
	list:= p.GetPublicUsers()

	self.ajaxList("成功", MSG_OK, 0, list)
}

//获取方案
func (self *RpcTestController) GetProgrammeConfiguration()  {

	userid,_:=self.GetInt("userid")

	var p models.ProgrammeConfigurationModel
	list:= p.Get(strconv.Itoa(userid))

	self.ajaxList("成功", MSG_OK, 0, list)
}
//获取特殊照片
func (self *RpcTestController) GetPicSpecialSet()  {

	userid,_:=self.GetInt("userid")

	var p models.PicSpecialSetModel
	list:= p.GetPicSpecialSet(strconv.Itoa(userid))

	self.ajaxList("成功", MSG_OK, 0, list)
}
//获取产品类型
func (self *RpcTestController) GetUserProductType()  {

	userid,_:=self.GetInt("userid")
	var p models.UserProductTypeModel
	list:= p.GetUserProductType(userid)
	self.ajaxList("成功", MSG_OK, 0, list)
}

func (self *RpcTestController) SaveOrder()  {
	userid,_:=self.GetInt("userid")
	configID,_:=self.GetInt("configID")
	procductlist,_:=self.GetInt("procductlist")
	vin:=self.GetString("vin")
	ordercount,_:=self.GetInt("ordercount")

	isPretrial,_:=self.GetInt("isPretrial")

	gocount,_:=self.GetInt("gocount")

	var usermodel models.PublicUsersModel
	usermodel.ID=userid
	usermodellist:=usermodel.GetPublicUsers()
	sourceid:=0
	if usermodellist != nil && len(usermodellist) > 0 {
		sourceid=usermodellist[0].TaskSourceId
	}
	var order models.OrderinfoModel
	order.Vin=vin
	order.Ordercount=ordercount
	order.Types=common.StrConvertNameByconfigID(configID)+"["+common.StrConvertNameByprocduct(procductlist)+"]"
	order.CreateName="下单"
	order.Gocount=gocount
	bool :=order.Save()
	if !bool || order.Id==0  {
		self.ajaxMsg("失败", MSG_ERR)
		return
	}
	vinchan:=make(chan string,1000)
	
	go InsertVin(sourceid,vin,ordercount,vinchan)

	for i:=0 ;i<gocount ;i++  {
		go Createorder(isPretrial,procductlist,userid,configID,order.Id,vinchan)
	}

	self.ajaxMsg("成功", MSG_OK)
}
//下单
func Createorder(isPretrial,procductlist,userid,configID int,orderId int64,vinc chan string)  {

	for {
		vinnew, ok := <-vinc
		if ok == false {
			logs.Debug("当前进程下单结束")
			break
		}

		if procductlist == 11 || procductlist == 13 || procductlist == 14 {//快估
			httpdate.Fast(userid,procductlist,vinnew,isPretrial,orderId)
		}else {//非快估
			if configID==5 {
				httpdate.SendPostFormFile9(userid,configID,procductlist,vinnew,orderId)
			}else if configID==2 {
				httpdate.SendPostFormFile6(userid,configID,procductlist,vinnew,orderId)
			}else {
				httpdate.SendPostFormFile(userid,configID,procductlist,vinnew,orderId)
			}
		}
	}

}
//生成vin
func InsertVin(sourceid int,vin string,ordercount int, cvin chan string)  {

	var tc models.TaskCarBasicModel
	vinlist:=make([]string,0)
	for i:=0;i<ordercount ;i++  {
		reset:
		vinnew:=common.GetRandvin(vin)
		if vinnew=="" {
			ordercount++
			continue
		}
		if collection.Collect(vinlist).Contains(vinnew) {
			goto reset
		}
		counts:= tc.IsVinRepeat(vinnew,sourceid)
		if counts>0 {
			goto reset
		}
		vinlist=append(vinlist,vinnew)
		cvin<-vinnew
	}
	close(cvin)
	logs.Debug("InsertVin 结束")
}