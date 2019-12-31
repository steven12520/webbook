package controllers

import (
	"../http"
	"github.com/astaxie/beego"
	"strconv"
	"encoding/json"
	"../models"
	"fmt"
	"time"
	"github.com/astaxie/beego/logs"
	"../common"
	"strings"
)
type YstestController struct {
	BaseController
}

var CityA = []string{"京A","京B","京C","京D","冀A","冀B","冀C","冀D"}

func (self *YstestController) Ystest()  {
	GetOperateRecord(168204)
	self.Data["pageTitle"]="预审测试"
	self.display()
}

func (self *YstestController) Plays()  {

	Useridlist:=self.GetString("Useridlist")

	Ranges,_:=self.GetInt("Ranges")
	Usercount,_:=self.GetInt("Usercount")

	if Useridlist=="" || Ranges==0  {
		self.ajaxMsg("参数错误",MSG_ERR)
	}
	arr:=strings.Split(strings.TrimRight(Useridlist,",") ,",")
	for _,Userid:=range arr  {
		id,_:=strconv.Atoi(Userid)
		if Ranges==1 {//预审通过提交
			logs.Debug("操作类型预审通过提交。。。。。。。。。")
			go YSPassO(id,Usercount)
		}else if Ranges==2{//退回修改
			logs.Debug("操作类型退回修改。。。。。。。。。")
			go YSReturnUpdate(id,Usercount)
		}else if Ranges==3{//关闭订单
			logs.Debug("操作类型关闭订单。。。。。。。。。")
			go ClossOrder(id,Usercount)
		}else if Ranges==4{//机构审批
			logs.Debug("操作类型机构审批。。。。。。。。。")
			go SourceSP(id,Usercount)
		}
	}

	self.ajaxMsg("成功",MSG_OK)
}

