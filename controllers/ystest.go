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
func GetOrderInfo(taskid,userId int)(models.OrderInfoModelDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"
	m:=make(map[string]string,0)
	m["op"]="GetOrderInfo"
	m["taskId"]=strconv.Itoa(taskid)
	m["userId"]=strconv.Itoa(userId)

	res,b:= httpdate.SendPost(url,m)
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

	res,b:= httpdate.SendPost(url,m)
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
	res,b:= httpdate.SendPost(url,m)
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


	res,b:= httpdate.SendPost(url,m)
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
func SaveFormData(taskId,userId,itemId,video int)(models.ResultPublicDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"
	m:=make(map[string]string,0)
	m["op"]="SaveFormData"
	m["taskId"]=strconv.Itoa(taskId)
	m["userId"]=strconv.Itoa(userId)
	model:=GetSaveFormData()
	bytes, _ := json.Marshal(model)
	m["data"]=string(bytes)

	res,b:= httpdate.SendPost(url,m)
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

func GetSaveFormData() models.TaskCarBasicEPModel  {

	var m models.TaskCarBasicEPModel
	m.TaskOwnerName="eeeeee"
	m.ShowArea=""
	m.ShowArea_v=""
	m.ProgramId=""
	m.Id=168083
	m.OrderNo="JZG9011576745749017"
	m.SourceID=204
	m.CityID=901
	m.Des="订单备注"
	m.LikeMan="Zjzj"
	m.LikeTel="15313666764"
	m.LikeAddr=""
	m.Vin="LJS1219F6W7CVE7K6"
	m.CarLicense="冀ACJXJX"
	m.RecordBrand="Xjxjjxdj"
	m.EngineNum="Xjxjjxjddj"
	m.RecordDate="2019-12-01T00:00:00"
	m.Color=0
	m.Mileage=12222
	m.Service=2
	m.AssessmentPrace=0
	m.SalePrice=0
	m.AssessmentDes=""
	m.UserID=0
	m.Status=7
	m.StatusName="预审中"
	m.OrderStatus=0
	m.CreateTime="2019-12-19T16:55:49"
	m.UpdateTime="2019-12-23T10:47:31"
	m.StartTime="2019-12-19T16:55:49"
	m.EndTime="2019-12-24T16:55:49"
	m.Exhaust=""
	m.Seating=0
	m.PerfSeatNum=""
	m.CarType=""
	m.DrivingMode=0
	m.Transmission=0
	m.FuelType=0
	m.ProductionTime="2019-11-01"
	m.Certificates=0
	m.ManufacturerPrice=0
	m.BusinessPrice=0
	m.SetGroupID=211
	m.TaskType=1
	m.TaskBackNum=0
	m.TaskBackReason=""
	m.AppraiseBackNum=0
	m.AppraiseBackReason=""
	m.TransferCount=0
	m.Insurance="1900-01-01T00:00:00"
	m.Inspection="1900-01-01T00:00:00"
	m.CreateUserId=1726
	m.ProvID=9
	m.ProName="河北"
	m.CityName="石家庄"
	m.YXOrderNo="JZG9011576745749017"
	m.VideoPath="group2/M01/0E/87/wKgAmF37qsGAU1rsAAkSzqdK-ms354.mp4"
	m.RegisterProvID=2
	m.RegisterProvName=""
	m.RegisterCityID=201
	m.RegisterCityname=""
	m.SourceName="测试机构-石家庄（勿动）"
	m.Perf_DriveType=""
	m.TransmissionType=""
	m.Engine_Exhaust=""
	m.Fuel=""
	m.Tasktel=""
	m.ProductType=9
	m.AppraiseBackReasonNew=""
	m.ShowSourceName=1
	m.ProgrammeId=4
	m.IsComplete=0
	m.ReViewType=0
	m.ReportPcLink=""
	m.ReportMLink=""
	m.ReportPrintLink=""
	m.AutoStar=0
	m.EstimatedTime="0001-01-01T00:00:00"
	m.TaskVersion=3
	m.OrderTelphone=""
	m.PretrialUser=2082
	m.CarFullName=""
	m.IsXing=0
	m.Channel=0
	m.CreateOrderName="csfyxsjzfb1(李丽-测试非易鑫分部石家庄1)"
	m.IsForTransfer=1
	m.FirstDate="2019-12-01"
	m.SecondDate="2019-12-24"
	m.AccidentBasis=""
	m.AccidentBasisType=""
	m.RandomYS=0
	m.RandomPGS=0
	m.IdUserName=""
	m.IdNumber=""
	m.IsMortgage=1
	m.ScrapValue=0
	m.MaintainStatus=0
	m.AssessmentPraceScopeStart=0
	m.AssessmentPraceScopeEnd=0
	m.SalePriceScopeStart=0
	m.SalePriceScopeEnd=0
	m.PgsSalePrice=0
	m.NewEdition=0
	m.Reconsideration=0

	return m
}
//预审通过提交
func PretrailSubmitPass(taskId,userId int)(models.ResultPublicDate,bool)  {
	url:= beego.AppConfig.String("pgs.url")+"/APP/Pretrial/PretrialV2.ashx"
	m:=make(map[string]string,0)
	m["op"]="PretrailSubmitPass"
	m["taskId"]=strconv.Itoa(taskId)
	m["userId"]=strconv.Itoa(userId)

	res,b:= httpdate.SendPost(url,m)
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



