package models

import "fmt"


type OrderinfoModel struct {
	Id	int64
	CreateName string
	Types string
	Ordercount int
	Vin	string
	Createtime string
}
type OrderinfodetailModel struct {
	Id 	int64
	Oid	int64
	Timelength float64
	Status	int
	Vin string
	Des string
}

func (o *OrderinfoModel)Save()bool  {

	sql := "INSERT INTO orderinfo(CreateName,Ordercount,Vin,Types)VALUES(?,?,?,?)"
	result, e := Db.Exec(sql, o.CreateName,o.Ordercount, o.Vin,o.Types)
	if e != nil {
		fmt.Println(" interfice add error", e.Error())
		return false
	} else {
		o.Id,_=result.LastInsertId()
		return true
	}

}
func (o *OrderinfodetailModel)Save()bool  {
	sql := "INSERT INTO orderinfodetail(oid,timelength,status,des,Vin)VALUES(?,?,?,?,?)"
	_, e := Db.Exec(sql, o.Oid, o.Timelength, o.Status, o.Des,o.Vin)
	if e != nil {
		fmt.Println(" interfice add error", e.Error())
		return false
	} else {
		return true
	}
}