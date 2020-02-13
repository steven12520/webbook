package models

import (
	"github.com/astaxie/beego/logs"
)

type Ysyinfo struct {
	Id         int64
	Yname      string
	Avgcount   int
	Timelength int
	Types      int
	Username   string
	Createtime string
}

func (br *Ysyinfo) Add() bool {

	sql := "INSERT INTO ysyinfo(yname,avgcount,timelength,types,username)VALUES(?,?,?,?,?)"

	result, e := Db.Exec(sql, br.Yname, br.Avgcount, br.Timelength, br.Types, br.Username)
	if e != nil {
		logs.Error(" Ysyinfo add error", e.Error())
		return false
	} else {
		br.Id, _ = result.LastInsertId()
		return true
	}
}

type Ysyinfodetail struct {
	Id       int64
	Ysyid    int64
	Vin      string
	Satus    int
	Satusmsg string
	Userid   int
}

func (br *Ysyinfodetail) Add() bool {

	sql := "INSERT INTO ysyinfodetail(ysyid,vin,satus,satusmsg,userid)VALUES(?,?,?,?,?)"

	result, e := Db.Exec(sql, br.Ysyid, br.Vin, br.Satus, br.Satusmsg, br.Userid)
	if e != nil {
		logs.Error(" Ysyinfodetail add error", e.Error())
		return false
	} else {
		br.Id, _ = result.LastInsertId()
		return true
	}
}
func (b Ysyinfodetail) Update() bool {

	sql := "UPDATE ysyinfodetail SET vin=?,satus=?,satusmsg=? WHERE id=?"
	_, e := Db.Exec(sql, b.Vin, b.Satus, b.Satusmsg, b.Id)

	if e == nil {
		return true
	} else {
		logs.Error("Ysyinfodetail 保存失败", e.Error())
		return false
	}
}

type Ysyinfodetailinterfice struct {
	Id        int64
	Ysyid     int64
	Ysydid    int64
	Iname     string
	Timelen   float64
	Txt       string
	Status    int
	Cratetime string
}

func (br Ysyinfodetailinterfice) Add() bool {

	sql := "INSERT INTO ysyinfodetailinterfice(ysyid,ysydid,iname,timelen,txt,status)VALUES(?,?,?,?,?,?)"

	_, e := Db.Exec(sql, br.Ysyid, br.Ysydid, br.Iname, br.Timelen, br.Txt, br.Status)
	if e != nil {
		logs.Error(" Ysyinfodetailinterfice add error", e.Error())
		return false
	} else {
		return true
	}
}
