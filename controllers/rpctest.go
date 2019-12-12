package controllers

import (
	"../models"
	"strconv"
	"../common"
	"../http"

	"github.com/chenhg5/collection"
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

	//gocount,_:=self.GetInt("gocount")

	var usermodel models.PublicUsersModel
	usermodel.ID=userid
	usermodellist:=usermodel.GetPublicUsers()
	sourceid:=0
	if usermodellist != nil && len(usermodellist) > 0 {
		sourceid=usermodellist[0].TaskSourceId
	}
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
		vinlist=append(vinlist,vinnew)
		counts:= tc.IsVinRepeat(vinnew,sourceid)
		if counts>0 {
			goto reset
		}
		go func() {
			if procductlist == 11 || procductlist == 13 || procductlist == 14 {//快估
				httpdate.Fast(userid,procductlist,vinnew,isPretrial)
			}else {//非快估
				if configID==5 {
					httpdate.SendPostFormFile9(userid,configID,procductlist,vinnew)
				}else if configID==2 {
					httpdate.SendPostFormFile6(userid,configID,procductlist,vinnew)
				}else {
					httpdate.SendPostFormFile(userid,configID,procductlist,vinnew)
				}
			}
		}()
	}


	self.ajaxMsg("成功", MSG_OK)
}