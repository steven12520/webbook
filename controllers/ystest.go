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

var CityA = []string{"京A", "京B", "京C", "京D", "冀A", "冀B", "冀C", "冀D"}

func (self *YstestController) Ystest() {


	//DeletePic(170389,5809422,0,0)
	//Working(2083,1,0,0)
	//UploadPic(170389, 0, 0 , 0, 0 )
	//UploadPic(170389, 0, 0 , 0, 0 )
	//DeletePic(170389,5809377,0,0)
	//GetOrderInfo(1917299,10269,0,0)
	self.Data["pageTitle"] = "预审测试"
	self.display()
}

func (self *YstestController) Plays() {

	defer func() {
		i := recover()
		if i != nil {
			common.Requestdderror("Plays error"+i.(string))
		}
	}()

	Useridlist := self.GetString("Useridlist")

	Ranges, _ := self.GetInt("Ranges")
	Usercount, _ := self.GetInt("Usercount")
	timelen, _ := self.GetInt("timelen")

	if Useridlist == "" || Ranges == 0 || Usercount == 0 || timelen == 0 {
		logs.Debug("参数错误", Useridlist, Ranges, Usercount, timelen)
		self.ajaxMsg("参数错误", MSG_ERR)
		return
	}
	var mo models.Ysyinfo
	mo.Timelength = timelen
	mo.Types = Ranges
	mo.Avgcount = Usercount
	mo.Username = Useridlist
	if Ranges == 1 {
		mo.Yname = "预审通过测试"
	} else if Ranges == 2 {
		mo.Yname = "预审退回修改测试"
	} else if Ranges == 3 {
		mo.Yname = "预审关闭订单测试"
	} else {
		mo.Yname = "预审机构审批测试"
	}
	bol := mo.Add()
	if !bol && mo.Id == 0 {
		return
	}

	arr := strings.Split(strings.TrimRight(Useridlist, ","), ",")
	for _, Userid := range arr {
		id, _ := strconv.Atoi(Userid)
		if Ranges == 1 { //预审通过提交
			logs.Debug("操作类型预审通过提交。。。。。。。。。")
			go YSPassO(id, Usercount, timelen, mo.Id)
		} else if Ranges == 2 { //退回修改
			logs.Debug("操作类型退回修改。。。。。。。。。")
			go YSReturnUpdate(id, Usercount, timelen, mo.Id)
		} else if Ranges == 3 { //关闭订单
			logs.Debug("操作类型关闭订单。。。。。。。。。")
			go ClossOrder(id, Usercount, timelen, mo.Id)
		} else if Ranges == 4 { //机构审批
			logs.Debug("操作类型机构审批。。。。。。。。。")
			go SourceSP(id, Usercount, timelen, mo.Id)
		}
	}
	fmt.Println(Usercount, timelen, arr)
	self.ajaxMsg("成功", MSG_OK)
}

