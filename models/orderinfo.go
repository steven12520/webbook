package models

import (
	"fmt"
	"github.com/astaxie/beego/logs"
)


type OrderinfoModel struct {
	Id         int64
	CreateName string
	Types      string
	Ordercount int
	Vin        string
	Gocount    int
	Gotype     int
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

func (b OrderinfoModel) GetList(page, pagesize int) []OrderinfoModel {

	sql := "SELECT Id,CreateName,Types,Ordercount,Vin,Gocount,Gotype,Createtime FROM orderinfo ORDER BY id DESC"
	sql += fmt.Sprintf(" limit %d,%d", (page-1)*pagesize, pagesize)

	rows, e := Db.Query(sql)
	if e != nil {
		logs.Error("GetBorrowBookList error", e.Error())
	}
	list := make([]OrderinfoModel, 0)
	for rows.Next() {
		var br OrderinfoModel
		//b_bookId,b_userId,b_createTime,b_toAlsoTime,b_alsoTime,b_type,b_bookName,b_author,b_loginname
		rows.Scan(&br.Id, &br.CreateName, &br.Types, &br.Ordercount, &br.Vin, &br.Gocount, &br.Gotype, &br.Createtime)
		list = append(list, br)
	}
	return list
}

func (b OrderinfoModel) GetCount() (count int64) {
	sql := "select count(*)as 'count' from orderinfo"
	rows, e := Db.Query(sql)
	for rows.Next() {
		rows.Scan(&count)
	}
	if e != nil {
		logs.Error("orderinfo GetCount error",e.Error())
	}
	return count
}

func (b OrderinfodetailModel) GetList(page, pagesize,id int) []OrderinfodetailModel {

	sql := "SELECT * FROM orderinfodetail"
	sql += fmt.Sprintf(" where oid=%d limit %d,%d",id, (page-1)*pagesize, pagesize)

	list := make([]OrderinfodetailModel, 0)
	er := Db.Select(&list, sql)
	if er != nil {
		logs.Error("OrderinfodetailModel GetList  error ", er.Error())
	}
	return list
}

func (b OrderinfodetailModel) GetCount(oid int) (count int64) {
	sql := "select count(*)as 'count' from orderinfodetail"
	sql+=fmt.Sprintf(" where oid=%d ",oid)
	rows, e := Db.Query(sql)
	for rows.Next() {
		rows.Scan(&count)
	}
	if e != nil {
		logs.Error("OrderinfodetailModel GetCount error",e.Error())
	}
	return count
}

func (b *OrderinfoModel) Delete() bool {

	sql := "delete from orderinfo WHERE id=?"
	_, e := Db.Exec(sql, b.Id)

	if e == nil {
		return true
	} else {
		logs.Error("删除失败", e.Error())
		return false
	}
}
func (b *OrderinfodetailModel) Delete() bool {

	sql := "delete from orderinfodetail WHERE oid=?;"
	_, e := Db.Exec(sql, b.Oid)

	if e == nil {
		return true
	} else {
		logs.Error("删除失败", e.Error())
		return false
	}
}



func (o *OrderinfoModel)Save()bool  {

	sql := "INSERT INTO orderinfo(CreateName,Ordercount,Vin,Types,Gocount,Gotype)VALUES(?,?,?,?,?,?)"
	result, e := Db.Exec(sql, o.CreateName,o.Ordercount, o.Vin,o.Types,o.Gocount,o.Gotype)
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