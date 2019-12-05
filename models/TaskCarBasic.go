package models

import (
	"fmt"
	"strconv"
)

type TaskCarBasicModel struct {
	Id                 int
	OrderNo            string
	SourceID           int
	CityID             int
	ProvID             int
	Des                string
	LikeMan            string
	LikeTel            string
	LikeAddr           string
	vin                string
	CarLicense         string
	RecordBrand        string
	EngineNum          string
	RecordDate         string
	MakeID             int
	ModelID            int
	StyleID            int
	Color              int
	Mileage            int
	Service            int
	AssessmentPrace    int
	AssessmentDes      string
	UserID             int
	Status             int
	CreateTime         string
	UpdateTime         string
	StartTime          string
	EndTime            string
	Exhaust            string
	Seating            int
	DrivingMode        int
	Transmission       int
	FuelType           int
	ProductionTime     string
	Certificates       int
	ManufacturerPrice  int
	SetGroupID         int
	TaskType           int
	TaskBackNum        int
	TaskBackReason     string
	AppraiseBackNum    int
	AppraiseBackReason string
	TransferCount      int
	Insurance          string
	Inspection         string
	CreateUserId       int
	YXOrderNo          string
	VideoPath          string
	RegisterCityID     int
	RegisterProvID     int
	BusinessPrice      int
	SalePrice          int
	Perf_DriveType     string
	TransmissionType   string
	Engine_Exhaust     string
	Fuel               string
	CarType            string
	ProductType        int
	ProgrammeId        int
	IsComplete         int
	OnCardCityID       int
	OnCardProvID       int
	OrderTelphone      string
	OrderStatus        int
	ReViewType         int
	EstimatedTime      string
	PretrialUser       int
	TaskVersion        int
	IsPriority         int
	Channel            int
	ProgramId          string
	CarFullName        string
	SignatureReport    string
	CertPdf            string
	AccidentBasisType  string
	RandomYS           int
	RandomPGS          int
	IsRead             int
}

func (t *TaskCarBasicModel) IsVinRepeat(Vin string, SourceID int)int  {

	sql := "select count(1) as counts  from TaskCarBasic WITH(NOLOCK) where vin='"+Vin+"'  and (([EndTime] >=GETDATE() and (status=0 or status=1 or status=7 or status=8 or status =10)) or  ((status=2 or status=4 or status=5) )) and status not in(-1,3,6)  and SourceID =  "+strconv.Itoa(SourceID)

	rows, e := Dbsql.Query(sql)
	counts:=0
	for rows.Next() {
		rows.Scan(&counts)
	}
	if e != nil {
		fmt.Println("(p *PicSpecialUserSetModel)get error!")
	}
	return counts
}




