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

	url := "http://localhost:45678/app/GetTaskDetail.ashx"

	token := beego.AppConfig.String("app.userTokenet")

	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)
	// 1. 要上传的数据

	resmap := make(map[string]string, 0)

	resmap["taskId"] = "1924614"
	resmap["userId"] = "890"
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

	url := "http://localhost:45678/api/Task/GetBackDetail"
	token := beego.AppConfig.String("app.userTokenet")

	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)
	// 1. 要上传的数据

	resmap := make(map[string]string, 0)

	resmap["Telephone"] = " "
	resmap["appVersion"] = "3.4.4"
	resmap["equipmentNo"] = "849367E75AD41C62FE6F9E6A283DB093"
	resmap["platType"] = "1"
	resmap["taskId"] = "1924595"
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

	token := beego.AppConfig.String("app.userTokenet")

	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)
	// 1. 要上传的数据

	resmap := make(map[string]string, 0)

	resmap["op"] = "login2"
	resmap["telphone"] = "13552077115"
	resmap["equipmentNo"] = "849367E75AD41C62FE6F9E6A28366655"
	resmap["lgcode"] = "yxdev001" //登录名
	resmap["tokenid"] = "6"
	resmap["checkcode"] = ""                                 //验证码
	resmap["loginCode"] = "FA407A5DDD70EB0D5F91A1189165336E" //登录秘钥

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
