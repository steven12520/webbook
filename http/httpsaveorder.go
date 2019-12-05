package httpdate

import (
	"net/http"
	"strings"
	"fmt"
	"io"
	"compress/gzip"
	"io/ioutil"
	"strconv"
	"os"
	"bytes"
	"mime/multipart"
)

type RequesModel struct {
	RecordBrand     string
	TaskOwnerName   string
	tokenid         int
	appVersion      string
	Address         string
	OrderCityId     int
	Latitude        string
	Tasktel         string
	carLicense      string
	sign            string
	BusinessPrice   int
	OrderProvinceId int
	productType     int
	op              string
	NewEdition      int
	RecordDate      string
	Service         int
	telephone       string
	equipmentNo     string
	ProgramId       string
	Longitude       string
	userId          int
	deviceInfo      string
	Des             string
	publishType     int
	drivingLicense  string
	VinCode         string
	EngineNum       string
	ConfigID        int
	platType        int
}






func GetRequesDate(userid int,configID int,procductlist int,vin string,ordercount int,gocount int)string  {


	m:=""

	m+="RecordBrand=Xjxjjxdj&"
	m+="TaskOwnerName=Zjzj&"
	m+="tokenid=6&"
	m+="appVersion=3.3.9&"
	m+="Address=beijing&"
	m+="OrderCityId=901&"
	m+="Latitude=39.985395&"
	m+="Tasktel=15313666764&"
	m+="carLicense=冀ACJXJX&"
	m+="sign=D53F7B90C8758CF6DFB0DACEEA3AABCD&"
	m+="BusinessPrice=100000&"
	m+="OrderProvinceId=9&"
	m+="productType=9&"
	m+="op=save&"
	m+="NewEdition=1&"
	m+="RecordDate=2019-12-01&"
	m+="Service=2&"
	m+="telephone=&"
	m+="equipmentNo=B3B51F6D80E09F24FD6652BE50E51D80&"
	m+="ProgramId=2_1&"
	m+="Longitude=116.312729&"
	m+="userId="+strconv.Itoa(userid) +"&"
	m+="deviceInfo={\"brand\":\"HONOR\",\"model\":\"PRA-AL00X\",\"osVersion\":\"8.0.0\",\"platform\":\"android\",\"resolution\":\"1080*1794\"}&"
	m+="Des=订单备注&"
	m+="publishType=3&"
	m+="drivingLicense=&"
	m+="VinCode="+vin+"&"
	m+="EngineNum=Xjxjjxjddj&"
	m+="ConfigID=1&"
	m+="platType=1"

	return m
}

func File2Bytes(filename string) ([]byte, error) {

	// File
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// FileInfo:
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// []byte
	data := make([]byte, stats.Size())
	count, err := file.Read(data)
	if err != nil {
		return nil, err
	}
	fmt.Printf("read file %s len: %d \n", filename, count)
	return data, nil
}

func RequesSaveOrder(userid int,configID int,procductlist int,vin string,ordercount int,gocount int)  {
	url:="http://localhost:45678/app/TaskSave20160303.ashx?"

	value:=GetRequesDate(userid ,configID ,procductlist ,vin ,ordercount ,gocount )

	bpdy:=strings.NewReader(value)

	b,_:=File2Bytes("E:\\dada.zip")

	bpdy.Read(b)
	request, _ := http.NewRequest("POST", url, bpdy)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("User-Agent", "PostmanRuntime/7.20.1")
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Cache-Control", "no-cache")
	request.Header.Add("Postman-Token", "e01acf5-2475-4de8-a5f3-52ddea7effc5")
	request.Header.Add("Host", "localhost:45678")
	request.Header.Add("Accept-Encoding", "gzip, deflate")
	request.Header.Add("Content-Length", "0")
	request.Header.Add("Connection", "keep-alive")

	client := http.Client{}

	res, e := client.Do(request)
	if e != nil {
		fmt.Println("error")
	}

	defer res.Body.Close()
	var err error
	var reader io.ReadCloser
	if res.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(res.Body)
		if err != nil {
			return
		}
	} else {
		reader = res.Body
	}

	txt, er := ioutil.ReadAll(reader)
	if er != nil {
		fmt.Println("ReadAlleoorr")
		return
	}
	fmt.Println(string(txt))
}


func SendPostFormFile(userid int,configID int,procductlist int,vin string) {
	url:="http://localhost:45678/app/TaskSave20160303.ashx?"
	filename:="E:\\dada.zip"
	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)
	// boundary默认会提供一组随机数，也可以自己设置。
	body_writer.SetBoundary("Pp7Ye2EeWaFDdAY")


	// 1. 要上传的数据
	body_writer.WriteField("RecordBrand", "Xjxjjxdj")
	body_writer.WriteField("TaskOwnerName", "Zjzj")
	body_writer.WriteField("tokenid", "6")
	body_writer.WriteField("appVersion", "3.3.9")
	body_writer.WriteField("Address", "beijing")
	body_writer.WriteField("OrderCityId", "901")
	body_writer.WriteField("Latitude", "39.985395")
	body_writer.WriteField("Tasktel", "15313666764")
	body_writer.WriteField("carLicense", "冀ACJXJX")
	body_writer.WriteField("BusinessPrice", "D53F7B90C8758CF6DFB0DACEEA3AABCD")
	body_writer.WriteField("OrderProvinceId", "9")
	body_writer.WriteField("op", "save")
	body_writer.WriteField("productType", "9")
	body_writer.WriteField("NewEdition", "1")
	body_writer.WriteField("RecordDate", "2019-12-01")
	body_writer.WriteField("Service", "2")
	body_writer.WriteField("telephone", "")
	body_writer.WriteField("equipmentNo", "B3B51F6D80E09F24FD6652BE50E51D80")
	body_writer.WriteField("ProgramId", "2_1")
	body_writer.WriteField("Longitude", "116.312729")
	body_writer.WriteField("userId", strconv.Itoa(userid))
	body_writer.WriteField("deviceInfo", "{\"brand\":\"HONOR\",\"model\":\"PRA-AL00X\",\"osVersion\":\"8.0.0\",\"platform\":\"android\",\"resolution\":\"1080*1794\"}")
	body_writer.WriteField("Des", "订单备注")
	body_writer.WriteField("publishType",strconv.Itoa(procductlist) )
	body_writer.WriteField("drivingLicense", "")
	body_writer.WriteField("VinCode", vin)
	body_writer.WriteField("EngineNum", "Xjxjjxjddj")
	body_writer.WriteField("ConfigID", strconv.Itoa(configID))
	body_writer.WriteField("platType", "1")


	// 3. 读取文件
	_, err := body_writer.CreateFormFile("application", filename)
	if err != nil {
		fmt.Printf("创建FormFile2文件信息异常！", err)
		return
	}
	fb2, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("打开文件异常!", err)
		return
	}
	body_buf.Write(fb2)

	// 结束整个消息body
	body_writer.Close()

	//
	req_reader := io.MultiReader(body_buf)
	req, err := http.NewRequest("POST", url, req_reader)
	if err != nil {
		fmt.Printf("站点相机上传图片，创建上次请求异常！异常信息:", err)
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
	fmt.Println("发送回应数据:"+vin,string(body))
	return
}