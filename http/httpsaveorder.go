package httpdate

import (
	"../common"
	"../models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
)

//6张下单
func SendPostFormFile6(userid int, configID int, procductlist int, vin string, id int64) {

	url := beego.AppConfig.String("app.url") + "/app/TaskSaveSimple.ashx"

	filename := beego.AppConfig.String("zip.pic6")
	token := beego.AppConfig.String("app.userTokenet")

	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)
	// 1. 要上传的数据
	resmap := GetFastalue6(userid, configID, procductlist, vin)
	sigin := common.GetSign(resmap, token)
	fmt.Println("sigin", sigin)
	for k, v := range resmap {
		body_writer.WriteField(k, v)
	}
	body_writer.WriteField("sign", sigin)

	// 3. 读取文件
	_, err := body_writer.CreateFormFile("application", filename)
	if err != nil {
		logs.Error("创建FormFile2文件信息异常！", err)
		return
	}
	fb2, err := ioutil.ReadFile(filename)
	if err != nil {
		logs.Error("打开文件异常!", err)
		return
	}
	body_buf.Write(fb2)

	// 结束整个消息body
	body_writer.Close()

	//
	req_reader := io.MultiReader(body_buf)
	req, err := http.NewRequest("POST", url, req_reader)
	if err != nil {
		logs.Error("站点相机上传图片，创建上次请求异常！异常信息:", err)
		return
	}
	// 添加Post头
	req.Header.Set("Connection", "close")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Content-Type", body_writer.FormDataContentType())
	req.ContentLength = int64(body_buf.Len())

	// 发送消息
	client := &http.Client{}
	starttime := time.Now()
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取回应消息异常:", err)
	}
	logs.Debug("发送回应数据:"+vin, string(body))
	endtime := time.Now()
	var res models.AppResultModel
	b := json.Unmarshal(body, &res)

	if b == nil {
		Timelength := endtime.Sub(starttime) //两个时间相减
		WriteOrderInfodetail(id, res, Timelength.Seconds(), vin)
	} else {
		logs.Error("SendPostFormFile9:解析json失败")
	}
	return
}

//18,20张下单
func SendPostFormFile(userid int, configID int, procductlist int, vin string, id int64) {
	url := beego.AppConfig.String("app.url") + "/app/TaskSave20160303.ashx"

	filename := beego.AppConfig.String("zip.pic20")
	NewEdition := 1
	if configID != 1 || (procductlist == 11 && procductlist == 13 && procductlist == 14) {
		NewEdition = 0
	}
	token := beego.AppConfig.String("app.userTokenet")

	if configID == 4 {
		filename = beego.AppConfig.String("zip.pic18")
	} else if configID == 6 {
		filename = beego.AppConfig.String("zip.pic13")
	} else if configID == 7 {
		filename = beego.AppConfig.String("zip.pic16")
	} else if configID == 8 {
		filename = beego.AppConfig.String("zip.pic26")
	}

	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)
	// boundary默认会提供一组随机数，也可以自己设置。

	// 1. 要上传的数据
	resmap := GetFastValue18(userid, configID, procductlist, vin, NewEdition)
	sigin := common.GetSign(resmap, token)

	for k, v := range resmap {
		body_writer.WriteField(k, v)
	}
	body_writer.WriteField("sign", sigin)
	// 3. 读取文件
	_, err := body_writer.CreateFormFile("application", filename)
	if err != nil {
		logs.Error("创建FormFile2文件信息异常！", err)
		return
	}
	fb2, err := ioutil.ReadFile(filename)
	if err != nil {
		logs.Error("打开文件异常!", err)
		return
	}
	body_buf.Write(fb2)

	// 结束整个消息body
	body_writer.Close()

	//
	req_reader := io.MultiReader(body_buf)
	req, err := http.NewRequest("POST", url, req_reader)
	if err != nil {
		logs.Error("error:", err)
		return
	}
	// 添加Post头
	req.Header.Set("Connection", "close")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Content-Type", body_writer.FormDataContentType())
	req.ContentLength = int64(body_buf.Len())

	// 发送消息
	client := &http.Client{}
	starttime := time.Now()
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("读取回应消息异常SendPostFormFile:", err)
	}
	logs.Debug("接收SendPostFormFile:"+vin, string(body))
	endtime := time.Now()
	var res models.AppResultModel
	b := json.Unmarshal(body, &res)

	if b == nil {
		Timelength := endtime.Sub(starttime) //两个时间相减
		go WriteOrderInfodetail(id, res, Timelength.Seconds(), vin)
	} else {
		logs.Error("SendPostFormFile:解析json失败", vin, string(body))
	}
	return
}