//预审通过完成提交
func YSPassO(Userid, Usercount, timelen int, Ysyid int64) {

	endtime := time.Now().Add(time.Minute * time.Duration(timelen))
	var ddid int64
	ddvin:=""

	defer func(ddid int64) {
		var mo models.Ysyinfodetail
		mo.Id=ddid
		if err := recover(); err != nil {
			mo.Satus = 0
			mo.Satusmsg = err.(string)
			mo.Update()
			common.Requestdderror(strconv.Itoa(Userid)+":"+err.(string))
			return
		}
	}(ddid)

	m, _ := PretrialPush(Userid, Ysyid, 0)
	if  m.Status == 100 {
		if m.Data.UserReceiptOrderStatus == 0 {
			logs.Debug("接单中...")
		} else {
			w, er := Working(Userid, 1, Ysyid, 0)
			if er && w.Status == 100 {
				logs.Debug("开始接单成功")
			} else {
				logs.Error("开始接单失败")
				fmt.Println("结束...")
				panic("开始接单失败")
			}
		}
	} else {
		logs.Error("开始接单失败", m.Msg)
		panic("开始接单失败"+m.Msg)
	}

	for i := 0; i < Usercount; i++ {
		//1，判断是否是接单状态，不是则改成接单,
	thisstart:
		m, _ = PretrialPush(Userid, Ysyid, 0)

		var temporder models.PretrialOrderInfo
		//2,优先级 退回，预审中，新订单
		if len(m.Data.Returnorder) > 0 {
			temporder = m.Data.Returnorder[0]
		} else if len(m.Data.Selforder) > 0 {
			temporder = m.Data.Selforder[0]
		} else if len(m.Data.Neworder) > 0 {
			date, _ := ReciveOrderConfirm(m.Data.Neworder[0].TaskID, Userid, Ysyid, 0)
			if date.Status != 100 && date.Msg == "操作失败此订单已被操作！" {
				Deleteorder(m.Data.Neworder[0].Vin)
			}
			goto thisstart
		} else { //无订单是停止3秒继续
			if endtime.Before(time.Now()) {
				logs.Debug("时间已到结束。。。")
				break
			}

			time.Sleep(time.Second * 2)
			logs.Debug("无订单停止2秒继续。。。")
			goto thisstart
		}

		var mo models.Ysyinfodetail
		mo.Ysyid = Ysyid
		mo.Userid = Userid
		mo.Satus = 0
		mo.Satusmsg = ""
		mo.Vin = temporder.Vin
		if temporder.Vin!=ddvin{
			mo.Add()
			if mo.Id == 0 {
				return
			}
			ddid=mo.Id
			ddvin=temporder.Vin
		}else {
			mo.Id=ddid
			temporder.Vin=ddvin
		}

		if temporder.Status == 4 { //挂起
			date, _ := UnlockForkedOrder(temporder.TaskID, Userid, Ysyid, mo.Id)
			if date.Status != 100 {
				logs.Error("解除挂起失败", temporder.TaskID, Userid)
				panic(date.Msg)
			}
			goto thisstart
		} else {
			if temporder.TaskID%2 == 0 && common.RandInt64(1, 4) == 1 {
				date, _ := ForkOrder(temporder.TaskID, Userid, Ysyid, mo.Id)
				if date.Status != 100 {
					logs.Error("挂起失败", temporder.TaskID, Userid)
					panic(date.Msg)
				}
				goto thisstart
			}

		}

		infoDate := TestInfoDate(temporder.Vin, temporder.TaskID, Userid, Ysyid, mo.Id)
		if !infoDate {
			logs.Error("TestInfoDate 错误停止所有执行")
			panic("错误停止所有执行TestInfoDate")
		}

		//添加附加图片2张
		publicDate, _ := UploadPic(temporder.TaskID, 0, 0, Ysyid, mo.Id)
		if publicDate.Status!=100 {
			logs.Error("上传附加图失败1", temporder.TaskID, 0, 0, Ysyid, mo.Id)
			panic(publicDate.Msg)
		}
		publicDate, _ = UploadPic(temporder.TaskID, 0, 0, Ysyid, mo.Id)
		if publicDate.Status!=100 {
			logs.Error("上传附加图失败2", temporder.TaskID, 0, 0, Ysyid, mo.Id)
			panic(publicDate.Msg)
		}

		//3，审核图片
		imgDate, bol := GetImgList(Userid, temporder.TaskID, Ysyid, mo.Id)
		l:=len(imgDate.Data.RedisPretrail.AddPicTemp)
		deleteid:=0
		if imgDate.Status==100 && l > 0 {
			for _,k:=range imgDate.Data.RedisPretrail.AddPicTemp  {
				deleteid=k.Id
			}
		}
		resultPublicDate, _ := DeletePic(temporder.TaskID, deleteid, Ysyid, mo.Id)
		if resultPublicDate.Status!=100 {
			logs.Error("DeletePic附加图失败", temporder.TaskID, deleteid, Ysyid, mo.Id)
			panic(resultPublicDate.Msg)

		}
		imgDate, bol = GetImgList(Userid, temporder.TaskID, Ysyid, mo.Id)
		if bol && len(imgDate.Data.CarPicList) > 0 {
			for _, list := range imgDate.Data.CarPicList {
				itemid := list.ItemId
				p := int64(len(imgDate.Data.CarPicList))
				r := common.RandInt64(1, p)
				if r%2 == 0 {
					date, _ := UploadPic(temporder.TaskID, list.Id, itemid, Ysyid, mo.Id)
					if date.Status != 100 {
						logs.Error("替换错误", temporder.TaskID, list.Id, itemid)
						panic("替换错误")
					}
					if itemid%2 == 0 {
						datefj, _ := UploadPic_YsAttach(temporder.TaskID, itemid, Ysyid, mo.Id)
						if datefj.Status != 100 {
							logs.Error("上传附件图片操作记录中显示", temporder.TaskID, itemid)
							panic("上传附件图片操作记录中显示")
						}
						if datefj.PicId%2 == 0 {
							date, _ = DeletePic_YsAttach(temporder.TaskID, datefj.PicId, Ysyid, mo.Id)
							if date.Status != 100 {
								logs.Error("删除附加错误", temporder.TaskID, datefj.PicId)
								panic("删除附加错误")
							}
						}
					}
				}
				date, _ := ImgCheckPass(temporder.TaskID, Userid, itemid, 0, Ysyid, mo.Id)
				if date.Status != 100 {
					logs.Error("图片审核通过出现错误", temporder.TaskID, Userid, itemid, 0)
					panic("图片审核通过出现错误")
				}

			}
		}
		//3，审核视频
		if imgDate.Data.VedioInfo.Path != "" {
			date, bol := ImgCheckPass(temporder.TaskID, Userid, -1, 1, Ysyid, mo.Id)
			if !bol || date.Status != 100 {
				logs.Error("图片审核通过出现错误", temporder.TaskID, Userid, -1, 1)
				panic("图片审核通过出现错误")
			}
		}
		//3，保存基本信息
		date, bol := SaveFormData(imgDate.Data.Tc, Userid, Ysyid, mo.Id)
		if !bol || date.Status != 100 {
			logs.Error("图片审核通过出现错误", temporder.TaskID, Userid, -1, 1)
			panic("图片审核通过出现错误")
		}
		//添加备注
		date, bol = Ys_AddRemark(temporder.TaskID, Userid, "预审自动测试添加备注", Ysyid, mo.Id)
		if !bol || date.Status != 100 {
			logs.Error("图片审添加备注", temporder.TaskID, Userid, -1, 1)
			panic("图片审添加备注")
		}
		date, _ = PretrailSubmitPass(temporder.TaskID, Userid, Ysyid, mo.Id)
		if  date.Status != 100 {
			logs.Error("预审通过提交", temporder.TaskID, Userid, -1, 1)
			goto thisstart
		}
		mo.Satus = 1
		mo.Satusmsg = "成功"
		mo.Update()
		mo.Id = 0

	}
	fmt.Println("执行结束。。。。。。。。。")

	if m.Data.UserReceiptOrderStatus == 0 {
	thisstop:
		w, er := Working(Userid, 0, Ysyid, 0)
		if er && w.Status == 100 {
			logs.Debug("停止接单")
		} else {
			time.Sleep(time.Second * 2)
			goto thisstop
		}
	}

}

