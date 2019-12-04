package httpdate

import (
	"net/http"
	"strings"
	"fmt"
	"io"
	"compress/gzip"
	"io/ioutil"
	"strconv"
	"github.com/typa01/go-utils"
)

func GetRequesDate(userid int,configID int,procductlist int,vin string,ordercount int,gocount int)string  {

	builder1 := tsgutils.NewStringBuilder()
	builder1.Append("{")

	builder1.Append("\"RecordBrand\":\"Xjxjjxdj\",")
	builder1.Append("\"TaskOwnerName\":\"Zjzj\",")
	builder1.Append("\"tokenid\":\"6\",")
	builder1.Append("\"appVersion\":\"3.3.9\",")
	builder1.Append("\"Address\":\"beijing\",")
	builder1.Append("\"OrderCityId\":\"901\",")
	builder1.Append("\"Latitude\":\"39.985395\",")
	builder1.Append("\"Tasktel\":\"15313666764\",")
	builder1.Append("\"carLicense\":\"冀ACJXJX\",")
	builder1.Append("\"sign\":\"D53F7B90C8758CF6DFB0DACEEA3AABCD\",")
	builder1.Append("\"BusinessPrice\":\"100000\",")
	builder1.Append("\"OrderProvinceId\":\"9\",")
	builder1.Append("\"productType\":\"9\",")
	builder1.Append("\"op\":\"save\",")
	builder1.Append("\"NewEdition\":\"1\",")
	builder1.Append("\"RecordDate\":\"2019-12-01\",")
	builder1.Append("\"Service\":\"2\",")
	builder1.Append("\"telephone\":\"\",")
	builder1.Append("\"equipmentNo\":\"B3B51F6D80E09F24FD6652BE50E51D80\",")
	builder1.Append("\"ProgramId\":\"2_1\",")
	builder1.Append("\"Longitude\":\"116.312729\",")
	builder1.Append("\"userId\":\""+strconv.Itoa(userid) +"\",")
	builder1.Append("\"deviceInfo\":\"{\"brand\":\"HONOR\",\"model\":\"PRA-AL00X\",\"osVersion\":\"8.0.0\",\"platform\":\"android\",\"resolution\":\"1080*1794\"}\",")
	builder1.Append("\"Des\":\"订单备注\",")
	builder1.Append("\"publishType\":\"3\",")
	builder1.Append("\"drivingLicense\":\"\",")
	builder1.Append("\"VinCode\":\""+vin+"\",")
	builder1.Append("\"EngineNum\":\"Xjxjjxjddj\",")
	builder1.Append("\"ConfigID\":\"1\",")
	builder1.Append("\"platType\":\"1\"")
	builder1.Append("}")


	return builder1.ToString()
}



func RequesSaveOrder(userid int,configID int,procductlist int,vin string,ordercount int,gocount int)  {
	url:="http://localhost:45678/app/TaskSave20160303.ashx"

	value:=GetRequesDate(userid ,configID ,procductlist ,vin ,ordercount ,gocount )

	value="{\"op\":\"save\",\"UserID\":1726,\"TaskID\":0}"

	request, _ := http.NewRequest("POST", url, strings.NewReader(value))



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