//9张下单
func SendPostFormFile9(userid int, configID int, procductlist int, vin string, id int64) {

	url := beego.AppConfig.String("app.url") + "/app/TaskSave9Pic.ashx"
	filename := beego.AppConfig.String("zip.pic9")
	token := beego.AppConfig.String("app.userTokenet")
	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)

	// 1. 要上传的数据
	resmap := GetFastValue9(userid, configID, procductlist, vin)
	sigin := common.GetSign(resmap, token)
	for k, v := range resmap {
		body_writer.WriteField(k, v)
	}
	body_writer.WriteField("sign", sigin)

	// 3. 读取文件
	_, err := body_writer.CreateFormFile("application", filename)
	if err != nil {
		logs.Error("创建FormFile2文件信息异常9！", err)
		return
	}
	fb2, err := ioutil.ReadFile(filename)
	if err != nil {
		logs.Error("打开文件异常!SendPostFormFile9", err)
		return
	}
	body_buf.Write(fb2)

	// 结束整个消息body
	body_writer.Close()

	req_reader := io.MultiReader(body_buf)
	req, err := http.NewRequest("POST", url, req_reader)
	if err != nil {
		logs.Error("SendPostFormFile9 error:", err)
		return
	}
	// 添加Post头
	req.Header.Set("Connection", "close")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Content-Type", body_writer.FormDataContentType())
	req.ContentLength = int64(body_buf.Len())

	// 发送消息
	client := &http.Client{}
	starttime := time.Now()
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("读取回应消息异常SendPostFormFile9:", err)
	}
	logs.Debug("接收返回数据SendPostFormFile9:"+vin, string(body))
	endtime := time.Now()
	var res models.AppResultModel
	b := json.Unmarshal(body, &res)
	fmt.Println("调用接口时间", starttime, endtime)
	if b == nil {
		Timelength := endtime.Sub(starttime) //两个时间相减
		go WriteOrderInfodetail(id, res, Timelength.Seconds(), vin)
	} else {
		logs.Error("SendPostFormFile9:解析json失败", vin, string(body))
	}

	return
}

//快估下单
func Fast(userid int, procductlist int, vin string, isPretrial int, id int64) {

	url := beego.AppConfig.String("app.Fasturl") + "/api/onLineTask7Pic/addEighteenthTask"

	token := beego.AppConfig.String("app.userToken")
	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)
	// boundary默认会提供一组随机数，也可以自己设置。
	//body_writer.SetBoundary("Pp7Ye2EeWaFDdAY")

	// 1. 要上传的数据
	resmap := GetFastValue(userid, procductlist, vin, isPretrial)
	sigin := common.GetSign(resmap, token)
	for k, v := range resmap {
		body_writer.WriteField(k, v)
	}
	body_writer.WriteField("sign", sigin)
	// 结束整个消息body
	body_writer.Close()

	//
	req_reader := io.MultiReader(body_buf)
	req, err := http.NewRequest("POST", url, req_reader)
	if err != nil {
		logs.Error("创建请求异常 Fast ", err)
		return
	}
	// 添加Post头
	req.Header.Set("Connection", "close")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Content-Type", body_writer.FormDataContentType())
	req.ContentLength = int64(body_buf.Len())

	// 发送消息
	client := &http.Client{}
	starttime := time.Now()
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("读取回应消息异常 Fast ", err)
	}
	logs.Debug("收到返回数据 Fast ", vin, string(body))
	endtime := time.Now()
	var res models.AppResultModel
	b := json.Unmarshal(body, &res)
	logs.Debug("调用接口时间 Fast", starttime, endtime)
	if b == nil {
		Timelength := endtime.Sub(starttime) //两个时间相减
		go WriteOrderInfodetail(id, res, Timelength.Seconds(), vin)
	} else {
		logs.Error("Fast:解析json失败", vin, string(body))
	}
	return
}