//预审退回
func YSReturnUpdate(Userid, Usercount, timelen int, Ysyid int64) {
	endtime := time.Now().Add(time.Minute * time.Duration(timelen))

	var ddid int64
	ddvin:=""

	defer func(ddid int64) {
		var mo models.Ysyinfodetail
		mo.Id=ddid
		if err := recover(); err != nil {
			mo.Satus = 0
			mo.Satusmsg = err.(string)
			mo.Update()
			common.Requestdderror(strconv.Itoa(Userid)+":"+err.(string))
			return
		}
	}(ddid)

	m, b := PretrialPush(Userid, Ysyid, 0)
	if b && m.Status == 100 {
		if m.Data.UserReceiptOrderStatus == 0 {
			logs.Debug("接单中...")
		} else {
			w, er := Working(Userid, 1, Ysyid, 0)
			if er && w.Status == 100 {
				logs.Debug("开始接单成功")
			} else {
				logs.Error("开始接单失败")
				fmt.Println("结束...")
				panic("开始接单失败")
			}
		}
	} else {
		logs.Error("开始接单失败", m.Msg)
		panic("开始接单失败")
	}

	for i := 0; i < Usercount; i++ {
		//1，判断是否是接单状态，不是则改成接单,
	thisstart:
		m, b = PretrialPush(Userid, Ysyid,0)

		var temporder models.PretrialOrderInfo
		//2,优先级 退回，预审中，新订单
		if len(m.Data.Returnorder) > 0 {
			temporder = m.Data.Returnorder[0]
		} else if len(m.Data.Selforder) > 0 {
			temporder = m.Data.Selforder[0]
		} else if len(m.Data.Neworder) > 0 {
			ReciveOrderConfirm(m.Data.Neworder[0].TaskID, Userid, Ysyid, 0)
			goto thisstart
		} else { //无订单是停止3秒继续
			if endtime.Before(time.Now()) {
				logs.Debug("时间已到结束。。。")
				break
			}

			time.Sleep(time.Second * 3)
			goto thisstart
		}
		if temporder.Status == 4 { //挂起
			UnlockForkedOrder(temporder.TaskID, Userid, Ysyid, 0)
			goto thisstart
		}
		var mo models.Ysyinfodetail
		mo.Ysyid = Ysyid
		mo.Userid = Userid
		mo.Satus = 0
		mo.Satusmsg = ""
		mo.Vin = temporder.Vin
		if temporder.Vin!=ddvin{
			mo.Add()
			if mo.Id == 0 {
				return
			}
			ddid=mo.Id
			ddvin=temporder.Vin
		}else {
			mo.Id=ddid
			temporder.Vin=ddvin
		}

		infoDate := TestInfoDate(temporder.Vin, temporder.TaskID, Userid, Ysyid, mo.Id)
		if !infoDate {
			logs.Error("TestInfoDate 错误停止所有执行")
			panic("错误停止所有执行")
		}
		//3，审核图片
		imgDate, bol := GetImgList(Userid, temporder.TaskID, Ysyid, mo.Id)
		if bol && len(imgDate.Data.CarPicList) > 0 {
			for _, list := range imgDate.Data.CarPicList {
				//查看图片详情
				ImgDetail, bol := GetImgDetail(temporder.TaskID, list.ItemId, Userid, Ysyid, mo.Id)
				if !bol {
					logs.Error("获取图片详情错误", temporder.TaskID, list.ItemId, Userid)
					panic("获取图片详情错误")
				}
				itemid := list.ItemId
				p := int64(len(imgDate.Data.CarPicList))
				r := common.RandInt64(1, p)
				if r%2 == 0 && common.RandInt64(1, 6)==2 {
					//替换图片
					date, bol := UploadPic(temporder.TaskID, list.Id, itemid, Ysyid, mo.Id)
					if !bol || date.Status != 100 {
						logs.Error("替换错误", temporder.TaskID, list.Id, itemid)
						panic("替换错误")
					}
				}
				r = common.RandInt64(1, p)
				if r%2 == 0 {
					date, bol := ImgCheckPass(temporder.TaskID, Userid, itemid, 0, Ysyid, mo.Id)
					if !bol || date.Status != 100 {
						logs.Error("图片审核通过出现错误", temporder.TaskID, Userid, itemid, 0)
						panic("图片审核通过出现错误")
					} else {
						date, bol, _ := ImgCheckUnPass(temporder.TaskID, Userid, itemid, 0, ImgDetail, Ysyid, mo.Id)
						if !bol || date.Status != 100 {
							logs.Error("图片审不核通过出现错误", temporder.TaskID, Userid, itemid, 0)
							panic("图片审不核通过出现错误")
						}
					}
				} else {
					date, bol, returnId := ImgCheckUnPass(temporder.TaskID, Userid, itemid, 0, ImgDetail, Ysyid, mo.Id)
					if !bol || date.Status != 100 {
						logs.Error("图片审不核通过出现错误", temporder.TaskID, Userid, itemid, 0)
						panic("图片审不核通过出现错误")
					}
					if itemid%2 == 0  && common.RandInt64(1, 6)==2  {
						date, bol = ImgCheckUnPass_AttachDelete(temporder.TaskID, Userid, itemid, returnId, Ysyid, mo.Id)
						if !bol || date.Status != 100 {
							logs.Error("驳回图片删除出现错误", temporder.TaskID, Userid, itemid, returnId)
							panic("驳回图片删除出现错误")
						}
					}

				}
				//图片审核通过
			}
		}
		//3，审核视频
		if imgDate.Data.VedioInfo.Path != "" {
			date, bol := ImgCheckPass(temporder.TaskID, Userid, -1, 1, Ysyid, mo.Id)
			if !bol || date.Status != 100 {
				logs.Error("图片审核通过出现错误", temporder.TaskID, Userid, -1, 1)
				panic("图片审核通过出现错误")
			}
		}
		//3，保存基本信息
		date, bol := SaveFormData(imgDate.Data.Tc, Userid, Ysyid, mo.Id)
		if !bol || date.Status != 100 {
			logs.Error("图片审核通过出现错误", temporder.TaskID, Userid, -1, 1)
			panic("图片审核通过出现错误")
		}
		//4，添加附加照片
		date, bol = YsyReturnSummaryReason_Save(temporder.TaskID, Userid, Ysyid, mo.Id)
		if !bol || date.Status != 100 {
			logs.Error("添加附加图片错误", temporder.TaskID, Userid)
			panic("添加附加图片错误")
		}
		//5,退回修改入缓存
		date, bol = Ys_ImgCheckUnPassSummaryRemark(temporder.TaskID, Userid, Ysyid, mo.Id)

		if !bol || date.Status != 100 {
			logs.Error("退回修改提存入缓存失败", temporder.TaskID, Userid)
			panic("退回修改提存入缓存失败")
		}
		//6,退回修改最终提交
		date, bol = PretrailSubmitBack(temporder.TaskID, Userid, Ysyid, mo.Id)

		if !bol || date.Status != 100 {
			logs.Error("退回修改最终提交失败", temporder.TaskID, Userid)
			panic("退回修改最终提交失败")
		}
		logs.Debug("退回修改提交成功", temporder.TaskID, Userid)

		mo.Satus = 1
		mo.Satusmsg = "成功"
		mo.Update()
		mo.Id = 0
	}
	fmt.Println("执行结束。。。。。。。。。")
	if m.Data.UserReceiptOrderStatus == 0 {
	thisstop:
		w, er := Working(Userid, 0, Ysyid, 0)
		if er && w.Status == 100 {
			logs.Debug("停止接单")
		} else {
			time.Sleep(time.Second * 2)
			goto thisstop
		}
	}
}

