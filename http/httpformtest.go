package httpdate

import (
	"../common"
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
)

func FastOnLineList() {

	url := "http://localhost:45678/app/FastOnLineList.ashx"

	token := beego.AppConfig.String("app.userTokenet")

	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)
	// 1. 要上传的数据

	resmap := make(map[string]string, 0)

	resmap["PageIndex"] = "1"
	resmap["Telephone"] = ""
	resmap["appVersion"] = "3.4.4"
	resmap["equipmentNo"] = "849367E75AD41C62FE6F9E6A283DB093"
	resmap["isShowAll"] = "4"
	resmap["platType"] = "1"
	resmap["sign"] = "8ECB4FDDDBF0E2F007292165B5908147"
	resmap["status"] = "-3"
	resmap["tokenid"] = "6"
	resmap["userId"] = "890"
	resmap["deviceInfo"] = "{\"brand\":\"honor\",\"model\":\"FRD-AL00\",\"osVersion\":\"8.0.0\",\"platform\":\"android\",\"resolution\":\"1080*1920\"}"

	sigin := common.GetSign(resmap, token)
	fmt.Println("sigin", sigin)
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
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取回应消息异常:", err)
	}
	logs.Debug("发送回应数据:", string(body))

	return
}

func GetTaskDetail() {

	url := "http://jiancetwo.sandbox.guchewang.com/app/GetTaskDetail.ashx"
	//url = "http://localhost:45678/app/GetTaskDetail.ashx"

	token := beego.AppConfig.String("app.userTokenet")

	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)
	// 1. 要上传的数据

	resmap := make(map[string]string, 0)

	resmap["taskId"] = "1926929"
	resmap["userId"] = "14692"
	resmap["tokenId"] = "6"

	sigin := common.GetSign(resmap, token)
	fmt.Println("sigin", sigin)
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
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取回应消息异常:", err)
	}
	fmt.Println("发送回应数据:", string(body))

	return
}

func GetBackDetail() {

	url := "http://jiancetwo.sandbox.guchewang.com/api/Task/GetBackDetail"
	token := beego.AppConfig.String("app.userTokenet")

	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)
	// 1. 要上传的数据

	resmap := make(map[string]string, 0)

	resmap["Telephone"] = " "
	resmap["appVersion"] = "3.4.4"
	resmap["equipmentNo"] = "849367E75AD41C62FE6F9E6A283DB093"
	resmap["platType"] = "1"
	resmap["taskId"] = "1925543"
	resmap["tokenid"] = "6"
	resmap["userId"] = "890"
	resmap["deviceInfo"] = "{\"brand\":\"honor\",\"model\":\"FRD-AL00\",\"osVersion\":\"8.0.0\",\"platform\":\"android\",\"resolution\":\"1080*1920\"}"

	sigin := common.GetSign(resmap, token)
	fmt.Println("sigin", sigin)
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
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取回应消息异常:", err)
	}
	fmt.Println("发送回应数据:", string(body))

	return
}

//获取验证码
func GetPhoneCheckNum() {

	url := "http://jiancetwo.sandbox.guchewang.com/app/GetPhoneCheckNum.ashx"
	//url = "http://jiancetwo.sandbox.guchewang.com/app/GetPhoneCheckNum.ashx"

	token := beego.AppConfig.String("app.userTokenet")

	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)
	// 1. 要上传的数据

	resmap := make(map[string]string, 0)

	resmap["tel"] = "13552077115"
	resmap["UserId"] = ""
	resmap["EquipmentNo"] = "849367E75AD41C62FE6F9E6A28366655"
	resmap["username"] = "yxdev001"
	resmap["type"] = "1"
	resmap["tokenid"] = "6"

	sigin := common.GetSign(resmap, token)
	fmt.Println("sigin", sigin)
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
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取回应消息异常:", err)
	}
	fmt.Println("发送回应数据:", string(body))

	return
}

//手机登录 UserHandler
func UserHandler() {

	url := "http://jiancetwo.sandbox.guchewang.com/app/UserHandler.ashx"
	url = "http://localhost:45678/app/UserHandler.ashx"
	token := beego.AppConfig.String("app.userTokenet")

	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)
	// 1. 要上传的数据

	resmap := make(map[string]string, 0)

	resmap["op"] = "login"
	resmap["telphone"] = "15313256075"
	resmap["equipmentNo"] = "849367E75AD41C62FE6F9E6A28366655"
	resmap["lgcode"] = "yxdev001" //登录名
	resmap["tokenid"] = "6"
	resmap["checkcode"] = ""
	resmap["pwd"] = "111111"
	//验证码
	resmap["loginCode"] = "7E428B97259A2E1A7858458B7D4ECB25" //登录秘钥

	sigin := common.GetSign(resmap, token)
	fmt.Println("sigin", sigin)
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
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取回应消息异常:", err)
	}
	fmt.Println("发送回应数据:", string(body))

	return
}

func Getmfe() {
	url := "https://yunfudao.mofangge.xin/Home/Login"

	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)
	// 1. 要上传的数据

	resmap := make(map[string]string, 0)

	resmap["useracount"] = "49698965368"
	resmap["password"] = "son22372"
	resmap["valicode"] = ""

	for k, v := range resmap {
		body_writer.WriteField(k, v)
	}

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
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "50")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Host", "yunfudao.mofangge.xin")
	req.Header.Set("Origin", "https://www.mofangge.xin")
	req.Header.Set("Referer", "https://www.mofangge.xin/home/index/word-detail.html")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")

	req.ContentLength = int64(body_buf.Len())

	// 发送消息
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取回应消息异常:", err)
	}
	fmt.Println("发送回应数据:", string(body))

	return
}

func PretrailSubmitPass(taskId, userId int) {
	url := beego.AppConfig.String("pgs.url") + "/APP/Pretrial/PretrialV2.ashx"
	url = "http://localhost:45678/APP/Pretrial/PretrialV2.ashx"
	m := make(map[string]string, 0)
	m["op"] = "PretrailSubmitPass"
	m["taskId"] = strconv.Itoa(taskId)
	m["userId"] = strconv.Itoa(userId)

	SendPostys(url, m, "")
}