//快估获取参数
func GetFastValue(userid int, procductlist int, vin string, isPretrial int) map[string]string {

	res := make(map[string]string, 0)
	// 1. 要上传的数据
	res["productType"] = strconv.Itoa(procductlist)
	res["vin"] = vin
	res["Longitude"] = "116.312877"
	res["equipmentNo"] = "868198FA7F3B804B8557021B509C8CD0"
	res["modelID"] = "4765"
	res["userId"] = strconv.Itoa(userid)
	res["provinceId"] = "5"
	res["transferCount"] = "0"
	res["service"] = "2"
	res["engineNum"] = "SDFASDF"
	res["tokenid"] = "6"
	res["seating"] = "2"
	res["Address"] = "中国北京市海淀区苏州街16号"
	res["appVersion"] = "3.3.9"
	res["registerCityId"] = "201"
	res["color"] = "黑色"
	res["Latitude"] = "39.985431"
	res["recordDate"] = "2018-06-07"
	res["cityId"] = "505"
	res["orderPhone"] = "15313636363"
	res["vinAnalyzing"] = "2"
	res["Telephone"] = ""
	res["productionTime"] = "2018-06-05"
	res["taskOwnerName"] = "LJS"
	res["makeID"] = "97"
	res["styleID"] = "119120"
	res["userMd5"] = ""
	res["orderName"] = "ljs"
	res["Des"] = "LJS"
	res["registerProvId"] = "2"
	res["recordBrand"] = "SADFASDF"
	res["mileage"] = "10"
	res["carLicense"] = "京A12347"
	res["deviceInfo"] = ""
	res["platType"] = "1"
	res["appendImageList"] = ""
	res["orderTelPhone"] = ""
	res["imageList"] = ""                        //图片集合
	res["videoPath"] = ""                        //视频地址
	res["isPretrial"] = strconv.Itoa(isPretrial) //0全量照片 1无照片

	list := make([]models.ImageList, 0)

	if procductlist == 13 { //2张
		var m1 models.ImageList
		m1.PartCode = 242
		m1.ImageUrl = "group4/M01/91/9C/wKgUk16fo1SASPZhAABPqjB__is104.jpg"
		list = append(list, m1)
		var m2 models.ImageList
		m2.PartCode = 253
		m2.ImageUrl = "group4/M01/91/9C/wKgUk16fo1SASPZhAABPqjB__is104.jpg"
		list = append(list, m2)
		j, e := json.Marshal(list)
		if e == nil {
			res["imageList"] = string(j)
		}

		res["isPretrial"] = "0" //0全量照片 1非全量 无照片
	} else { //7张
		if isPretrial == 0 {
			var m1 models.ImageList
			m1.PartCode = 242
			m1.ImageUrl = "group4/M01/91/9C/wKgUk16fo1SASPZhAABPqjB__is104.jpg"
			list = append(list, m1)
			var m2 models.ImageList
			m2.PartCode = 247
			m2.ImageUrl = "group4/M01/91/9C/wKgUk16fo1SASPZhAABPqjB__is104.jpg"
			list = append(list, m2)
			var m3 models.ImageList
			m3.PartCode = 291
			m3.ImageUrl = "group4/M01/91/9C/wKgUk16fo1SASPZhAABPqjB__is104.jpg"
			list = append(list, m3)
			var m4 models.ImageList
			m4.PartCode = 292
			m4.ImageUrl = "group4/M01/91/9C/wKgUk16fo1SASPZhAABPqjB__is104.jpg"
			list = append(list, m4)
			var m5 models.ImageList
			m5.PartCode = 296
			m5.ImageUrl = "group4/M01/91/9C/wKgUk16fo1SASPZhAABPqjB__is104.jpg"
			list = append(list, m5)
			var m6 models.ImageList
			m6.PartCode = 298
			m6.ImageUrl = "group4/M01/91/9C/wKgUk16fo1SASPZhAABPqjB__is104.jpg"
			list = append(list, m6)

			var m7 models.ImageList
			m7.PartCode = 295
			m7.ImageUrl = "group4/M01/91/9C/wKgUk16fo1SASPZhAABPqjB__is104.jpg"
			list = append(list, m7)

			var m8 models.ImageList
			m8.PartCode = 293
			m8.ImageUrl = "group4/M01/91/9C/wKgUk16fo1SASPZhAABPqjB__is104.jpg"
			list = append(list, m8)
			var m9 models.ImageList
			m9.PartCode = 309
			m9.ImageUrl = "group4/M01/91/9C/wKgUk16fo1SASPZhAABPqjB__is104.jpg"
			list = append(list, m9)

			var m10 models.ImageList
			m10.PartCode = 297
			m10.ImageUrl = "group4/M01/91/9C/wKgUk16fo1SASPZhAABPqjB__is104.jpg"
			list = append(list, m10)

			var m11 models.ImageList
			m11.PartCode = 308
			m11.ImageUrl = "group4/M01/91/9C/wKgUk16fo1SASPZhAABPqjB__is104.jpg"
			list = append(list, m11)

			j, e := json.Marshal(list)
			if e == nil {
				res["imageList"] = string(j)
			}
			res["videoPath"] = "https://imgv5.jingzhengu.com/group3/M03/83/25/wKgUkV3wXFyAJvFzAExV9qPD074781.mp4" //视频地址
			res["isPretrial"] = "0"                                                                               //0全量照片 1无照片
		} else {
			res["videoPath"] = ""   //视频地址
			res["imageList"] = ""   //图片集合
			res["isPretrial"] = "1" //0全量照片 1无照片
		}

	}

	return res
}