//预审通过完成提交
func YSPassO(Userid,Usercount int) {
	for i := 0; i < 5; i++ {

		//1，判断是否是接单状态，不是则改成接单,
		thisstart:
			m, b := PretrialPush(Userid)
		if b && m.Status == 100 {
			if m.Data.UserReceiptOrderStatus == 0 {
				logs.Debug("接单中...")
			} else {
				w, er := Working(Userid, 1)
				if er && w.Status == 100 {
					logs.Debug("开始接单成功")
				} else {
					logs.Error("开始接单失败")
					fmt.Println("结束...")
					return
				}
			}
		}else {
			logs.Error("开始接单失败",m.Msg)
		}
		var temporder models.PretrialOrderInfo
		//2,优先级 退回，预审中，新订单
		if len(m.Data.Returnorder) > 0 {
			temporder = m.Data.Returnorder[0]
		} else if len(m.Data.Selforder) > 0 {
			temporder = m.Data.Selforder[0]
		} else if len(m.Data.Neworder) > 0 {
			ReciveOrderConfirm(m.Data.Neworder[0].TaskID, Userid)
			goto thisstart
		} else { //无订单是停止3秒继续
			time.Sleep(time.Second * 2)
			goto thisstart
		}
		if temporder.Status == 4 { //挂起
			UnlockForkedOrder(temporder.TaskID, Userid)
			goto thisstart
		}
		infoDate := TestInfoDate(temporder.Vin, temporder.TaskID, Userid)
		if !infoDate {
			logs.Error("TestInfoDate 错误停止所有执行")
			return
		}
		//3，审核图片
		imgDate, bol := GetImgList(Userid, temporder.TaskID)
		if bol && len(imgDate.Data.CarPicList) > 0 {
			for _, list := range imgDate.Data.CarPicList {
				itemid := list.ItemId
				p:= int64(len(imgDate.Data.CarPicList))
				r:= common.RandInt64(1,p)
				if r%2==0 {
					date, bol := UploadPic(temporder.TaskID,list.Id,itemid)
					if !bol || date.Status != 100 {
						logs.Error("替换错误", temporder.TaskID,list.Id,itemid)
						return
					}
				}
				date, bol := ImgCheckPass(temporder.TaskID, Userid, itemid, 0)
				if !bol || date.Status != 100 {
					logs.Error("图片审核通过出现错误", temporder.TaskID, Userid, itemid, 0)
					return
				}

			}
		}
		//3，审核视频
		if imgDate.Data.VedioInfo.Path != "" {
			date, bol := ImgCheckPass(temporder.TaskID, Userid, -1, 1)
			if !bol || date.Status != 100 {
				logs.Error("图片审核通过出现错误", temporder.TaskID, Userid, -1, 1)
				return
			}
		}
		//3，保存基本信息
		date, bol := SaveFormData(imgDate.Data.Tc, Userid)
		if !bol || date.Status != 100 {
			logs.Error("图片审核通过出现错误", temporder.TaskID, Userid, -1, 1)
			return
		}
		date, bol = PretrailSubmitPass(temporder.TaskID, Userid)
		if !bol || date.Status != 100 {
			logs.Error("预审通过提交", temporder.TaskID, Userid, -1, 1)
			return
		}
	}
	fmt.Println("执行结束。。。。。。。。。")
}
//预审退回
func YSReturnUpdate(Userid,Usercount int)  {
	for i:=0;i<Usercount ;i++ {
		//1，判断是否是接单状态，不是则改成接单,
	thisstart:
		m, b := PretrialPush(Userid)
		if b && m.Status == 100 {
			if m.Data.UserReceiptOrderStatus == 0 {
				logs.Debug("接单中...")
			} else {
				w, er := Working(Userid, 1)
				if er && w.Status == 100 {
					logs.Debug("开始接单成功")
				} else {
					logs.Error("开始接单失败")
					fmt.Println("结束...")
					return
				}
			}
		} else {
			logs.Error("开始接单失败", m.Msg)
		}
		var temporder models.PretrialOrderInfo
		//2,优先级 退回，预审中，新订单
		if len(m.Data.Returnorder) > 0 {
			temporder = m.Data.Returnorder[0]
		} else if len(m.Data.Selforder) > 0 {
			temporder = m.Data.Selforder[0]
		} else if len(m.Data.Neworder) > 0 {
			ReciveOrderConfirm(m.Data.Neworder[0].TaskID, Userid)
			goto thisstart
		} else { //无订单是停止3秒继续
			time.Sleep(time.Second * 3)
			goto thisstart
		}
		if temporder.Status == 4 { //挂起
			UnlockForkedOrder(temporder.TaskID, Userid)
			goto thisstart
		}
		infoDate := TestInfoDate(temporder.Vin, temporder.TaskID, Userid)
		if !infoDate {
			logs.Error("TestInfoDate 错误停止所有执行")
			return
		}
		//3，审核图片
		imgDate, bol := GetImgList(Userid, temporder.TaskID)
		if bol && len(imgDate.Data.CarPicList) > 0 {
			for _, list := range imgDate.Data.CarPicList {
				//查看图片详情
				ImgDetail, bol := GetImgDetail(temporder.TaskID, list.ItemId, Userid)
				if !bol {
					logs.Error("获取图片详情错误", temporder.TaskID, list.ItemId, Userid)
					return
				}
				itemid := list.ItemId
				p := int64(len(imgDate.Data.CarPicList))
				r := common.RandInt64(1, p)
				if r%2 == 0 {
					//替换图片
					date, bol := UploadPic(temporder.TaskID, list.Id, itemid)
					if !bol || date.Status != 100 {
						logs.Error("替换错误", temporder.TaskID, list.Id, itemid)
						return
					}
				}
				r = common.RandInt64(1, p)
				if r%2 == 0 {
					date, bol := ImgCheckPass(temporder.TaskID, Userid, itemid, 0)
					if !bol || date.Status != 100 {
						logs.Error("图片审核通过出现错误", temporder.TaskID, Userid, itemid, 0)
						return
					}
				} else {
					date, bol := ImgCheckUnPass(temporder.TaskID, Userid, itemid, 0, ImgDetail)
					if !bol || date.Status != 100 {
						logs.Error("图片审不核通过出现错误", temporder.TaskID, Userid, itemid, 0)
						return
					}
				}
				//图片审核通过
			}
		}
		//3，审核视频
		if imgDate.Data.VedioInfo.Path != "" {
			date, bol := ImgCheckPass(temporder.TaskID, Userid, -1, 1)
			if !bol || date.Status != 100 {
				logs.Error("图片审核通过出现错误", temporder.TaskID, Userid, -1, 1)
				return
			}
		}
		//3，保存基本信息
		date, bol := SaveFormData(imgDate.Data.Tc, Userid)
		if !bol || date.Status != 100 {
			logs.Error("图片审核通过出现错误", temporder.TaskID, Userid, -1, 1)
			return
		}
		//4，添加附加照片
		date, bol = YsyReturnSummaryReason_Save(temporder.TaskID, Userid)
		if !bol || date.Status != 100 {
			logs.Error("添加附加图片错误", temporder.TaskID, Userid)
			return
		}
		//5,退回修改入缓存
		date, bol = Ys_ImgCheckUnPassSummaryRemark(temporder.TaskID, Userid)

		if !bol || date.Status != 100 {
			logs.Error("退回修改提存入缓存失败", temporder.TaskID, Userid)
			return
		}
		//6,退回修改最终提交
		date, bol = PretrailSubmitBack(temporder.TaskID, Userid)

		if !bol || date.Status != 100 {
			logs.Error("退回修改最终提交失败", temporder.TaskID, Userid)
			return
		}
		logs.Debug("退回修改提交成功", temporder.TaskID, Userid)
	}
}
//关闭订单
func ClossOrder(Userid,Usercount int)  {
	for i:=0;i<Usercount ;i++ {
		//1，判断是否是接单状态，不是则改成接单,
	thisstart:
		m, b := PretrialPush(Userid)
		if b && m.Status == 100 {
			if m.Data.UserReceiptOrderStatus == 0 {
				logs.Debug("接单中...")
			} else {
				w, er := Working(Userid, 1)
				if er && w.Status == 100 {
					logs.Debug("开始接单成功")
				} else {
					logs.Error("开始接单失败")
					fmt.Println("结束...")
					return
				}
			}
		} else {
			logs.Error("开始接单失败", m.Msg)
		}
		var temporder models.PretrialOrderInfo
		//2,优先级 退回，预审中，新订单
		if len(m.Data.Returnorder) > 0 {
			temporder = m.Data.Returnorder[0]
		} else if len(m.Data.Selforder) > 0 {
			temporder = m.Data.Selforder[0]
		} else if len(m.Data.Neworder) > 0 {
			ReciveOrderConfirm(m.Data.Neworder[0].TaskID, Userid)
			goto thisstart
		} else { //无订单是停止3秒继续
			time.Sleep(time.Second * 3)
			goto thisstart
		}
		if temporder.Status == 4 { //挂起
			UnlockForkedOrder(temporder.TaskID, Userid)
			goto thisstart
		}
		infoDate := TestInfoDate(temporder.Vin, temporder.TaskID, Userid)
		if !infoDate {
			logs.Error("TestInfoDate 错误停止所有执行")
			return
		}
		//3，审核图片
		imgDate, bol := GetImgList(Userid, temporder.TaskID)
		count := len(imgDate.Data.RejectReasons)
		txt := ""

		if bol && count > 0 {
			r := common.RandInt64(0, int64(count)-1)
			txt = imgDate.Data.RejectReasons[r].Reason
		}
		txt += ",自动关闭备注"
		date, b := PretrailSubmitReject(temporder.TaskID, Userid, txt)
		if !b && date.Status != 100 {
			logs.Error("订单关闭失败。。", temporder.TaskID, Userid, txt)
		}

		logs.Debug("订单关闭操作结束", temporder.TaskID, Userid, txt)
	}
}
//机构审批
func SourceSP(Userid,Usercount int)  {

	for i:=0;i<Usercount ;i++ {
		//1，判断是否是接单状态，不是则改成接单,
	thisstart:
		m, b := PretrialPush(Userid)
		if b && m.Status == 100 {
			if m.Data.UserReceiptOrderStatus == 0 {
				logs.Debug("接单中...")
			} else {
				w, er := Working(Userid, 1)
				if er && w.Status == 100 {
					logs.Debug("开始接单成功")
				} else {
					logs.Error("开始接单失败")
					fmt.Println("结束...")
					return
				}
			}
		} else {
			logs.Error("开始接单失败", m.Msg)
			return
		}
		var temporder models.PretrialOrderInfo
		//2,优先级 退回，预审中，新订单
		if len(m.Data.Returnorder) > 0 {
			temporder = m.Data.Returnorder[0]
		} else if len(m.Data.Selforder) > 0 {
			temporder = m.Data.Selforder[0]
		} else if len(m.Data.Neworder) > 0 {
			ReciveOrderConfirm(m.Data.Neworder[0].TaskID, Userid)
			goto thisstart
		} else { //无订单是停止3秒继续
			time.Sleep(time.Second * 3)
			goto thisstart
		}
		if temporder.Status == 4 { //挂起
			UnlockForkedOrder(temporder.TaskID, Userid)
			goto thisstart
		}
		infoDate := TestInfoDate(temporder.Vin, temporder.TaskID, Userid)
		if !infoDate {
			logs.Error("TestInfoDate 错误停止所有执行")
			return
		}
		//3，审核图片
		imgDate, bol := GetImgList(Userid, temporder.TaskID)
		count := len(imgDate.Data.RejectReasons)
		txt := ""

		if bol && count > 0 {
			r := common.RandInt64(0, int64(count)-1)
			txt = imgDate.Data.RejectReasons[r].Reason
		} else {
			logs.Error("订单审批操错误。。", temporder.TaskID, Userid, txt)
		}
		txt += ",自动审批备注"
		date, b := PretrailSubmitBackOrg(temporder.TaskID, Userid, txt)
		if !b && date.Status != 100 {
			logs.Error("订单审批失败。。", temporder.TaskID, Userid, txt)
		}
		logs.Debug("订单审批操作结束", temporder.TaskID, Userid, txt)
	}
}
//获取图片列表
func GetImgList(userid,taskid int) (models.ResultDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx?"
	m:=make(map[string]string,0)
	m["op"]="GetImgList"
	m["taskId"]=strconv.Itoa(taskid)
	m["userId"]=strconv.Itoa(userid)

	res,b:= httpdate.SendPost(url,m,"")
	var Data models.ResultDate
	if b {
		fmt.Println(string(res))
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

	res,b:= httpdate.SendPost(url,m,"")
	var Data models.ResultDate
	if b {
		err:= json.Unmarshal(res,&Data)
		if err!=nil {
			logs.Error(string(res))
			return Data,false
		}
	}
	if Data.Status==100 {
		return Data,true
	}else {
		logs.Error(string(res))
		return Data,false
	}
}
//拉取派单数据
func PretrialPush(userid int)(models.PretrialPush,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialPush.ashx"
	m:=make(map[string]string,0)
	m["userId"]=strconv.Itoa(userid)
	res,b:= httpdate.SendPost(url,m,"")
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
	res,b:= httpdate.SendPost(url,m,"")
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

	res,b:= httpdate.SendPost(url,m,"")
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

	res,b:= httpdate.SendPost(url,m,"")
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
func GetOrderInfo(taskid,userId int)(models.OrderInfoModelDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"
	m:=make(map[string]string,0)
	m["op"]="GetOrderInfo"
	m["taskId"]=strconv.Itoa(taskid)
	m["userId"]=strconv.Itoa(userId)

	res,b:= httpdate.SendPost(url,m,"")
	var Data models.OrderInfoModelDate
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
//获取省市列表
func GetProvincesAndCitys(taskid,userId int)(models.ProvincesAndCitysVoDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"
	m:=make(map[string]string,0)
	m["op"]="GetProvincesAndCitys"

	res,b:= httpdate.SendPost(url,m,"")
	var Data models.ProvincesAndCitysVoDate
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
//根据车牌定位城市
func GetCityAndProvinceByPlatName(plateName string)(models.ProvinceCityModelDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"
	m:=make(map[string]string,0)
	m["op"]="GetCityAndProvinceByPlatName"
	m["plateName"]=plateName
	res,b:= httpdate.SendPost(url,m,"")
	var Data models.ProvinceCityModelDate
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
//图片和视频审核通过
func ImgCheckPass(taskId,userId,itemId,video int)(models.ResultPublicDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"

	m:=make(map[string]string,0)
	m["op"]="ImgCheckPass"
	m["taskId"]=strconv.Itoa(taskId)
	m["userId"]=strconv.Itoa(userId)
	m["itemId"]=strconv.Itoa(itemId)//-1为视频
	if video==1 {//视频
		m["itemId"]="-1"//-1为视频
	}


	res,b:= httpdate.SendPost(url,m,"")
	var Data models.ResultPublicDate
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
//基本信息保存
func SaveFormData(model models.TaskCarBasicEPModel,userId int)(models.ResultPublicDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"
	var Data models.ResultPublicDate
	m:=make(map[string]string,0)
	m["op"]="SaveFormData"
	m["taskId"]=strconv.Itoa(model.Id)
	m["userId"]=strconv.Itoa(userId)
	model=GetSaveFormData(model)

	if model.TaskType==1 || model.TaskType==5 {
		c := common.RandInt64(0, int64(len(CityA))-1)
		date, b := GetCityAndProvinceByPlatName(CityA[c])
		if !b || date.Status != 100 {
			logs.Error("GetCityAndProvinceByPlatName 错误", CityA[c])
			return Data, false
		}
		model.CarLicense = CityA[c]
		this:
		if len(model.CarLicense)<7 {
			model.CarLicense+=strconv.FormatInt(common.RandInt64(0, 9),10)
			goto this
		}

		model.ProName = date.Data.ProvinceName
		model.RegisterProvID = date.Data.ProvinceID
		model.CityName =  date.Data.CityName
		model.RegisterCityID = date.Data.CityId
	}

	bytes, _ := json.Marshal(model)
	m["data"]=string(bytes)

	res,b:= httpdate.SendPost(url,m,"")

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
func GetSaveFormData(m models.TaskCarBasicEPModel) models.TaskCarBasicEPModel  {


	if m.TaskType==1 {//18张
		m.CarLicense = "冀ACJXJX"
		m.CityID = 901
		m.ProName = "河北"
		m.RegisterProvID = 9
		m.CityName = "石家庄"
		m.RegisterCityID = 901
		m.Color = 4
		m.EngineNum = "LIJIANSONG"
		m.IsForTransfer = 0
		m.FirstDate = "2019-12-01"
		m.SecondDate = "2019-12-24"
		m.IsMortgage = 1
		m.Mileage = 180000
		m.RecordBrand = "LIJIANSONG"
		m.RecordDate = "2019-12-01"
		m.ProductionTime = "2015-06-06"
		m.Service = 2
		m.TaskOwnerName = "LIJIANSONG"
		m.TransferCount = 0
	}else if m.TaskType==2 {///6张
		m.Mileage = 60000
	}else if m.TaskType==5{///9张
		m.CarLicense = "冀ACJXJX"
		m.CityID = 901
		m.CityName = "石家庄"
		m.Color = 4
		m.EngineNum = "LIJIANSONG"
		m.Mileage = 180000
		m.ProName = "河北"
		m.RecordBrand = "LIJIANSONG"
		m.RecordDate = "2019-12-01"
		m.ProductionTime = "2015-06-06"
		m.RegisterCityID = 901
		m.RegisterProvID = 9
		m.Service = 2
	}
	return m
	//18 张
	//VIN码：	品牌型号	出厂日期：	登记日期：	发动机号：	使用性质：	车主姓名：	车牌号码：	上牌地区：	车身颜色：
	//过户次数：	是否循环过户： 是  否	第一次时间：	第二次时间	表显里程
	//
	//6张
	//表显里程：
	//
	//9张
	//VIN码：	品牌型号：	出厂日期：	登记日期：	发动机号：	使用性质：	车牌号码：	上牌地区：	车身颜色：	表显里程
}
//预审通过提交
func PretrailSubmitPass(taskId,userId int)(models.ResultPublicDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"
	m:=make(map[string]string,0)
	m["op"]="PretrailSubmitPass"
	m["taskId"]=strconv.Itoa(taskId)
	m["userId"]=strconv.Itoa(userId)

	res,b:= httpdate.SendPost(url,m,"")
	var Data models.ResultPublicDate
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
//解除挂起
func UnlockForkedOrder(taskId,userId int)(models.ResultPublicDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"
	m:=make(map[string]string,0)
	m["op"]="UnlockForkedOrder"
	m["taskId"]=strconv.Itoa(taskId)
	m["userId"]=strconv.Itoa(userId)
	res,b:= httpdate.SendPost(url,m,"")
	var Data models.ResultPublicDate
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
//认领验证
func ReciveOrderCheck(taskId,userId int)(models.ResultPublicDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"
	m:=make(map[string]string,0)
	m["op"]="ReciveOrderCheck"
	m["taskId"]=strconv.Itoa(taskId)
	m["userId"]=strconv.Itoa(userId)
	m["types"]="1"

	res,b:= httpdate.SendPost(url,m,"")
	var Data models.ResultPublicDate
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
//订单认领
func ReciveOrderConfirm(taskId,userId int)(models.ResultPublicDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"

	date, er := ReciveOrderCheck(taskId, userId)
	if !er || date.Status != 100 {
		return date,false
	}
	m:=make(map[string]string,0)
	m["op"]="ReciveOrderConfirm"
	m["taskId"]=strconv.Itoa(taskId)
	m["userId"]=strconv.Itoa(userId)
	res,b:= httpdate.SendPost(url,m,"")
	var Data models.ResultPublicDate
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
//替换图片
func UploadPic(taskId,picId,itemId int)(models.ResultPublicDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"

	m:=make(map[string]string,0)
	m["op"]="UploadPic"
	m["TaskID"]=strconv.Itoa(taskId)
	m["picId"]=strconv.Itoa(picId)
	m["itemId"]=strconv.Itoa(itemId)
	filename:= beego.AppConfig.String("pic.picth")
	res,b:= httpdate.SendPost(url,m,filename)
	var Data models.ResultPublicDate
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
//审核不通过加照片
func Upload_SampleImg(taskId,userId,itemId,returnId int)(models.UploadPic,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"

	m:=make(map[string]string,0)
	m["op"]="Upload_SampleImg"
	m["taskId"]=strconv.Itoa(taskId)
	m["userId"]=strconv.Itoa(userId)
	m["itemId"]=strconv.Itoa(itemId)
	m["returnId"]=strconv.Itoa(returnId)

	filename:= beego.AppConfig.String("pic.picthpic")
	res,b:= httpdate.SendPost(url,m,filename)
	var Data models.UploadPic
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
//审核不通过
func ImgCheckUnPass(taskId,userId,itemId,video int,pic models.GetImgDetailReplyDate)(models.ResultPublicDate,bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"

	//unPassReasonJson: [{"Type":1,"ItemID":270,"ReturnID":253,"ReturnReason":"拍摄位置错误，请重拍","ReturnName":"后排座椅","SampleImg":"http://192.168.0.156/group2/M01/0E/91/wKgAl14GI8qAAYUOAAK-Sixj33g537.jpg","FastDFSBasePath":"","TitleText":"eeeeeeeeeeeee","FileName":"2.jpg","DefaultAttachUrl":"https://image.jingzhengu.com/JCXTImgUrl/SampleImage/eighteen6.png"}]
	var Data models.ResultPublicDate
	m := make(map[string]string, 0)
	m["op"] = "ImgCheckUnPass"
	m["taskId"] = strconv.Itoa(taskId)
	m["userId"] = strconv.Itoa(userId)
	m["itemId"] = strconv.Itoa(itemId) //-1为视频

	len:=len(pic.Data.TxtReturnList)
	if len == 0 {
		return Data,false
	}

	r:= common.RandInt64(1,int64(len))
	mo:=pic.Data.TxtReturnList[r]

	relist := make([]models.TaskReturnLogModel, 0)
	mo.TitleText="预审自动测试提交"
	//可以上传图片
	uploadPic, bol := Upload_SampleImg(taskId, userId, itemId, mo.TaskReturnLogModel.ReturnID)
	if !bol || uploadPic.Status!=100 {
		logs.Error("驳回上传图片报错")
		return Data,false
	}
	mo.TaskReturnLogModel.SampleImg=uploadPic.PicPath
	mo.TaskReturnLogModel.FileName=uploadPic.PicName
	mo.TaskReturnLogModel.FastDFSBasePath=uploadPic.FastDFSBasePath

	relist = append(relist, mo.TaskReturnLogModel)

	t, _ := json.Marshal(relist)
	m["unPassReasonJson"] = string(t)

	if video == 1 { //视频
		m["itemId"] = "-1" //-1为视频
	}

	res, b := httpdate.SendPost(url, m, "")

	if b {
		err := json.Unmarshal(res, &Data)
		if err != nil {
			return Data, false
		}
	}
	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}
//驳回新增附加图片
func YsyReturnSummaryReason_Save(taskId,userId int)(models.ResultPublicDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"

	m:=make(map[string]string,0)
	m["op"]="YsyReturnSummaryReason_Save"
	m["taskId"]=strconv.Itoa(taskId)
	m["userId"]=strconv.Itoa(userId)

	list:=make([]models.AddFJ,0)
	var mo models.AddFJ
	mo.ReturnReason ="新增附加1"
	mo.TitleText  ="新增附加1说明"
	mo.SampleImg =beego.AppConfig.String("pgs.fjurl")
	mo.FileName ="yscstjfu.jpg"
	mo.PicID=0
	list= append(list,mo )
	jso,_:=json.Marshal(list)
	m["data"]=string(jso)

	res,b:= httpdate.SendPost(url,m,"")
	var Data models.ResultPublicDate
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
//退回修改
func Ys_ImgCheckUnPassSummaryRemark(taskId,userId int)(models.ResultPublicDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"

	m:=make(map[string]string,0)
	m["op"]="Ys_ImgCheckUnPassSummaryRemark"
	m["taskId"]=strconv.Itoa(taskId)
	m["userId"]=strconv.Itoa(userId)
	m["remark"]="订单退回修改备注123"

	res,b:= httpdate.SendPost(url,m,"")
	var Data models.ResultPublicDate
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
//退回修改
func PretrailSubmitBack(taskId,userId int)(models.ResultPublicDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"

	m:=make(map[string]string,0)
	m["op"]="PretrailSubmitBack"
	m["taskId"]=strconv.Itoa(taskId)
	m["userId"]=strconv.Itoa(userId)
	m["txt"]="订单退回修改备注123"

	res,b:= httpdate.SendPost(url,m,"")
	var Data models.ResultPublicDate
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
//订单关闭
func PretrailSubmitReject(taskId,userId int,txt string)(models.ResultPublicDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"

	m:=make(map[string]string,0)
	m["op"]="PretrailSubmitReject"
	m["taskId"]=strconv.Itoa(taskId)
	m["userId"]=strconv.Itoa(userId)
	m["txt"]=txt


	res,b:= httpdate.SendPost(url,m,"")
	var Data models.ResultPublicDate
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
//机构审批提交
func PretrailSubmitBackOrg(taskId,userId int,txt string)(models.ResultPublicDate,bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"
	m := make(map[string]string, 0)
	m["op"] = "PretrailSubmitBackOrg"
	m["taskId"] = strconv.Itoa(taskId)
	m["userId"] = strconv.Itoa(userId)
	m["txt"] = txt
	res, b := httpdate.SendPost(url, m, "")
	var Data models.ResultPublicDate
	if b {
		err := json.Unmarshal(res, &Data)
		if err != nil {
			return Data, false
		}
	}
	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}
//订单其他数据测试
func TestInfoDate(vin string,taskid,userId int)bool {
	//历史报告
	date, _ := GetHistoryReports(vin)
	if date.Status != 100 {
		logs.Error("TestInfoDate——GetHistoryReports 错误", vin)
		return false
	}
	//操作历史
	logDate, _ := GetOperateRecord(taskid)
	if logDate.Status != 100 {
		logs.Error("TestInfoDate——GetOperateRecord 错误", taskid)
		return false
	}
	//获取订单基本信息
	modelDate, _ := GetOrderInfo(taskid, userId)
	if modelDate.Status != 100 {
		logs.Error("TestInfoDate——GetOrderInfo 错误", taskid, userId)
		return false
	}
	//获取省市列表
	voDate, _ := GetProvincesAndCitys(taskid, userId)
	if voDate.Status != 100 {
		logs.Error("TestInfoDate——GetOrderInfo 错误", taskid, userId)
		return false
	}
	return true
}