//关闭订单
func ClossOrder(Userid, Usercount, timelen int, Ysyid int64) {
	endtime := time.Now().Add(time.Minute * time.Duration(timelen))

	var ddid int64
	ddvin:=""

	defer func(ddid int64) {
		var mo models.Ysyinfodetail
		mo.Id=ddid
		if err := recover(); err != nil {
			mo.Satus = 0
			mo.Satusmsg = err.(string)
			mo.Update()
			common.Requestdderror(strconv.Itoa(Userid)+":"+err.(string))
			return
		}
	}(ddid)

	m, b := PretrialPush(Userid, Ysyid, 0)
	if b && m.Status == 100 {
		if m.Data.UserReceiptOrderStatus == 0 {
			logs.Debug("接单中...")
		} else {
			w, er := Working(Userid, 1, Ysyid, 0)
			if er && w.Status == 100 {
				logs.Debug("开始接单成功")
			} else {
				logs.Error("开始接单失败")
				fmt.Println("结束...")
				panic("开始接单失败")
			}
		}
	} else {
		logs.Error("开始接单失败", m.Msg)
	}
	for i := 0; i < Usercount; i++ {
		//1，判断是否是接单状态，不是则改成接单,
	thisstart:
		m, b = PretrialPush(Userid, Ysyid, 0)
		var temporder models.PretrialOrderInfo
		//2,优先级 退回，预审中，新订单
		if len(m.Data.Returnorder) > 0 {
			temporder = m.Data.Returnorder[0]
		} else if len(m.Data.Selforder) > 0 {
			temporder = m.Data.Selforder[0]
		} else if len(m.Data.Neworder) > 0 {
			ReciveOrderConfirm(m.Data.Neworder[0].TaskID, Userid, Ysyid, 0)
			goto thisstart
		} else { //无订单是停止3秒继续
			if endtime.Before(time.Now()) {
				logs.Debug("时间已到结束。。。")
				break
			}

			time.Sleep(time.Second * 3)
			goto thisstart
		}
		if temporder.Status == 4 { //挂起
			UnlockForkedOrder(temporder.TaskID, Userid, Ysyid, 0)
			goto thisstart
		}
		var mo models.Ysyinfodetail
		mo.Ysyid = Ysyid
		mo.Userid = Userid
		mo.Satus = 0
		mo.Satusmsg = ""
		mo.Vin = temporder.Vin
		if temporder.Vin!=ddvin{
			mo.Add()
			if mo.Id == 0 {
				return
			}
			ddid=mo.Id
			ddvin=temporder.Vin
		}else {
			mo.Id=ddid
			temporder.Vin=ddvin
		}
		infoDate := TestInfoDate(temporder.Vin, temporder.TaskID, Userid, Ysyid, mo.Id)
		if !infoDate {
			logs.Error("TestInfoDate 错误停止所有执行")
			panic("错误停止所有执行")
		}
		//3，审核图片
		imgDate, bol := GetImgList(Userid, temporder.TaskID, Ysyid, mo.Id)
		count := len(imgDate.Data.RejectReasons)
		txt := ""

		if bol && count > 0 {
			r := common.RandInt64(0, int64(count)-1)
			txt = imgDate.Data.RejectReasons[r].Reason
		}
		txt += ",自动关闭备注"
		date, b := PretrailSubmitReject(temporder.TaskID, Userid, txt, Ysyid, mo.Id)
		if !b && date.Status != 100 {
			logs.Error("订单关闭失败。。", temporder.TaskID, Userid, txt)
			panic("订单关闭失败")
		}
		mo.Satus = 1
		mo.Satusmsg = "成功"
		mo.Update()
		mo.Id = 0

		logs.Debug("订单关闭操作结束", temporder.TaskID, Userid, txt)
	}
	fmt.Println("执行结束。。。。。。。。。")
	if m.Data.UserReceiptOrderStatus == 0 {
	thisstop:
		w, er := Working(Userid, 0, Ysyid, 0)
		if er && w.Status == 100 {
			logs.Debug("停止接单")
		} else {
			time.Sleep(time.Second * 2)
			goto thisstop
		}
	}
}