//6张获取参数
func GetFastalue6(userid int, configID int, procductlist int, vin string) map[string]string {

	res := make(map[string]string, 0)

	res["Longitude"] = "116.312861"
	res["userId"] = strconv.Itoa(userid)
	res["productType"] = strconv.Itoa(procductlist)
	res["VinCode"] = vin
	res["ConfigID"] = strconv.Itoa(configID)
	res["equipmentNo"] = "868198FA7F3B804B8557021B509C8CD0"
	res["MakeId"] = "136"
	res["ModelId"] = "4303"
	res["RegDate"] = "2017-12-01"
	res["Perf_DriveType"] = "前轮驱动"
	res["op"] = "save"
	res["tokenid"] = "6"
	res["OrderCityId"] = "901"
	res["ProgramId"] = "2_1"
	res["Address"] = "中国北京市海淀区苏州街18号a2-512室"
	res["Engine_Exhaust"] = "1.6"
	res["appVersion"] = "3.3.9"
	res["telephone"] = ""
	res["Latitude"] = "39.985435"
	res["OrderProvinceId"] = "9"
	res["TransmissionType"] = "手自一体"
	res["ConfigID"] = "2"
	res["Des"] = "LJS备注"
	res["platType"] = "1"
	res["deviceInfo"] = "{\"brand\":\"OPPO\",\"model\":\"\",\"osVersion\":\"5.1.1\",\"platform\":\"android\",\"resolution\":\"1080*1920\"}"

	return res
}

