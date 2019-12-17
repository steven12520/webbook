package models



type AppResultModel struct {
	Status        int    `json:"status"`
	Msg           string `json:"msg"`
	Tasktype      string    `json:"tasktype"`
	EstimatedTime string `json:"estimatedTime"`
}