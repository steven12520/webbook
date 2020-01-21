package models

import (
	"fmt"
	"strconv"
	"strings"
)

type TaskCarBasicModel struct {
	Id                 int
	IdList 			   string
	OrderNo            string
	SourceID           int
	CityID             int
	ProvID             int
	Des                string
	LikeMan            string
	LikeTel            string
	LikeAddr           string
	Vin                string
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

func (t *TaskCarBasicModel) GetList()[]TaskCarBasicModel  {
	sql := "SELECT Id, CreateUserId, Status, vin, OrderNo FROM TaskCarBasic where 1=1 "

	strwhere:=""
	if t.Vin != "" && len(t.Vin)>=3 {
		strwhere+= " and vin like '"+strings.Replace(t.Vin,"-","",0) +"%' "
	}
	if t.StartTime != "" {
		strwhere+= " and CreateTime >'"+t.StartTime+" 00:00:00' "
	}
	if t.EndTime != "" {
		strwhere+= " and CreateTime <='"+t.EndTime+" 23:59:59' "
	}
	if t.CreateUserId>0 {
		strwhere+= " and CreateUserId ="+ strconv.Itoa(t.CreateUserId)
	}
	if t.Id>0 {
		strwhere+= " and id ="+ strconv.Itoa(t.Id)
	}
	if strwhere=="" {
		return nil
	}
	sql+=strwhere

	rows, e := Dbsql.Query(sql)
	if e != nil {
		fmt.Println("TaskCarBasicModel List error", e.Error())
	}
	list := make([]TaskCarBasicModel, 0)
	for rows.Next() {
		var br TaskCarBasicModel
		rows.Scan(&br.Id, &br.CreateUserId, &br.Status, &br.Vin, &br.OrderNo)
		list = append(list, br)
	}
	return list
}

func (t *TaskCarBasicModel) GetRes()int  {
	if t.Id==0 {
		return 0
	}
	sql := "SELECT  Reconsideration  from TaskOnLineExpand where TaskID ="+strconv.Itoa(t.Id)

	rows, e := Dbsql.Query(sql)
	if e != nil {
		fmt.Println("GetRes  error", e.Error())
	}
	Reconsideration := 0
	for rows.Next() {
		rows.Scan(&Reconsideration)
	}
	return Reconsideration
}

//删除订单
func (t *TaskCarBasicModel)DeleteOrder()int  {

	sql := "delete from TaskCarBasic WHERE id in("+t.IdList+") "

	_, e := Dbsql.Exec(sql)

	if e == nil {
		return 1
	} else {
		fmt.Println("DeleteOrder 删除失败", e.Error())
		return 0
	}
}
//删除派单表数据
func (t *TaskCarBasicModel)DeleteAssignedTask()int  {

	sql := "delete from AssignedTask WHERE TaskID in("+t.IdList+");delete from AppraiserOrderQueue WHERE TaskID in("+t.IdList+");delete from PreAuditOrderQueue WHERE TaskID in("+t.IdList+"); "

	_, e := Dbsql.Exec(sql)

	if e == nil {
		return 1
	} else {
		fmt.Println("DeleteAssignedTask 删除失败", e.Error())
		return 0
	}
}
//删除登录状态
func (t *TaskCarBasicModel)DeleteLoginStatusRecords()int  {

	sql := "delete from LoginStatusRecords ; DELETE from PreAuditAppraiserQueue ;"

	_, e := Dbsql.Exec(sql)

	if e == nil {
		return 1
	} else {
		fmt.Println("DeleteLoginStatusRecords 删除失败", e.Error())
		return 0
	}
}