//9张获取参数
func GetFastValue9(userid int, configID int, procductlist int, vin string) map[string]string {

	res := make(map[string]string, 0)

	res["RecordBrand"] = "Xjxjjxdj"
	res["TaskOwnerName"] = "Zjzj"
	res["tokenid"] = "6"
	res["appVersion"] = "3.3.9"
	res["Address"] = "beijing"
	res["OrderCityId"] = "901"
	res["Latitude"] = "39.985395"
	res["Tasktel"] = "15313666764"
	res["carLicense"] = "冀ACJXJX"
	res["BusinessPrice"] = "D53F7B90C8758CF6DFB0DACEEA3AABCD"
	res["OrderProvinceId"] = "9"
	res["op"] = "save"
	res["productType"] = "9"
	res["RecordDate"] = "2019-12-01"
	res["Service"] = "2"
	res["telephone"] = ""
	res["equipmentNo"] = "B3B51F6D80E09F24FD6652BE50E51D80"
	res["ProgramId"] = ""
	res["Longitude"] = "116.312729"
	res["userId"] = strconv.Itoa(userid)
	res["Des"] = "订单备注"
	res["publishType"] = strconv.Itoa(procductlist)
	res["drivingLicense"] = ""
	res["VinCode"] = vin
	res["EngineNum"] = "Xjxjjxjddj"
	res["ConfigID"] = strconv.Itoa(configID)
	res["platType"] = ""
	res["deviceInfo"] = "{\"brand\":\"OPPO\",\"model\":\"\",\"osVersion\":\"5.1.1\",\"platform\":\"android\",\"resolution\":\"1080*1920\"}"

	return res
}

//18,20张获取参数
func GetFastValue18(userid int, configID int, procductlist int, vin string, NewEdition int) map[string]string {

	res := make(map[string]string, 0)

	res["RecordBrand"] = "Xjxjjxdj"
	res["TaskOwnerName"] = "Zjzj"
	res["tokenid"] = "6"
	res["appVersion"] = "399.399.999"
	res["Address"] = "beijing"
	res["OrderCityId"] = "901"
	res["Latitude"] = "39.985395"
	res["Tasktel"] = "15313666764"
	res["carLicense"] = "冀ACJXJX"
	res["BusinessPrice"] = "D53F7B90C8758CF6DFB0DACEEA3AABCD"
	res["OrderProvinceId"] = "9"
	res["op"] = "save"
	res["productType"] = "9"
	res["NewEdition"] = strconv.Itoa(NewEdition)
	res["RecordDate"] = "2019-12-01"
	res["Service"] = "2"
	res["telephone"] = ""
	res["equipmentNo"] = "B3B51F6D80E09F24FD6652BE50E51D80"
	res["ProgramId"] = ""
	res["Longitude"] = "116.312729"
	res["userId"] = strconv.Itoa(userid)
	res["Des"] = "订单备注"
	res["publishType"] = strconv.Itoa(procductlist)
	res["drivingLicense"] = ""
	res["VinCode"] = vin
	res["EngineNum"] = "Xjxjjxjddj"
	res["ConfigID"] = strconv.Itoa(configID)
	res["platType"] = "1"
	res["deviceInfo"] = "{\"brand\":\"OPPO\",\"model\":\"\",\"osVersion\":\"5.1.1\",\"platform\":\"android\",\"resolution\":\"1080*1920\"}"

	return res
}

//写入操作记录
func WriteOrderInfodetail(id int64, mo models.AppResultModel, Timelengthstr float64, vin string) {

	var tail models.OrderinfodetailModel
	tail.Status = 1
	tail.Oid = id
	if mo.Status != 100 {
		tail.Status = 2
	}
	tail.Des = mo.Msg
	tail.Vin = vin
	tail.Timelength = Timelengthstr
	tail.Save()
}