//机构审批
func SourceSP(Userid, Usercount, timelen int, Ysyid int64) {
	endtime := time.Now().Add(time.Minute * time.Duration(timelen))

	var ddid int64
	ddvin:=""

	defer func(ddid int64) {
		var mo models.Ysyinfodetail
		mo.Id=ddid
		if err := recover(); err != nil {
			mo.Satus = 0
			mo.Satusmsg = err.(string)
			mo.Update()
			common.Requestdderror(strconv.Itoa(Userid)+":"+err.(string))
			return
		}
	}(ddid)
	m, b := PretrialPush(Userid, Ysyid, 0)
	if b && m.Status == 100 {
		if m.Data.UserReceiptOrderStatus == 0 {
			logs.Debug("接单中...")
		} else {
			w, er := Working(Userid, 1, Ysyid,0)
			if er && w.Status == 100 {
				logs.Debug("开始接单成功")
			} else {
				logs.Error("开始接单失败")
				fmt.Println("结束...")
				panic("开始接单失败")
			}
		}
	} else {
		logs.Error("开始接单失败", m.Msg)
		panic("开始接单失败")
	}

	for i := 0; i < Usercount; i++ {
		//1，判断是否是接单状态，不是则改成接单,
	thisstart:
		m, b = PretrialPush(Userid, Ysyid, 0)

		var temporder models.PretrialOrderInfo
		//2,优先级 退回，预审中，新订单
		if len(m.Data.Returnorder) > 0 {
			temporder = m.Data.Returnorder[0]
		} else if len(m.Data.Selforder) > 0 {
			temporder = m.Data.Selforder[0]
		} else if len(m.Data.Neworder) > 0 {
			ReciveOrderConfirm(m.Data.Neworder[0].TaskID, Userid, Ysyid, 0)
			goto thisstart
		} else { //无订单是停止3秒继续
			if endtime.Before(time.Now()) {
				logs.Debug("时间已到结束。。。")
				break
			}

			time.Sleep(time.Second * 3)
			goto thisstart
		}
		if temporder.Status == 4 { //挂起
			UnlockForkedOrder(temporder.TaskID, Userid, Ysyid, 0)
			goto thisstart
		}
		var mo models.Ysyinfodetail
		mo.Ysyid = Ysyid
		mo.Userid = Userid
		mo.Satus = 0
		mo.Satusmsg = ""
		mo.Vin = temporder.Vin
		if temporder.Vin!=ddvin{
			mo.Add()
			if mo.Id == 0 {
				return
			}
			ddid=mo.Id
			ddvin=temporder.Vin
		}else {
			mo.Id=ddid
			temporder.Vin=ddvin
		}
		infoDate := TestInfoDate(temporder.Vin, temporder.TaskID, Userid, Ysyid, mo.Id)
		if !infoDate {
			logs.Error("TestInfoDate 错误停止所有执行")
			panic("错误停止所有执行")
		}
		//3，审核图片
		imgDate, bol := GetImgList(Userid, temporder.TaskID, Ysyid, mo.Id)
		count := len(imgDate.Data.RejectReasons)
		txt := ""

		if bol && count > 0 {
			r := common.RandInt64(0, int64(count)-1)
			txt = imgDate.Data.RejectReasons[r].Reason
		} else {
			logs.Error("订单审批操错误。。", temporder.TaskID, Userid, txt)
			panic("订单审批操错误")
		}
		txt += ",自动审批备注"
		date, b := PretrailSubmitBackOrg(temporder.TaskID, Userid, txt, Ysyid, mo.Id)
		if !b && date.Status != 100 {
			logs.Error("订单审批失败。。", temporder.TaskID, Userid, txt)
			panic("订单审批失败")
		}

		mo.Satus = 1
		mo.Satusmsg = "成功"
		mo.Update()
		mo.Id = 0

		logs.Debug("订单审批操作结束", temporder.TaskID, Userid, txt)
	}
	fmt.Println("执行结束。。。。。。。。。")
	if m.Data.UserReceiptOrderStatus == 0 {
	thisstop:
		w, er := Working(Userid, 0, Ysyid, 0)
		if er && w.Status == 100 {
			logs.Debug("停止接单")
		} else {
			time.Sleep(time.Second * 2)
			goto thisstop
		}
	}
}


////////////////////////////////接口逻辑//////////////////////////////////

//获取图片列表 ok
func GetImgList(userid, taskid int, Ysyid, Ysydid int64) (models.ResultDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx?"
	m := make(map[string]string, 0)
	m["op"] = "GetImgList"
	m["taskId"] = strconv.Itoa(taskid)
	m["userId"] = strconv.Itoa(userid)

	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.ResultDate
	if b {
		fmt.Println(string(res))
		 json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"GetImgList",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}


}

//开始接单/停止接单 ok
func Working(userid, Working int, Ysyid, Ysydid int64) (models.ResultDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"
	m := make(map[string]string, 0)
	m["op"] = "Working"
	m["Working"] = strconv.Itoa(Working) //1接单，0停止接单
	m["userId"] = strconv.Itoa(userid)

	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.ResultDate
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			logs.Error(string(res))
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"Working",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		logs.Error(string(res))
		return Data, false
	}
}

