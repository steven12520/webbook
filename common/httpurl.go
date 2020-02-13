package common

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func Requestdderror(txt string) {
	//如果有未发送新闻 请求钉钉webhook
	body := "{\"msgtype\": \"text\",\"text\": {\"content\": \"" + txt + "！\"},\"at\": {\"atMobiles\": [	\"15313256075\"],\"isAtAll\": false}}"

	url := "https://oapi.dingtalk.com/robot/send?access_token=1d67c5c56787d95de9cd364884f4b9aca6a58a2cd76500c8150fc1bdc9055920"
	jsonValue := []byte(body)
	//发送消息到钉钉群使用webhook
	//钉钉文档 https://open-doc.dingtalk.com/docs/doc.htm?spm=a219a.7629140.0.0.karFPe&treeId=257&articleId=105735&docType=1
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {

		log.Println(err)
	} else {
		fmt.Println(resp)
	}
}