func SendPost(url string, resmap map[string]string, filename string) ([]byte, bool) {

	token := beego.AppConfig.String("app.userToken")
	resmap["sign"] = common.GetSign(resmap, token)

	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)
	var resbyte []byte
	for k, v := range resmap {
		body_writer.WriteField(k, v)
	}

	if filename != "" { //上传文件
		//  读取文件
		_, err := body_writer.CreateFormFile("application", filename)
		if err != nil {
			logs.Error("创建FormFile2文件信息异常！", err)
			return resbyte, false
		}
		fb2, err := ioutil.ReadFile(filename)
		if err != nil {
			logs.Error("打开文件异常!", err)
			return resbyte, false
		}
		body_buf.Write(fb2)
	}
	// 结束整个消息body
	body_writer.Close()

	req_reader := io.MultiReader(body_buf)
	req, err := http.NewRequest("POST", url, req_reader)
	if err != nil {
		logs.Error("SendPost error:", err)
		return resbyte, false
	}
	// 添加Post头
	req.Header.Set("Connection", "close")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Content-Type", body_writer.FormDataContentType())
	req.ContentLength = int64(body_buf.Len())
	// 发送消息
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logs.Error("读取回应消息异常SendPost:", err)
		return resbyte, false
	}
	logs.Debug("接收返回数据SendPost:", string(body))

	fmt.Println(string(body))

	return body, true
}

func SendPostys(url string, resmap map[string]string, filename string) ([]byte, bool, float64) {

	token := beego.AppConfig.String("app.userToken")
	resmap["sign"] = common.GetSign(resmap, token)
	var Timelength float64 = 0
	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)
	var resbyte []byte
	for k, v := range resmap {
		body_writer.WriteField(k, v)
	}

	defer func() {
		i := recover()
		if i != nil {
			marshal, _ := json.Marshal(resmap)
			common.Requestdderror(string(marshal))
		}
	}()

	if filename != "" { //上传文件
		//  读取文件
		_, err := body_writer.CreateFormFile("application", filename)
		if err != nil {
			logs.Error("创建FormFile2文件信息异常！", err)
			return resbyte, false, Timelength
		}
		fb2, err := ioutil.ReadFile(filename)
		if err != nil {
			logs.Error("打开文件异常!", err)
			return resbyte, false, Timelength
		}
		body_buf.Write(fb2)
	}
	// 结束整个消息body
	body_writer.Close()

	req_reader := io.MultiReader(body_buf)
	req, err := http.NewRequest("POST", url, req_reader)
	if err != nil {
		logs.Error("SendPost error:", err)
		return resbyte, false, Timelength
	}
	// 添加Post头
	req.Header.Set("Connection", "close")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Content-Type", body_writer.FormDataContentType())
	req.ContentLength = int64(body_buf.Len())
	// 发送消息
	client := &http.Client{}
	starttime := time.Now()
	resp, err := client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)
	endtime := time.Now()
	Timelength = endtime.Sub(starttime).Seconds()
	defer resp.Body.Close()
	if err != nil {
		logs.Error("读取回应消息异常SendPost:", err)
		logs.Debug("接收返回数据SendPost:", string(body))
		return resbyte, false, Timelength
	}

	return body, true, Timelength
}

func SendPostKG(taskid, userid int) int {

	url := beego.AppConfig.String("app.url") + "/App/TaskReconsideration.ashx"
	filename := beego.AppConfig.String("zip.kg")
	token := beego.AppConfig.String("app.userTokenet")
	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)

	// 1. 要上传的数据
	resmap := GetFastValueKG(taskid, userid)
	sigin := common.GetSign(resmap, token)
	for k, v := range resmap {
		body_writer.WriteField(k, v)
	}
	body_writer.WriteField("sign", sigin)

	// 3. 读取文件
	_, err := body_writer.CreateFormFile("application", filename)
	if err != nil {
		logs.Error("创建FormFile2文件信息异常9！", err)
		return 0
	}
	fb2, err := ioutil.ReadFile(filename)
	if err != nil {
		logs.Error("打开文件异常!SendPostFormFile9", err)
		return 0
	}
	body_buf.Write(fb2)

	// 结束整个消息body
	body_writer.Close()

	req_reader := io.MultiReader(body_buf)
	req, err := http.NewRequest("POST", url, req_reader)
	if err != nil {
		logs.Error("SendPostKG error:", err)
		return 0
	}
	// 添加Post头
	req.Header.Set("Connection", "close")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Content-Type", body_writer.FormDataContentType())
	req.ContentLength = int64(body_buf.Len())

	// 发送消息
	client := &http.Client{}
	starttime := time.Now()
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("读取回应消息异常SendPostKG:", err)
	}
	logs.Debug("接收返回数据SendPostKG:", string(body))
	endtime := time.Now()
	var res models.AppResultModel
	json.Unmarshal(body, &res)
	fmt.Println("调用接口时间", starttime, endtime)
	if res.Status == 100 {
		return 1
	} else {
		return 0
	}

}