//拉取派单数据 ok
func PretrialPush(userid int, Ysyid, Ysydid int64) (models.PretrialPush, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialPush.ashx"
	m := make(map[string]string, 0)
	m["userId"] = strconv.Itoa(userid)
	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.PretrialPush
	if b {
		 json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"PretrialPush",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//获取历史订单 ok
func GetHistoryReports(vin string, Ysyid, Ysydid int64) (models.OrderHistoryDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"
	m := make(map[string]string, 0)
	m["vin"] = vin
	m["op"] = "GetHistoryReports"
	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.OrderHistoryDate
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			fmt.Print(string(res))
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"GetHistoryReports",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//获取操作历史 ok
func GetOperateRecord(taskid int, Ysyid, Ysydid int64) (models.OperateLogDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"
	m := make(map[string]string, 0)
	m["op"] = "GetOperateRecord"
	m["taskId"] = strconv.Itoa(taskid)

	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.OperateLogDate
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			fmt.Print(string(res))
			return Data, false
		}

	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"GetOperateRecord",Timelength)
	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//获取图片详情 ok
func GetImgDetail(taskid, itemId, userId int, Ysyid, Ysydid int64) (models.GetImgDetailReplyDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"
	m := make(map[string]string, 0)
	m["op"] = "GetImgDetail"
	m["taskId"] = strconv.Itoa(taskid)
	m["itemId"] = strconv.Itoa(itemId)
	m["userId"] = strconv.Itoa(userId)

	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.GetImgDetailReplyDate
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}

	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"GetImgDetail",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//获取订单基本信息 ok
func GetOrderInfo(taskid, userId int, Ysyid, Ysydid int64) (models.OrderInfoModelDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"
	m := make(map[string]string, 0)
	m["op"] = "GetOrderInfo"
	m["taskId"] = strconv.Itoa(taskid)
	m["userId"] = strconv.Itoa(userId)

	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.OrderInfoModelDate
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"GetOrderInfo",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//获取省市列表 ok
func GetProvincesAndCitys(taskid, userId int, Ysyid, Ysydid int64) (models.ProvincesAndCitysVoDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"
	m := make(map[string]string, 0)
	m["op"] = "GetProvincesAndCitys"

	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.ProvincesAndCitysVoDate
	if b {
		 json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"GetProvincesAndCitys",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//根据车牌定位城市 ok
func GetCityAndProvinceByPlatName(plateName string, Ysyid, Ysydid int64) (models.ProvinceCityModelDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"
	m := make(map[string]string, 0)
	m["op"] = "GetCityAndProvinceByPlatName"
	m["plateName"] = plateName
	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.ProvinceCityModelDate
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"GetCityAndProvinceByPlatName",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//图片和视频审核通过 ok
func ImgCheckPass(taskId, userId, itemId, video int, Ysyid, Ysydid int64) (models.ResultPublicDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"

	m := make(map[string]string, 0)
	m["op"] = "ImgCheckPass"
	m["taskId"] = strconv.Itoa(taskId)
	m["userId"] = strconv.Itoa(userId)
	m["itemId"] = strconv.Itoa(itemId) //-1为视频
	if video == 1 { //视频
		m["itemId"] = "-1" //-1为视频
	}

	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.ResultPublicDate
	if b {
		 json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"ImgCheckPass",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//基本信息保存 ok
func SaveFormData(model models.TaskCarBasicEPModel, userId int, Ysyid, Ysydid int64) (models.ResultPublicDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"
	var Data models.ResultPublicDate
	m := make(map[string]string, 0)
	m["op"] = "SaveFormData"
	m["taskId"] = strconv.Itoa(model.Id)
	m["userId"] = strconv.Itoa(userId)
	model = GetSaveFormData(model)

	if model.TaskType == 1 || model.TaskType == 5 {
		c := common.RandInt64(0, int64(len(CityA))-1)
		date, b := GetCityAndProvinceByPlatName(CityA[c], Ysyid, Ysydid)
		if !b || date.Status != 100 {
			logs.Error("GetCityAndProvinceByPlatName 错误", CityA[c])
			return Data, false
		}
		model.CarLicense = CityA[c]
	this:
		if len(model.CarLicense) < 7 {
			model.CarLicense += strconv.FormatInt(common.RandInt64(0, 9), 10)
			goto this
		}

		model.ProName = date.Data.ProvinceName
		model.RegisterProvID = date.Data.ProvinceID
		model.CityName = date.Data.CityName
		model.RegisterCityID = date.Data.CityId
	}

	bytes, _ := json.Marshal(model)
	m["data"] = string(bytes)

	res, b,Timelength := httpdate.SendPostys(url, m, "")

	if b {
		 json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"SaveFormData",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}

}
func GetSaveFormData(m models.TaskCarBasicEPModel) models.TaskCarBasicEPModel {

	if m.TaskType == 1 { //18张
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
	} else if m.TaskType == 2 { ///6张
		m.Mileage = 60000
	} else if m.TaskType == 5 { ///9张
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

//预审通过提交 ok
func PretrailSubmitPass(taskId, userId int, Ysyid, Ysydid int64) (models.ResultPublicDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"
	m := make(map[string]string, 0)
	m["op"] = "PretrailSubmitPass"
	m["taskId"] = strconv.Itoa(taskId)
	m["userId"] = strconv.Itoa(userId)

	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.ResultPublicDate
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"PretrailSubmitPass",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//解除挂起 ok
func UnlockForkedOrder(taskId, userId int, Ysyid, Ysydid int64) (models.ResultPublicDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"
	m := make(map[string]string, 0)
	m["op"] = "UnlockForkedOrder"
	m["taskId"] = strconv.Itoa(taskId)
	m["userId"] = strconv.Itoa(userId)
	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.ResultPublicDate
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"UnlockForkedOrder",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//认领验证 ok
func ReciveOrderCheck(taskId, userId int, Ysyid, Ysydid int64) (models.ResultPublicDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"
	m := make(map[string]string, 0)
	m["op"] = "ReciveOrderCheck"
	m["taskId"] = strconv.Itoa(taskId)
	m["userId"] = strconv.Itoa(userId)
	m["types"] = "1"

	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.ResultPublicDate
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"ReciveOrderCheck",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//订单认领 ok
func ReciveOrderConfirm(taskId, userId int, Ysyid, Ysydid int64) (models.ResultPublicDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"

	date, er := ReciveOrderCheck(taskId, userId, Ysyid, Ysydid)
	if !er || date.Status != 100 {
		return date, false
	}
	m := make(map[string]string, 0)
	m["op"] = "ReciveOrderConfirm"
	m["taskId"] = strconv.Itoa(taskId)
	m["userId"] = strconv.Itoa(userId)
	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.ResultPublicDate
	if b {
		 json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"ReciveOrderConfirm",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//替换图片，上传附加 ok
func UploadPic(taskId, picId, itemId int, Ysyid, Ysydid int64) (models.ResultPublicDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"

	m := make(map[string]string, 0)
	m["op"] = "UploadPic"
	m["TaskID"] = strconv.Itoa(taskId)
	m["picId"] = strconv.Itoa(picId)
	m["itemId"] = strconv.Itoa(itemId)
	filename := beego.AppConfig.String("pic.picth")
	res, b,Timelength := httpdate.SendPostys(url, m, filename)
	var Data models.ResultPublicDate
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"UploadPic",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//审核不通过加照片 ok
func Upload_SampleImg(taskId, userId, itemId, returnId int, Ysyid, Ysydid int64) (models.UploadPic, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"

	m := make(map[string]string, 0)
	m["op"] = "Upload_SampleImg"
	m["taskId"] = strconv.Itoa(taskId)
	m["userId"] = strconv.Itoa(userId)
	m["itemId"] = strconv.Itoa(itemId)
	m["returnId"] = strconv.Itoa(returnId)

	filename := beego.AppConfig.String("pic.picthpic")
	res, b,Timelength := httpdate.SendPostys(url, m, filename)
	var Data models.UploadPic
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"Upload_SampleImg",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//审核不通过 ok
func ImgCheckUnPass(taskId, userId, itemId, video int, pic models.GetImgDetailReplyDate, Ysyid, Ysydid int64) (models.ResultPublicDate, bool, int) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"

	//unPassReasonJson: [{"Type":1,"ItemID":270,"ReturnID":253,"ReturnReason":"拍摄位置错误，请重拍","ReturnName":"后排座椅","SampleImg":"http://192.168.0.156/group2/M01/0E/91/wKgAl14GI8qAAYUOAAK-Sixj33g537.jpg","FastDFSBasePath":"","TitleText":"eeeeeeeeeeeee","FileName":"2.jpg","DefaultAttachUrl":"https://image.jingzhengu.com/JCXTImgUrl/SampleImage/eighteen6.png"}]
	var Data models.ResultPublicDate
	m := make(map[string]string, 0)
	m["op"] = "ImgCheckUnPass"
	m["taskId"] = strconv.Itoa(taskId)
	m["userId"] = strconv.Itoa(userId)
	m["itemId"] = strconv.Itoa(itemId) //-1为视频

	lens := len(pic.Data.TxtReturnList)
	if lens == 0 {
		return Data, false, 0
	}

	r := common.RandInt64(1, int64(lens))
	mo := pic.Data.TxtReturnList[r]

	relist := make([]models.TaskReturnLogModel, 0)
	mo.TitleText = "预审自动测试提交"
	//可以上传图片
	uploadPic, bol := Upload_SampleImg(taskId, userId, itemId, mo.TaskReturnLogModel.ReturnID,Ysyid, Ysydid )
	if !bol || uploadPic.Status != 100 {
		logs.Error("驳回上传图片报错")
		return Data, false, 0
	}
	mo.TaskReturnLogModel.SampleImg = uploadPic.PicPath
	mo.TaskReturnLogModel.FileName = uploadPic.PicName
	mo.TaskReturnLogModel.FastDFSBasePath = uploadPic.FastDFSBasePath

	relist = append(relist, mo.TaskReturnLogModel)

	t, _ := json.Marshal(relist)
	m["unPassReasonJson"] = string(t)

	if video == 1 { //视频
		m["itemId"] = "-1" //-1为视频
	}

	res, b,Timelength := httpdate.SendPostys(url, m, "")

	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false, 0
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"ImgCheckUnPass",Timelength)

	if Data.Status == 100 {
		return Data, true, mo.TaskReturnLogModel.ReturnID
	} else {
		return Data, false, 0
	}
}

//驳回新增附加图片 ok
func YsyReturnSummaryReason_Save(taskId, userId int, Ysyid, Ysydid int64) (models.ResultPublicDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"

	m := make(map[string]string, 0)
	m["op"] = "YsyReturnSummaryReason_Save"
	m["taskId"] = strconv.Itoa(taskId)
	m["userId"] = strconv.Itoa(userId)

	list := make([]models.AddFJ, 0)
	var mo models.AddFJ
	mo.ReturnReason = "新增附加1"
	mo.TitleText = "新增附加1说明"
	mo.SampleImg = beego.AppConfig.String("pgs.fjurl")
	mo.FileName = "yscstjfu.jpg"
	mo.PicID = 0
	list = append(list, mo)
	jso, _ := json.Marshal(list)
	m["data"] = string(jso)

	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.ResultPublicDate
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"YsyReturnSummaryReason_Save",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}

}

//退回修改 ok
func Ys_ImgCheckUnPassSummaryRemark(taskId, userId int, Ysyid, Ysydid int64) (models.ResultPublicDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"

	m := make(map[string]string, 0)
	m["op"] = "Ys_ImgCheckUnPassSummaryRemark"
	m["taskId"] = strconv.Itoa(taskId)
	m["userId"] = strconv.Itoa(userId)
	m["remark"] = "订单退回修改备注123"

	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.ResultPublicDate
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"Ys_ImgCheckUnPassSummaryRemark",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//退回修改 ok
func PretrailSubmitBack(taskId, userId int, Ysyid, Ysydid int64) (models.ResultPublicDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"

	m := make(map[string]string, 0)
	m["op"] = "PretrailSubmitBack"
	m["taskId"] = strconv.Itoa(taskId)
	m["userId"] = strconv.Itoa(userId)
	m["txt"] = "订单退回修改备注123"

	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.ResultPublicDate
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"PretrailSubmitBack",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//订单关闭 ok
func PretrailSubmitReject(taskId, userId int, txt string, Ysyid, Ysydid int64) (models.ResultPublicDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"

	m := make(map[string]string, 0)
	m["op"] = "PretrailSubmitReject"
	m["taskId"] = strconv.Itoa(taskId)
	m["userId"] = strconv.Itoa(userId)
	m["txt"] = txt

	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.ResultPublicDate
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"PretrailSubmitReject",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//机构审批提交 ok
func PretrailSubmitBackOrg(taskId, userId int, txt string, Ysyid, Ysydid int64) (models.ResultPublicDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"
	m := make(map[string]string, 0)
	m["op"] = "PretrailSubmitBackOrg"
	m["taskId"] = strconv.Itoa(taskId)
	m["userId"] = strconv.Itoa(userId)
	m["txt"] = txt
	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.ResultPublicDate
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"PretrailSubmitBackOrg",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//订单其他数据测试
func TestInfoDate(vin string, taskid, userId int, Ysyid, Ysydid int64) bool {
	//历史报告
	date, _ := GetHistoryReports(vin, Ysyid, Ysydid)
	if date.Status != 100 {
		panic("TestInfoDate——GetHistoryReports 错误")
	}
	//操作历史
	logDate, _ := GetOperateRecord(taskid, Ysyid, Ysydid)
	if logDate.Status != 100 {
		panic("TestInfoDate——GetOperateRecord 错误")
	}
	//获取订单基本信息
	modelDate, _ := GetOrderInfo(taskid, userId, Ysyid, Ysydid)
	if modelDate.Status != 100 {
		panic("TestInfoDate——GetOrderInfo 错误")
	}
	//获取省市列表
	voDate, _ := GetProvincesAndCitys(taskid, userId, Ysyid, Ysydid)
	if voDate.Status != 100 {
		panic("TestInfoDate——GetOrderInfo 错误")
	}
	groupDate, _ := CheckPassDescSearch(taskid, "订单", Ysyid, Ysydid)
	if groupDate.Status != 100 {
		logs.Error("TestInfoDate——CheckPassDescSearch 错误", taskid, userId)
		panic("TestInfoDate——CheckPassDescSearch 错误")
	}
	return true
}

//机构标准搜索 ok
func CheckPassDescSearch(taskid int, keyWord string, Ysyid, Ysydid int64) (models.CheckPassDescGroupDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"
	m := make(map[string]string, 0)
	m["op"] = "CheckPassDescSearch"
	m["TaskID"] = strconv.Itoa(taskid)
	m["keyWord"] = keyWord

	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.CheckPassDescGroupDate
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"CheckPassDescSearch",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//添加备注 ok
func Ys_AddRemark(taskid, userId int, remark string, Ysyid, Ysydid int64) (models.ResultPublicDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"
	m := make(map[string]string, 0)
	m["op"] = "Ys_AddRemark"
	m["TaskID"] = strconv.Itoa(taskid)
	m["userId"] = strconv.Itoa(userId)
	m["remark"] = remark

	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.ResultPublicDate
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"Ys_AddRemark",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//删除图片,预审附件 ok
func DeletePic_YsAttach(taskid, picId int, Ysyid, Ysydid int64) (models.ResultPublicDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"
	m := make(map[string]string, 0)
	m["op"] = "DeletePic_YsAttach"
	m["TaskID"] = strconv.Itoa(taskid)
	m["picId"] = strconv.Itoa(picId)

	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.ResultPublicDate
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"DeletePic_YsAttach",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//驳回图片删除 ok
func ImgCheckUnPass_AttachDelete(taskid, userId, itemId, returnId int, Ysyid, Ysydid int64) (models.ResultPublicDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"
	m := make(map[string]string, 0)
	m["op"] = "ImgCheckUnPass_AttachDelete"
	m["TaskID"] = strconv.Itoa(taskid)
	m["userId"] = strconv.Itoa(userId)
	m["itemId"] = strconv.Itoa(itemId)
	m["returnId"] = strconv.Itoa(returnId)

	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.ResultPublicDate
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"ImgCheckUnPass_AttachDelete",Timelength)

	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//挂起订单 ok
func ForkOrder(taskid, userId int, Ysyid, Ysydid int64) (models.ResultPublicDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"
	m := make(map[string]string, 0)
	m["op"] = "ForkOrder"
	m["TaskID"] = strconv.Itoa(taskid)
	m["userId"] = strconv.Itoa(userId)

	res, b,Timelength := httpdate.SendPostys(url, m, "")
	var Data models.ResultPublicDate
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"ForkOrder",Timelength)
	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}
}

//附件图片操作记录显示 ok
func UploadPic_YsAttach(taskid, itemId int, Ysyid, Ysydid int64) (models.FJpic, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"
	m := make(map[string]string, 0)
	m["op"] = "UploadPic_YsAttach"
	m["TaskID"] = strconv.Itoa(taskid)
	m["itemId"] = strconv.Itoa(itemId)
	pic := beego.AppConfig.String("pic.picczjl")

	res, b,Timelength := httpdate.SendPostys(url, m, pic)

	var Data models.FJpic
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"UploadPic_YsAttach",Timelength)
	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}

}
//删除附加
func DeletePic(taskId,picId int, Ysyid, Ysydid int64) (models.ResultPublicDate, bool) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"
	m := make(map[string]string, 0)
	m["op"] = "DeletePic"
	m["picId"] = strconv.Itoa(picId)
	m["taskId"] = strconv.Itoa(taskId)

	res, b,Timelength := httpdate.SendPostys(url, m, "")

	var Data models.ResultPublicDate
	if b {
		json.Unmarshal(res, &Data)
		if Data.Status != 100 {
			return Data, false
		}
	}
	go insetinterfice(Ysyid, Ysydid,Data.Status,Data.Msg,"UploadPic_YsAttach",Timelength)
	if Data.Status == 100 {
		return Data, true
	} else {
		return Data, false
	}

}

func insetinterfice(Ysyid, Ysydid int64, Status int, Txt, Iname string, Timelen float64) {

	if Ysyid==0  {
		return
	}

	var m models.Ysyinfodetailinterfice
	m.Ysyid = Ysyid
	m.Ysydid = Ysydid
	m.Status = Status
	m.Txt = Txt
	m.Iname = Iname
	m.Timelen = Timelen
	m.Add()
}