//6张获取参数
func GetFastValueKG(taskid, userId int) map[string]string {

	res := make(map[string]string, 0)
	res["op"] = "AllPic"
	res["appVersion"] = "3.4.2"
	res["tokenid"] = "6"
	res["telephone"] = ""
	res["id"] = strconv.Itoa(taskid)
	res["equipmentNo"] = "B3B51F6D80E09F24FD6652BE50E51D80"
	res["platType"] = "1"
	res["userId"] = strconv.Itoa(userId)
	res["deviceInfo"] = "{\"brand\":\"OPPO\",\"model\":\"\",\"osVersion\":\"5.1.1\",\"platform\":\"android\",\"resolution\":\"1080*1920\"}"

	return res
}

func IniteData(userid int) {

	url := beego.AppConfig.String("app.url") + "/api/Banner/IniteData"
	token := beego.AppConfig.String("app.userTokenet")
	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)

	// 1. 要上传的数据

	resmap := make(map[string]string, 0)
	resmap["userId"] = strconv.Itoa(userid)
	sigin := common.GetSign(resmap, token)

	for k, v := range resmap {
		body_writer.WriteField(k, v)
	}
	body_writer.WriteField("sign", sigin)

	// 结束整个消息body
	body_writer.Close()

	req_reader := io.MultiReader(body_buf)
	req, err := http.NewRequest("POST", url, req_reader)
	if err != nil {
		logs.Error("SendPostKG error:", err)
		//return 0
	}
	// 添加Post头
	req.Header.Set("Connection", "close")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Content-Type", body_writer.FormDataContentType())
	req.ContentLength = int64(body_buf.Len())

	// 发送消息
	client := &http.Client{}

	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("读取回应消息异常SendPostKG:", err)
	}
	logs.Debug("接收返回数据SendPostKG:", string(body))
	//var res models.AppResultModel
	//json.Unmarshal(body, &res)

}
func GetServiceProgram(userid int) {

	url := beego.AppConfig.String("app.url") + "/api/PicSpecial/GetServiceProgram"
	//url ="http://localhost:45678/api/PicSpecial/GetServiceProgram"
	token := beego.AppConfig.String("app.userTokenet")
	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)

	// 1. 要上传的数据

	resmap := make(map[string]string, 0)
	resmap["appVersion"] = "3.5.0"
	resmap["Telephone"] = ""
	resmap["tokenid"] = "6"
	resmap["equipmentNo"] = "DCBFC057FC223134DEE10BDE04843A70"
	resmap["platType"] = "1"
	resmap["userId"] = strconv.Itoa(userid)
	resmap["deviceInfo"] = `{"brand":"HUAWEI","model":"EVA-AL00","osVersion":"8.0.0","platform":"android","resolution":"1080*1792"}`

	sigin := common.GetSign(resmap, token)

	for k, v := range resmap {
		body_writer.WriteField(k, v)
	}
	body_writer.WriteField("sign", sigin)

	// 结束整个消息body
	body_writer.Close()

	req_reader := io.MultiReader(body_buf)
	req, err := http.NewRequest("POST", url, req_reader)
	if err != nil {
		logs.Error("SendPostKG error:", err)
		//return 0
	}
	// 添加Post头
	req.Header.Set("Connection", "close")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Content-Type", body_writer.FormDataContentType())
	req.ContentLength = int64(body_buf.Len())

	// 发送消息
	client := &http.Client{}

	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("读取回应消息异常SendPostKG:", err)
	}
	logs.Debug("接收返回数据SendPostKG:", string(body))
	//var res models.AppResultModel
	//json.Unmarshal(body, &res)

}
